// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquiditypool

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LiquidityPoolMetaData contains all meta data concerning the LiquidityPool contract.
var LiquidityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_settlementWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minFeeBps\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBond\",\"type\":\"uint256\"}],\"name\":\"BondPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBond\",\"type\":\"uint256\"}],\"name\":\"BondWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"ChannelExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ChannelFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeGain\",\"type\":\"uint256\"}],\"name\":\"ChannelSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesMinted\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesBurned\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channelDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"expireChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lockedByChannel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"minSettlementValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operatorBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"}],\"name\":\"previewDepositShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"previewWithdrawETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelId\",\"type\":\"bytes32\"}],\"name\":\"settleChannel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToBurn\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethAmountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawableETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161182b38038061182b83398101604081905261002f9161019c565b600160005561003d3361014a565b6001600160a01b03831661008b5760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b60448201526064015b60405180910390fd5b600082116100db5760405162461bcd60e51b815260206004820152601960248201527f496e76616c696420736574746c656d656e742077696e646f77000000000000006044820152606401610082565b612710811061011e5760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206d696e2066656560881b6044820152606401610082565b600480546001600160a01b0319166001600160a01b03949094169390931790925560805260a0526101df565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806000606084860312156101b157600080fd5b83516001600160a01b03811681146101c857600080fd5b602085015160409095015190969495509392505050565b60805160a0516116126102196000396000818161022401528181610ba101526111a401526000818161043101526106f001526116126000f3fe6080604052600436106101855760003560e01c80638afe47cf116100d1578063b6cc49da1161008a578063e1a4521811610064578063e1a45218146104a8578063f2fde38b146104be578063f5eb42dc146104de578063fd3d31991461051457600080fd5b8063b6cc49da14610453578063c3daab9614610480578063d0e30db0146104a057600080fd5b80638afe47cf146103a35780638da5cb5b146103b9578063a517458f146103d7578063b3ab15fb146103f7578063b453de0914610417578063b4a7bdf91461041f57600080fd5b806342161b741161013e5780636574bcd4116101185780636574bcd4146103395780636913e7cc14610359578063715018a61461036e57806378f3bfe21461038357600080fd5b806342161b74146102b4578063570ca735146102d45780635f693d0a1461030c57600080fd5b806301e1d114146101ea578063071de20d14610212578063260b6086146102465780632e1a7d4d1461026857806339b8b076146102885780633a98ef391461029e57600080fd5b366101e55760405162461bcd60e51b815260206004820152602560248201527f557365206465706f7369742c20626f6e64455448206f7220736574746c654368604482015264185b9b995b60da1b60648201526084015b60405180910390fd5b600080fd5b3480156101f657600080fd5b506101ff610527565b6040519081526020015b60405180910390f35b34801561021e57600080fd5b506101ff7f000000000000000000000000000000000000000000000000000000000000000081565b34801561025257600080fd5b50610266610261366004611467565b610549565b005b34801561027457600080fd5b506101ff610283366004611489565b6107b9565b34801561029457600080fd5b506101ff60055481565b3480156102aa57600080fd5b506101ff60025481565b3480156102c057600080fd5b506102666102cf366004611489565b6109a7565b3480156102e057600080fd5b506004546102f4906001600160a01b031681565b6040516001600160a01b039091168152602001610209565b34801561031857600080fd5b506101ff610327366004611489565b60076020526000908152604090205481565b34801561034557600080fd5b506101ff610354366004611489565b610af3565b34801561036557600080fd5b506101ff610b0c565b34801561037a57600080fd5b50610266610b1c565b34801561038f57600080fd5b506101ff61039e366004611489565b610b52565b3480156103af57600080fd5b506101ff60085481565b3480156103c557600080fd5b506001546001600160a01b03166102f4565b3480156103e357600080fd5b506101ff6103f2366004611489565b610b8a565b34801561040357600080fd5b506102666104123660046114a2565b610be1565b610266610ca6565b34801561042b57600080fd5b506101ff7f000000000000000000000000000000000000000000000000000000000000000081565b34801561045f57600080fd5b506101ff61046e366004611489565b60066020526000908152604090205481565b34801561048c57600080fd5b5061026661049b366004611489565b610d70565b6101ff610f17565b3480156104b457600080fd5b506101ff61271081565b3480156104ca57600080fd5b506102666104d93660046114a2565b611062565b3480156104ea57600080fd5b506101ff6104f93660046114a2565b6001600160a01b031660009081526003602052604090205490565b610266610522366004611489565b6110fd565b60006008546005544761053a91906114e1565b61054491906114f4565b905090565b6004546001600160a01b031633146105735760405162461bcd60e51b81526004016101dc90611507565b6002600054036105955760405162461bcd60e51b81526004016101dc90611537565b6002600055806105dc5760405162461bcd60e51b81526020600482015260126024820152710416d6f756e74206d757374206265203e20360741b60448201526064016101dc565b6000828152600660205260409020546001116106335760405162461bcd60e51b815260206004820152601660248201527510da185b9b995b08185b1c9958591e48199d5b99195960521b60448201526064016101dc565b61063b610b0c565b81111561068a5760405162461bcd60e51b815260206004820152601b60248201527f496e73756666696369656e742066726565206c6971756964697479000000000060448201526064016101dc565b8060085461069891906114f4565b60055410156106e95760405162461bcd60e51b815260206004820152601a60248201527f496e73756666696369656e7420626f6e6420636f76657261676500000000000060448201526064016101dc565b60006107157f0000000000000000000000000000000000000000000000000000000000000000426114f4565b60008481526006602090815260408083208690556007909152812082905560088054929350849290919061074a9084906114f4565b9091555050600454610765906001600160a01b0316836112c8565b60045460408051848152602081018490526001600160a01b039092169185917f921e723c5ab46e7fad6ecc352c41767d15f80a4d322198355af1f032581da3ea910160405180910390a35050600160005550565b60006002600054036107dd5760405162461bcd60e51b81526004016101dc90611537565b6002600055816108245760405162461bcd60e51b81526020600482015260126024820152710536861726573206d757374206265203e20360741b60448201526064016101dc565b336000908152600360205260409020548281101561087a5760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e742073686172657360681b60448201526064016101dc565b61088383610b52565b9150600082116108cb5760405162461bcd60e51b81526020600482015260136024820152720576974686472617720616d6f756e74203d203606c1b60448201526064016101dc565b6108d3610b0c565b8211156109225760405162461bcd60e51b815260206004820152601b60248201527f496e73756666696369656e742066726565206c6971756964697479000000000060448201526064016101dc565b61092c83826114e1565b33600090815260036020526040812091909155600280548592906109519084906114e1565b90915550610961905033836112c8565b604080518381526020810185905233917f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6910160405180910390a2506001600055919050565b6002600054036109c95760405162461bcd60e51b81526004016101dc90611537565b600260009081558181526006602052604090205480610a1c5760405162461bcd60e51b815260206004820152600f60248201526e155b9adb9bdddb8818da185b9b995b608a1b60448201526064016101dc565b6000828152600760205260409020544211610a675760405162461bcd60e51b815260206004820152600b60248201526a139bdd08195e1c1a5c995960aa1b60448201526064016101dc565b8060056000828254610a7991906114e1565b925050819055508060086000828254610a9291906114e1565b9091555050600082815260066020908152604080832083905560078252808320929092559051828152339184917fd2c7b3913e8b91b3448548f3fc1965a5455e0fa256af80705cab84affce61fa6910160405180910390a350506001600055565b6000610b0682610b01610527565b6113e6565b92915050565b60006005544761054491906114e1565b6001546001600160a01b03163314610b465760405162461bcd60e51b81526004016101dc9061156e565b610b506000611415565b565b60006002546001610b6391906114f4565b610b6b610527565b610b769060016114f4565b610b8090846115a3565b610b0691906115ba565b600081815260066020526040812054612710610bc67f0000000000000000000000000000000000000000000000000000000000000000836115a3565b610bd091906115ba565b610bda90826114f4565b9392505050565b6001546001600160a01b03163314610c0b5760405162461bcd60e51b81526004016101dc9061156e565b6001600160a01b038116610c545760405162461bcd60e51b815260206004820152601060248201526f24b73b30b634b21037b832b930ba37b960811b60448201526064016101dc565b600480546001600160a01b038381166001600160a01b0319831681179093556040519116919082907ffbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad0390600090a35050565b6004546001600160a01b03163314610cd05760405162461bcd60e51b81526004016101dc90611507565b60003411610d135760405162461bcd60e51b815260206004820152601060248201526f0426f6e64206d757374206265203e20360841b60448201526064016101dc565b3460056000828254610d2591906114f4565b909155505060055460405133917fc4e1648d08804ab9c896cf4e8349ac6b9d4849cbeffe50bb73a3b9d7970a8c5491610d6691348252602082015260400190565b60405180910390a2565b6004546001600160a01b03163314610d9a5760405162461bcd60e51b81526004016101dc90611507565b600260005403610dbc5760405162461bcd60e51b81526004016101dc90611537565b600260005580610e035760405162461bcd60e51b81526020600482015260126024820152710416d6f756e74206d757374206265203e20360741b60448201526064016101dc565b600554811115610e495760405162461bcd60e51b8152602060048201526011602482015270125b9cdd59999a58da595b9d08189bdb99607a1b60448201526064016101dc565b60085481600554610e5a91906114e1565b1015610e9e5760405162461bcd60e51b8152602060048201526013602482015272426f6e642062656c6f7720636f76657261676560681b60448201526064016101dc565b8060056000828254610eb091906114e1565b9091555050600454610ecb906001600160a01b0316826112c8565b60055460405133917ff0566a10f495c405474bbbfa43e7d77dffab0bdbf3240d1880b2b0624968540991610f0791858252602082015260400190565b60405180910390a2506001600055565b6000600260005403610f3b5760405162461bcd60e51b81526004016101dc90611537565b60026000553480610f845760405162461bcd60e51b815260206004820152601360248201527204465706f736974206d757374206265203e203606c1b60448201526064016101dc565b610f9b8182610f91610527565b610b0191906114e1565b915060008211610fe15760405162461bcd60e51b815260206004820152601160248201527004d696e74656420736861726573203d203607c1b60448201526064016101dc565b33600090815260036020526040812080548492906110009084906114f4565b92505081905550816002600082825461101991906114f4565b9091555050604080518281526020810184905233917f73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca910160405180910390a250600160005590565b6001546001600160a01b0316331461108c5760405162461bcd60e51b81526004016101dc9061156e565b6001600160a01b0381166110f15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101dc565b6110fa81611415565b50565b6004546001600160a01b031633146111275760405162461bcd60e51b81526004016101dc90611507565b6002600054036111495760405162461bcd60e51b81526004016101dc90611537565b60026000908155818152600660205260409020548061119c5760405162461bcd60e51b815260206004820152600f60248201526e155b9adb9bdddb8818da185b9b995b608a1b60448201526064016101dc565b6127106111c97f0000000000000000000000000000000000000000000000000000000000000000836115a3565b6111d391906115ba565b6111dd90826114f4565b3410156112405760405162461bcd60e51b815260206004820152602b60248201527f52657475726e6564204554482062656c6f77207072696e636970616c20706c7560448201526a39903332b290333637b7b960a91b60648201526084016101dc565b806008600082825461125291906114e1565b909155505060008281526006602090815260408083208390556007909152812055817fbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc782346112a182826114e1565b6040805193845260208401929092529082015260600160405180910390a250506001600055565b804710156113185760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016101dc565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611365576040519150601f19603f3d011682016040523d82523d6000602084013e61136a565b606091505b50509050806113e15760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016101dc565b505050565b60006113f38260016114f4565b6002546114019060016114f4565b61140b90856115a3565b610bda91906115ba565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000806040838503121561147a57600080fd5b50508035926020909101359150565b60006020828403121561149b57600080fd5b5035919050565b6000602082840312156114b457600080fd5b81356001600160a01b0381168114610bda57600080fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610b0657610b066114cb565b80820180821115610b0657610b066114cb565b60208082526016908201527513db9b1e481bdc195c985d1bdc8818d85b8818d85b1b60521b604082015260600190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b8082028115828204841417610b0657610b066114cb565b6000826115d757634e487b7160e01b600052601260045260246000fd5b50049056fea26469706673582212201efb57fc5831c72f739ce7138767eed04bc524ea4ef340bd1471afc71e8469bf64736f6c63430008220033",
}

// LiquidityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityPoolMetaData.ABI instead.
var LiquidityPoolABI = LiquidityPoolMetaData.ABI

// LiquidityPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidityPoolMetaData.Bin instead.
var LiquidityPoolBin = LiquidityPoolMetaData.Bin

// DeployLiquidityPool deploys a new Ethereum contract, binding an instance of LiquidityPool to it.
func DeployLiquidityPool(auth *bind.TransactOpts, backend bind.ContractBackend, _operator common.Address, _settlementWindow *big.Int, _minFeeBps *big.Int) (common.Address, *types.Transaction, *LiquidityPool, error) {
	parsed, err := LiquidityPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidityPoolBin), backend, _operator, _settlementWindow, _minFeeBps)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// LiquidityPool is an auto generated Go binding around an Ethereum contract.
type LiquidityPool struct {
	LiquidityPoolCaller     // Read-only binding to the contract
	LiquidityPoolTransactor // Write-only binding to the contract
	LiquidityPoolFilterer   // Log filterer for contract events
}

// LiquidityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityPoolSession struct {
	Contract     *LiquidityPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityPoolCallerSession struct {
	Contract *LiquidityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// LiquidityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityPoolTransactorSession struct {
	Contract     *LiquidityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// LiquidityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityPoolRaw struct {
	Contract *LiquidityPool // Generic contract binding to access the raw methods on
}

// LiquidityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityPoolCallerRaw struct {
	Contract *LiquidityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityPoolTransactorRaw struct {
	Contract *LiquidityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityPool creates a new instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPool(address common.Address, backend bind.ContractBackend) (*LiquidityPool, error) {
	contract, err := bindLiquidityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityPool{LiquidityPoolCaller: LiquidityPoolCaller{contract: contract}, LiquidityPoolTransactor: LiquidityPoolTransactor{contract: contract}, LiquidityPoolFilterer: LiquidityPoolFilterer{contract: contract}}, nil
}

// NewLiquidityPoolCaller creates a new read-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolCaller(address common.Address, caller bind.ContractCaller) (*LiquidityPoolCaller, error) {
	contract, err := bindLiquidityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolCaller{contract: contract}, nil
}

// NewLiquidityPoolTransactor creates a new write-only instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityPoolTransactor, error) {
	contract, err := bindLiquidityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolTransactor{contract: contract}, nil
}

// NewLiquidityPoolFilterer creates a new log filterer instance of LiquidityPool, bound to a specific deployed contract.
func NewLiquidityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityPoolFilterer, error) {
	contract, err := bindLiquidityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolFilterer{contract: contract}, nil
}

// bindLiquidityPool binds a generic wrapper to an already deployed contract.
func bindLiquidityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.LiquidityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.LiquidityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityPool *LiquidityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityPool *LiquidityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityPool.Contract.contract.Transact(opts, method, params...)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) BPSDENOMINATOR() (*big.Int, error) {
	return _LiquidityPool.Contract.BPSDENOMINATOR(&_LiquidityPool.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _LiquidityPool.Contract.BPSDENOMINATOR(&_LiquidityPool.CallOpts)
}

// ChannelDeadline is a free data retrieval call binding the contract method 0x5f693d0a.
//
// Solidity: function channelDeadline(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) ChannelDeadline(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "channelDeadline", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChannelDeadline is a free data retrieval call binding the contract method 0x5f693d0a.
//
// Solidity: function channelDeadline(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) ChannelDeadline(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.ChannelDeadline(&_LiquidityPool.CallOpts, arg0)
}

// ChannelDeadline is a free data retrieval call binding the contract method 0x5f693d0a.
//
// Solidity: function channelDeadline(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) ChannelDeadline(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.ChannelDeadline(&_LiquidityPool.CallOpts, arg0)
}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) LockedByChannel(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "lockedByChannel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) LockedByChannel(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.LockedByChannel(&_LiquidityPool.CallOpts, arg0)
}

// LockedByChannel is a free data retrieval call binding the contract method 0xb6cc49da.
//
// Solidity: function lockedByChannel(bytes32 ) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) LockedByChannel(arg0 [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.LockedByChannel(&_LiquidityPool.CallOpts, arg0)
}

// MinFeeBps is a free data retrieval call binding the contract method 0x071de20d.
//
// Solidity: function minFeeBps() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) MinFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "minFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinFeeBps is a free data retrieval call binding the contract method 0x071de20d.
//
// Solidity: function minFeeBps() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) MinFeeBps() (*big.Int, error) {
	return _LiquidityPool.Contract.MinFeeBps(&_LiquidityPool.CallOpts)
}

// MinFeeBps is a free data retrieval call binding the contract method 0x071de20d.
//
// Solidity: function minFeeBps() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) MinFeeBps() (*big.Int, error) {
	return _LiquidityPool.Contract.MinFeeBps(&_LiquidityPool.CallOpts)
}

// MinSettlementValue is a free data retrieval call binding the contract method 0xa517458f.
//
// Solidity: function minSettlementValue(bytes32 channelId) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) MinSettlementValue(opts *bind.CallOpts, channelId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "minSettlementValue", channelId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSettlementValue is a free data retrieval call binding the contract method 0xa517458f.
//
// Solidity: function minSettlementValue(bytes32 channelId) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) MinSettlementValue(channelId [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.MinSettlementValue(&_LiquidityPool.CallOpts, channelId)
}

// MinSettlementValue is a free data retrieval call binding the contract method 0xa517458f.
//
// Solidity: function minSettlementValue(bytes32 channelId) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) MinSettlementValue(channelId [32]byte) (*big.Int, error) {
	return _LiquidityPool.Contract.MinSettlementValue(&_LiquidityPool.CallOpts, channelId)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) Operator() (common.Address, error) {
	return _LiquidityPool.Contract.Operator(&_LiquidityPool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) Operator() (common.Address, error) {
	return _LiquidityPool.Contract.Operator(&_LiquidityPool.CallOpts)
}

// OperatorBond is a free data retrieval call binding the contract method 0x39b8b076.
//
// Solidity: function operatorBond() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) OperatorBond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "operatorBond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorBond is a free data retrieval call binding the contract method 0x39b8b076.
//
// Solidity: function operatorBond() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) OperatorBond() (*big.Int, error) {
	return _LiquidityPool.Contract.OperatorBond(&_LiquidityPool.CallOpts)
}

// OperatorBond is a free data retrieval call binding the contract method 0x39b8b076.
//
// Solidity: function operatorBond() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) OperatorBond() (*big.Int, error) {
	return _LiquidityPool.Contract.OperatorBond(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidityPool *LiquidityPoolCallerSession) Owner() (common.Address, error) {
	return _LiquidityPool.Contract.Owner(&_LiquidityPool.CallOpts)
}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) PreviewDepositShares(opts *bind.CallOpts, ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "previewDepositShares", ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) PreviewDepositShares(ethAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewDepositShares(&_LiquidityPool.CallOpts, ethAmount)
}

// PreviewDepositShares is a free data retrieval call binding the contract method 0x6574bcd4.
//
// Solidity: function previewDepositShares(uint256 ethAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) PreviewDepositShares(ethAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewDepositShares(&_LiquidityPool.CallOpts, ethAmount)
}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) PreviewWithdrawETH(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "previewWithdrawETH", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) PreviewWithdrawETH(sharesAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewWithdrawETH(&_LiquidityPool.CallOpts, sharesAmount)
}

// PreviewWithdrawETH is a free data retrieval call binding the contract method 0x78f3bfe2.
//
// Solidity: function previewWithdrawETH(uint256 sharesAmount) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) PreviewWithdrawETH(sharesAmount *big.Int) (*big.Int, error) {
	return _LiquidityPool.Contract.PreviewWithdrawETH(&_LiquidityPool.CallOpts, sharesAmount)
}

// SettlementWindow is a free data retrieval call binding the contract method 0xb4a7bdf9.
//
// Solidity: function settlementWindow() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) SettlementWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "settlementWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SettlementWindow is a free data retrieval call binding the contract method 0xb4a7bdf9.
//
// Solidity: function settlementWindow() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) SettlementWindow() (*big.Int, error) {
	return _LiquidityPool.Contract.SettlementWindow(&_LiquidityPool.CallOpts)
}

// SettlementWindow is a free data retrieval call binding the contract method 0xb4a7bdf9.
//
// Solidity: function settlementWindow() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) SettlementWindow() (*big.Int, error) {
	return _LiquidityPool.Contract.SettlementWindow(&_LiquidityPool.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) SharesOf(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "sharesOf", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) SharesOf(provider common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.SharesOf(&_LiquidityPool.CallOpts, provider)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address provider) view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) SharesOf(provider common.Address) (*big.Int, error) {
	return _LiquidityPool.Contract.SharesOf(&_LiquidityPool.CallOpts, provider)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalAssets() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalAssets(&_LiquidityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalAssets() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalAssets(&_LiquidityPool.CallOpts)
}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalLockedETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalLockedETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalLockedETH() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalLockedETH(&_LiquidityPool.CallOpts)
}

// TotalLockedETH is a free data retrieval call binding the contract method 0x8afe47cf.
//
// Solidity: function totalLockedETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalLockedETH() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalLockedETH(&_LiquidityPool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) TotalShares() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalShares(&_LiquidityPool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) TotalShares() (*big.Int, error) {
	return _LiquidityPool.Contract.TotalShares(&_LiquidityPool.CallOpts)
}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCaller) WithdrawableETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityPool.contract.Call(opts, &out, "withdrawableETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolSession) WithdrawableETH() (*big.Int, error) {
	return _LiquidityPool.Contract.WithdrawableETH(&_LiquidityPool.CallOpts)
}

// WithdrawableETH is a free data retrieval call binding the contract method 0x6913e7cc.
//
// Solidity: function withdrawableETH() view returns(uint256)
func (_LiquidityPool *LiquidityPoolCallerSession) WithdrawableETH() (*big.Int, error) {
	return _LiquidityPool.Contract.WithdrawableETH(&_LiquidityPool.CallOpts)
}

// BondETH is a paid mutator transaction binding the contract method 0xb453de09.
//
// Solidity: function bondETH() payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) BondETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "bondETH")
}

// BondETH is a paid mutator transaction binding the contract method 0xb453de09.
//
// Solidity: function bondETH() payable returns()
func (_LiquidityPool *LiquidityPoolSession) BondETH() (*types.Transaction, error) {
	return _LiquidityPool.Contract.BondETH(&_LiquidityPool.TransactOpts)
}

// BondETH is a paid mutator transaction binding the contract method 0xb453de09.
//
// Solidity: function bondETH() payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) BondETH() (*types.Transaction, error) {
	return _LiquidityPool.Contract.BondETH(&_LiquidityPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolSession) Deposit() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns(uint256 mintedShares)
func (_LiquidityPool *LiquidityPoolTransactorSession) Deposit() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Deposit(&_LiquidityPool.TransactOpts)
}

// ExpireChannel is a paid mutator transaction binding the contract method 0x42161b74.
//
// Solidity: function expireChannel(bytes32 channelId) returns()
func (_LiquidityPool *LiquidityPoolTransactor) ExpireChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "expireChannel", channelId)
}

// ExpireChannel is a paid mutator transaction binding the contract method 0x42161b74.
//
// Solidity: function expireChannel(bytes32 channelId) returns()
func (_LiquidityPool *LiquidityPoolSession) ExpireChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ExpireChannel(&_LiquidityPool.TransactOpts, channelId)
}

// ExpireChannel is a paid mutator transaction binding the contract method 0x42161b74.
//
// Solidity: function expireChannel(bytes32 channelId) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) ExpireChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.ExpireChannel(&_LiquidityPool.TransactOpts, channelId)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) FundChannel(opts *bind.TransactOpts, channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "fundChannel", channelId, amount)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) FundChannel(channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.FundChannel(&_LiquidityPool.TransactOpts, channelId, amount)
}

// FundChannel is a paid mutator transaction binding the contract method 0x260b6086.
//
// Solidity: function fundChannel(bytes32 channelId, uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) FundChannel(channelId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.FundChannel(&_LiquidityPool.TransactOpts, channelId, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidityPool.Contract.RenounceOwnership(&_LiquidityPool.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolTransactor) SetOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "setOperator", newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetOperator(&_LiquidityPool.TransactOpts, newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address newOperator) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SetOperator(newOperator common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SetOperator(&_LiquidityPool.TransactOpts, newOperator)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) SettleChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "settleChannel", channelId)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolSession) SettleChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SettleChannel(&_LiquidityPool.TransactOpts, channelId)
}

// SettleChannel is a paid mutator transaction binding the contract method 0xfd3d3199.
//
// Solidity: function settleChannel(bytes32 channelId) payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) SettleChannel(channelId [32]byte) (*types.Transaction, error) {
	return _LiquidityPool.Contract.SettleChannel(&_LiquidityPool.TransactOpts, channelId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidityPool.Contract.TransferOwnership(&_LiquidityPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolTransactor) Withdraw(opts *bind.TransactOpts, sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdraw", sharesToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolSession) Withdraw(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, sharesToBurn)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 sharesToBurn) returns(uint256 ethAmountOut)
func (_LiquidityPool *LiquidityPoolTransactorSession) Withdraw(sharesToBurn *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.Withdraw(&_LiquidityPool.TransactOpts, sharesToBurn)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactor) WithdrawBond(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.contract.Transact(opts, "withdrawBond", amount)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolSession) WithdrawBond(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawBond(&_LiquidityPool.TransactOpts, amount)
}

// WithdrawBond is a paid mutator transaction binding the contract method 0xc3daab96.
//
// Solidity: function withdrawBond(uint256 amount) returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) WithdrawBond(amount *big.Int) (*types.Transaction, error) {
	return _LiquidityPool.Contract.WithdrawBond(&_LiquidityPool.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquidityPool *LiquidityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquidityPool.Contract.Receive(&_LiquidityPool.TransactOpts)
}

// LiquidityPoolBondPostedIterator is returned from FilterBondPosted and is used to iterate over the raw logs and unpacked data for BondPosted events raised by the LiquidityPool contract.
type LiquidityPoolBondPostedIterator struct {
	Event *LiquidityPoolBondPosted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolBondPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolBondPosted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolBondPosted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolBondPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolBondPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolBondPosted represents a BondPosted event raised by the LiquidityPool contract.
type LiquidityPoolBondPosted struct {
	Operator common.Address
	Amount   *big.Int
	NewBond  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondPosted is a free log retrieval operation binding the contract event 0xc4e1648d08804ab9c896cf4e8349ac6b9d4849cbeffe50bb73a3b9d7970a8c54.
//
// Solidity: event BondPosted(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) FilterBondPosted(opts *bind.FilterOpts, operator []common.Address) (*LiquidityPoolBondPostedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "BondPosted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolBondPostedIterator{contract: _LiquidityPool.contract, event: "BondPosted", logs: logs, sub: sub}, nil
}

// WatchBondPosted is a free log subscription operation binding the contract event 0xc4e1648d08804ab9c896cf4e8349ac6b9d4849cbeffe50bb73a3b9d7970a8c54.
//
// Solidity: event BondPosted(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) WatchBondPosted(opts *bind.WatchOpts, sink chan<- *LiquidityPoolBondPosted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "BondPosted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolBondPosted)
				if err := _LiquidityPool.contract.UnpackLog(event, "BondPosted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondPosted is a log parse operation binding the contract event 0xc4e1648d08804ab9c896cf4e8349ac6b9d4849cbeffe50bb73a3b9d7970a8c54.
//
// Solidity: event BondPosted(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) ParseBondPosted(log types.Log) (*LiquidityPoolBondPosted, error) {
	event := new(LiquidityPoolBondPosted)
	if err := _LiquidityPool.contract.UnpackLog(event, "BondPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolBondWithdrawnIterator is returned from FilterBondWithdrawn and is used to iterate over the raw logs and unpacked data for BondWithdrawn events raised by the LiquidityPool contract.
type LiquidityPoolBondWithdrawnIterator struct {
	Event *LiquidityPoolBondWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolBondWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolBondWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolBondWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolBondWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolBondWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolBondWithdrawn represents a BondWithdrawn event raised by the LiquidityPool contract.
type LiquidityPoolBondWithdrawn struct {
	Operator common.Address
	Amount   *big.Int
	NewBond  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBondWithdrawn is a free log retrieval operation binding the contract event 0xf0566a10f495c405474bbbfa43e7d77dffab0bdbf3240d1880b2b06249685409.
//
// Solidity: event BondWithdrawn(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) FilterBondWithdrawn(opts *bind.FilterOpts, operator []common.Address) (*LiquidityPoolBondWithdrawnIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "BondWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolBondWithdrawnIterator{contract: _LiquidityPool.contract, event: "BondWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBondWithdrawn is a free log subscription operation binding the contract event 0xf0566a10f495c405474bbbfa43e7d77dffab0bdbf3240d1880b2b06249685409.
//
// Solidity: event BondWithdrawn(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) WatchBondWithdrawn(opts *bind.WatchOpts, sink chan<- *LiquidityPoolBondWithdrawn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "BondWithdrawn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolBondWithdrawn)
				if err := _LiquidityPool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBondWithdrawn is a log parse operation binding the contract event 0xf0566a10f495c405474bbbfa43e7d77dffab0bdbf3240d1880b2b06249685409.
//
// Solidity: event BondWithdrawn(address indexed operator, uint256 amount, uint256 newBond)
func (_LiquidityPool *LiquidityPoolFilterer) ParseBondWithdrawn(log types.Log) (*LiquidityPoolBondWithdrawn, error) {
	event := new(LiquidityPoolBondWithdrawn)
	if err := _LiquidityPool.contract.UnpackLog(event, "BondWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolChannelExpiredIterator is returned from FilterChannelExpired and is used to iterate over the raw logs and unpacked data for ChannelExpired events raised by the LiquidityPool contract.
type LiquidityPoolChannelExpiredIterator struct {
	Event *LiquidityPoolChannelExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolChannelExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolChannelExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolChannelExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolChannelExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolChannelExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolChannelExpired represents a ChannelExpired event raised by the LiquidityPool contract.
type LiquidityPoolChannelExpired struct {
	ChannelId [32]byte
	Principal *big.Int
	Caller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelExpired is a free log retrieval operation binding the contract event 0xd2c7b3913e8b91b3448548f3fc1965a5455e0fa256af80705cab84affce61fa6.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId, uint256 principal, address indexed caller)
func (_LiquidityPool *LiquidityPoolFilterer) FilterChannelExpired(opts *bind.FilterOpts, channelId [][32]byte, caller []common.Address) (*LiquidityPoolChannelExpiredIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ChannelExpired", channelIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolChannelExpiredIterator{contract: _LiquidityPool.contract, event: "ChannelExpired", logs: logs, sub: sub}, nil
}

// WatchChannelExpired is a free log subscription operation binding the contract event 0xd2c7b3913e8b91b3448548f3fc1965a5455e0fa256af80705cab84affce61fa6.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId, uint256 principal, address indexed caller)
func (_LiquidityPool *LiquidityPoolFilterer) WatchChannelExpired(opts *bind.WatchOpts, sink chan<- *LiquidityPoolChannelExpired, channelId [][32]byte, caller []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ChannelExpired", channelIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolChannelExpired)
				if err := _LiquidityPool.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelExpired is a log parse operation binding the contract event 0xd2c7b3913e8b91b3448548f3fc1965a5455e0fa256af80705cab84affce61fa6.
//
// Solidity: event ChannelExpired(bytes32 indexed channelId, uint256 principal, address indexed caller)
func (_LiquidityPool *LiquidityPoolFilterer) ParseChannelExpired(log types.Log) (*LiquidityPoolChannelExpired, error) {
	event := new(LiquidityPoolChannelExpired)
	if err := _LiquidityPool.contract.UnpackLog(event, "ChannelExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolChannelFundedIterator is returned from FilterChannelFunded and is used to iterate over the raw logs and unpacked data for ChannelFunded events raised by the LiquidityPool contract.
type LiquidityPoolChannelFundedIterator struct {
	Event *LiquidityPoolChannelFunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolChannelFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolChannelFunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolChannelFunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolChannelFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolChannelFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolChannelFunded represents a ChannelFunded event raised by the LiquidityPool contract.
type LiquidityPoolChannelFunded struct {
	ChannelId [32]byte
	Principal *big.Int
	Operator  common.Address
	Deadline  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelFunded is a free log retrieval operation binding the contract event 0x921e723c5ab46e7fad6ecc352c41767d15f80a4d322198355af1f032581da3ea.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator, uint256 deadline)
func (_LiquidityPool *LiquidityPoolFilterer) FilterChannelFunded(opts *bind.FilterOpts, channelId [][32]byte, operator []common.Address) (*LiquidityPoolChannelFundedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ChannelFunded", channelIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolChannelFundedIterator{contract: _LiquidityPool.contract, event: "ChannelFunded", logs: logs, sub: sub}, nil
}

// WatchChannelFunded is a free log subscription operation binding the contract event 0x921e723c5ab46e7fad6ecc352c41767d15f80a4d322198355af1f032581da3ea.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator, uint256 deadline)
func (_LiquidityPool *LiquidityPoolFilterer) WatchChannelFunded(opts *bind.WatchOpts, sink chan<- *LiquidityPoolChannelFunded, channelId [][32]byte, operator []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ChannelFunded", channelIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolChannelFunded)
				if err := _LiquidityPool.contract.UnpackLog(event, "ChannelFunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelFunded is a log parse operation binding the contract event 0x921e723c5ab46e7fad6ecc352c41767d15f80a4d322198355af1f032581da3ea.
//
// Solidity: event ChannelFunded(bytes32 indexed channelId, uint256 principal, address indexed operator, uint256 deadline)
func (_LiquidityPool *LiquidityPoolFilterer) ParseChannelFunded(log types.Log) (*LiquidityPoolChannelFunded, error) {
	event := new(LiquidityPoolChannelFunded)
	if err := _LiquidityPool.contract.UnpackLog(event, "ChannelFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolChannelSettledIterator is returned from FilterChannelSettled and is used to iterate over the raw logs and unpacked data for ChannelSettled events raised by the LiquidityPool contract.
type LiquidityPoolChannelSettledIterator struct {
	Event *LiquidityPoolChannelSettled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolChannelSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolChannelSettled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolChannelSettled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolChannelSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolChannelSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolChannelSettled represents a ChannelSettled event raised by the LiquidityPool contract.
type LiquidityPoolChannelSettled struct {
	ChannelId     [32]byte
	Principal     *big.Int
	TotalReturned *big.Int
	FeeGain       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChannelSettled is a free log retrieval operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) FilterChannelSettled(opts *bind.FilterOpts, channelId [][32]byte) (*LiquidityPoolChannelSettledIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "ChannelSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolChannelSettledIterator{contract: _LiquidityPool.contract, event: "ChannelSettled", logs: logs, sub: sub}, nil
}

// WatchChannelSettled is a free log subscription operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) WatchChannelSettled(opts *bind.WatchOpts, sink chan<- *LiquidityPoolChannelSettled, channelId [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "ChannelSettled", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolChannelSettled)
				if err := _LiquidityPool.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChannelSettled is a log parse operation binding the contract event 0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7.
//
// Solidity: event ChannelSettled(bytes32 indexed channelId, uint256 principal, uint256 totalReturned, uint256 feeGain)
func (_LiquidityPool *LiquidityPoolFilterer) ParseChannelSettled(log types.Log) (*LiquidityPoolChannelSettled, error) {
	event := new(LiquidityPoolChannelSettled)
	if err := _LiquidityPool.contract.UnpackLog(event, "ChannelSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the LiquidityPool contract.
type LiquidityPoolDepositedIterator struct {
	Event *LiquidityPoolDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolDeposited represents a Deposited event raised by the LiquidityPool contract.
type LiquidityPoolDeposited struct {
	Provider     common.Address
	EthAmount    *big.Int
	SharesMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) FilterDeposited(opts *bind.FilterOpts, provider []common.Address) (*LiquidityPoolDepositedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Deposited", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolDepositedIterator{contract: _LiquidityPool.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *LiquidityPoolDeposited, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Deposited", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolDeposited)
				if err := _LiquidityPool.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed provider, uint256 ethAmount, uint256 sharesMinted)
func (_LiquidityPool *LiquidityPoolFilterer) ParseDeposited(log types.Log) (*LiquidityPoolDeposited, error) {
	event := new(LiquidityPoolDeposited)
	if err := _LiquidityPool.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the LiquidityPool contract.
type LiquidityPoolOperatorUpdatedIterator struct {
	Event *LiquidityPoolOperatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolOperatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolOperatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolOperatorUpdated represents a OperatorUpdated event raised by the LiquidityPool contract.
type LiquidityPoolOperatorUpdated struct {
	PreviousOperator common.Address
	NewOperator      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, previousOperator []common.Address, newOperator []common.Address) (*LiquidityPoolOperatorUpdatedIterator, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "OperatorUpdated", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolOperatorUpdatedIterator{contract: _LiquidityPool.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *LiquidityPoolOperatorUpdated, previousOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "OperatorUpdated", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolOperatorUpdated)
				if err := _LiquidityPool.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUpdated is a log parse operation binding the contract event 0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03.
//
// Solidity: event OperatorUpdated(address indexed previousOperator, address indexed newOperator)
func (_LiquidityPool *LiquidityPoolFilterer) ParseOperatorUpdated(log types.Log) (*LiquidityPoolOperatorUpdated, error) {
	event := new(LiquidityPoolOperatorUpdated)
	if err := _LiquidityPool.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferredIterator struct {
	Event *LiquidityPoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidityPool contract.
type LiquidityPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolOwnershipTransferredIterator{contract: _LiquidityPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolOwnershipTransferred)
				if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidityPool *LiquidityPoolFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityPoolOwnershipTransferred, error) {
	event := new(LiquidityPoolOwnershipTransferred)
	if err := _LiquidityPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityPoolWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the LiquidityPool contract.
type LiquidityPoolWithdrawnIterator struct {
	Event *LiquidityPoolWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityPoolWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityPoolWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityPoolWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityPoolWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityPoolWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityPoolWithdrawn represents a Withdrawn event raised by the LiquidityPool contract.
type LiquidityPoolWithdrawn struct {
	Provider     common.Address
	EthAmount    *big.Int
	SharesBurned *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) FilterWithdrawn(opts *bind.FilterOpts, provider []common.Address) (*LiquidityPoolWithdrawnIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.FilterLogs(opts, "Withdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityPoolWithdrawnIterator{contract: _LiquidityPool.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LiquidityPoolWithdrawn, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityPool.contract.WatchLogs(opts, "Withdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityPoolWithdrawn)
				if err := _LiquidityPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed provider, uint256 ethAmount, uint256 sharesBurned)
func (_LiquidityPool *LiquidityPoolFilterer) ParseWithdrawn(log types.Log) (*LiquidityPoolWithdrawn, error) {
	event := new(LiquidityPoolWithdrawn)
	if err := _LiquidityPool.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
