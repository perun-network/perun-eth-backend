package channel

import (
	"context"

	"github.com/pkg/errors"
	"perun.network/go-perun/channel"
	"perun.network/go-perun/wallet"
)

type CoordinatorBackend interface {
	Coordinate(
		ctx context.Context,
		req channel.AdjudicatorReq,
		subStates []channel.SignedState, // includes sub-channels and virtual channels
		coordSigs []wallet.Sig, // depth-first ordered coordinator signatures
	) error
}

// Ensure Adjudicator implements CoordinatorBackend.
var _ CoordinatorBackend = (*Adjudicator)(nil)

// Coordinate coordinates the state of the channel directly on the blockchain.
func (a *Adjudicator) Coordinate(
	ctx context.Context,
	req channel.AdjudicatorReq,
	subChannels []channel.SignedState,
	coordSigs []wallet.Sig) error {
	// In the case of final states, we already call concludeFinal on the
	// adjudicator. Method ensureCoordinated calls concludeFinal for final states.
	if err := a.ensureCoordinated(ctx, req, subChannels, coordSigs); err != nil {
		return errors.WithMessage(err, "ensuring Coordinated")
	}

	return nil
}

func needCoordinate(ctx context.Context, req channel.AdjudicatorReq) bool {
	if !channel.IsNoApp(req.Params.App) && false /* !channel.HasCoordinator(req.Params)*/ {
		return true
	}
	return false
}
