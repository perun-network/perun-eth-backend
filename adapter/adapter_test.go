package adapter

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/perun-network/perun-eth-backend/bindings/liquiditypool"
	"github.com/stretchr/testify/require"
)

type mockContract struct {
	withdrawable *big.Int
	shares       *big.Int
	locked       *big.Int
	totalAssets  *big.Int
	totalLocked  *big.Int
	operator     common.Address

	withdrawErr error
	sharesErr   error
	lockedErr   error
	fundErr     error
	settleErr   error
	totalErr    error
	lockedTErr  error
	opErr       error

	fundTx   *types.Transaction
	settleTx *types.Transaction

	lastSettleValue *big.Int
}

func (m *mockContract) WithdrawableETH(_ *bind.CallOpts) (*big.Int, error) {
	if m.withdrawErr != nil {
		return nil, m.withdrawErr
	}
	return new(big.Int).Set(m.withdrawable), nil
}

func (m *mockContract) SharesOf(_ *bind.CallOpts, _ common.Address) (*big.Int, error) {
	if m.sharesErr != nil {
		return nil, m.sharesErr
	}
	return new(big.Int).Set(m.shares), nil
}

func (m *mockContract) LockedByChannel(_ *bind.CallOpts, _ [32]byte) (*big.Int, error) {
	if m.lockedErr != nil {
		return nil, m.lockedErr
	}
	return new(big.Int).Set(m.locked), nil
}

func (m *mockContract) FundChannel(_ *bind.TransactOpts, _ [32]byte, _ *big.Int) (*types.Transaction, error) {
	if m.fundErr != nil {
		return nil, m.fundErr
	}
	return m.fundTx, nil
}

func (m *mockContract) SettleChannel(opts *bind.TransactOpts, _ [32]byte) (*types.Transaction, error) {
	if opts.Value != nil {
		m.lastSettleValue = new(big.Int).Set(opts.Value)
	}
	if m.settleErr != nil {
		return nil, m.settleErr
	}
	return m.settleTx, nil
}

func (m *mockContract) Operator(_ *bind.CallOpts) (common.Address, error) {
	if m.opErr != nil {
		return common.Address{}, m.opErr
	}
	return m.operator, nil
}

func (m *mockContract) TotalAssets(_ *bind.CallOpts) (*big.Int, error) {
	if m.totalErr != nil {
		return nil, m.totalErr
	}
	return new(big.Int).Set(m.totalAssets), nil
}

func (m *mockContract) TotalLockedETH(_ *bind.CallOpts) (*big.Int, error) {
	if m.lockedTErr != nil {
		return nil, m.lockedTErr
	}
	return new(big.Int).Set(m.totalLocked), nil
}

func fakeTx(selector []byte, value *big.Int) *types.Transaction {
	to := common.HexToAddress("0x1234")
	return types.NewTx(&types.LegacyTx{
		Nonce:    1,
		To:       &to,
		Gas:      21000,
		GasPrice: big.NewInt(1),
		Value:    value,
		Data:     selector,
	})
}

func TestFundChannel_HappyPath(t *testing.T) {
	m := &mockContract{
		withdrawable: big.NewInt(100),
		locked:       big.NewInt(0),
		fundTx:       fakeTx([]byte{0x26, 0x0b, 0x60, 0x86}, big.NewInt(0)),
	}

	var captured *types.Transaction
	a := newTestAdapter(nil, m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		captured = tx
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	err := a.FundChannel(context.Background(), [32]byte{1}, big.NewInt(42))
	require.NoError(t, err)
	require.NotNil(t, captured)
	require.Equal(t, []byte{0x26, 0x0b, 0x60, 0x86}, captured.Data())
}

func TestFundChannel_AlreadyFunded(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(100), locked: big.NewInt(1)}
	a := newTestAdapter(nil, m, nil, nil)

	err := a.FundChannel(context.Background(), [32]byte{2}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestFundChannel_InsufficientFree(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(9), locked: big.NewInt(0)}
	a := newTestAdapter(nil, m, nil, nil)

	err := a.FundChannel(context.Background(), [32]byte{2}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestSettleChannel_HappyPath(t *testing.T) {
	m := &mockContract{
		locked:   big.NewInt(15),
		settleTx: fakeTx([]byte{0xfd, 0x3d, 0x31, 0x99}, big.NewInt(15)),
	}

	var captured *types.Transaction
	a := newTestAdapter(nil, m, func(context.Context) (*bind.TransactOpts, error) {
		return &bind.TransactOpts{}, nil
	}, func(_ context.Context, tx *types.Transaction) (*types.Receipt, error) {
		captured = tx
		return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
	})

	ret := big.NewInt(20)
	err := a.SettleChannel(context.Background(), [32]byte{3}, ret)
	require.NoError(t, err)
	require.Equal(t, ret, m.lastSettleValue)
	require.Equal(t, []byte{0xfd, 0x3d, 0x31, 0x99}, captured.Data())
}

func TestSettleChannel_BelowPrincipal(t *testing.T) {
	m := &mockContract{locked: big.NewInt(11)}
	a := newTestAdapter(nil, m, nil, nil)

	err := a.SettleChannel(context.Background(), [32]byte{4}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestSettleChannel_ChannelNotFound(t *testing.T) {
	m := &mockContract{locked: big.NewInt(0)}
	a := newTestAdapter(nil, m, nil, nil)

	err := a.SettleChannel(context.Background(), [32]byte{4}, big.NewInt(10))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}

func TestGetPoolState_HappyPath(t *testing.T) {
	m := &mockContract{totalAssets: big.NewInt(123), totalLocked: big.NewInt(45)}
	a := newTestAdapter(nil, m, nil, nil)

	reserve, locked, err := a.GetPoolState(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 123, reserve)
	require.EqualValues(t, 45, locked)
}

func TestGetOperator_Delegated(t *testing.T) {
	want := common.HexToAddress("0x1000")
	m := &mockContract{operator: want}
	a := newTestAdapter(nil, m, nil, nil)

	got, err := a.GetOperator(context.Background())
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestSharesOf_HappyPath(t *testing.T) {
	m := &mockContract{shares: big.NewInt(777)}
	a := newTestAdapter(nil, m, nil, nil)

	got, err := a.SharesOf(context.Background(), common.HexToAddress("0xABCD"))
	require.NoError(t, err)
	require.Equal(t, big.NewInt(777), got)
}

func TestSharesOf_RetriableOnError(t *testing.T) {
	m := &mockContract{sharesErr: errors.New("node down")}
	a := newTestAdapter(nil, m, nil, nil)

	_, err := a.SharesOf(context.Background(), common.HexToAddress("0xABCD"))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestWithdrawableETH_HappyPath(t *testing.T) {
	m := &mockContract{withdrawable: big.NewInt(555)}
	a := newTestAdapter(nil, m, nil, nil)

	got, err := a.WithdrawableETH(context.Background())
	require.NoError(t, err)
	require.Equal(t, big.NewInt(555), got)
}

func TestWithdrawableETH_RetriableOnError(t *testing.T) {
	m := &mockContract{withdrawErr: errors.New("node down")}
	a := newTestAdapter(nil, m, nil, nil)

	_, err := a.WithdrawableETH(context.Background())
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestPoolMetadata(t *testing.T) {
	addr := common.HexToAddress("0xABCD")
	a := newTestAdapter(nil, &mockContract{}, nil, nil)
	a.poolAddress = addr
	a.chainID = big.NewInt(1337)

	gotAddr, gotChain := a.PoolMetadata()
	require.Equal(t, addr, gotAddr)
	require.Equal(t, big.NewInt(1337), gotChain)

	// Returned chainID must be a copy; mutating it must not affect the adapter.
	gotChain.SetInt64(1)
	_, again := a.PoolMetadata()
	require.Equal(t, big.NewInt(1337), again)
}

func TestPoolMetadata_NilChainID(t *testing.T) {
	a := newTestAdapter(nil, &mockContract{}, nil, nil)
	addr, chainID := a.PoolMetadata()
	require.Equal(t, common.Address{}, addr)
	require.Nil(t, chainID)
}

func TestRetriableOnRPCTimeout(t *testing.T) {
	err := classifyEthError(context.DeadlineExceeded)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrRetriable)
}

func TestOperatorUpdatedEvent(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(liquiditypool.LiquidityPoolABI))
	require.NoError(t, err)

	prev := common.HexToAddress("0x1111")
	next := common.HexToAddress("0x2222")
	log := types.Log{
		Topics: []common.Hash{
			parsed.Events["OperatorUpdated"].ID,
			common.BytesToHash(common.LeftPadBytes(prev.Bytes(), 32)),
			common.BytesToHash(common.LeftPadBytes(next.Bytes(), 32)),
		},
	}

	decoded := struct {
		PreviousOperator common.Address
		NewOperator      common.Address
	}{}
	err = abi.ParseTopics(&decoded, parsed.Events["OperatorUpdated"].Inputs, log.Topics[1:])
	require.NoError(t, err)

	event := OperatorUpdatedEvent{Previous: decoded.PreviousOperator, New: decoded.NewOperator}
	require.Equal(t, prev, event.Previous)
	require.Equal(t, next, event.New)
	require.Equal(t, liquiditypool.LiquidityPoolOperatorUpdatedTopic, log.Topics[0])
}

func TestSelectors_MatchABI(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(liquiditypool.LiquidityPoolABI))
	require.NoError(t, err)
	require.Equal(t, []byte{0x26, 0x0b, 0x60, 0x86}, parsed.Methods["fundChannel"].ID)
	require.Equal(t, []byte{0xfd, 0x3d, 0x31, 0x99}, parsed.Methods["settleChannel"].ID)
}

func TestClassifyUnknownIsDeterministic(t *testing.T) {
	err := classifyEthError(errors.New("some deterministic issue"))
	require.Error(t, err)
	require.ErrorIs(t, err, ErrDeterministic)
}
