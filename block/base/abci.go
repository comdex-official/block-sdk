package base

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/block-sdk/block"
	"github.com/skip-mev/block-sdk/block/proposals"
	"github.com/skip-mev/block-sdk/block/utils"
)

// PrepareLane will prepare a partial proposal for the lane. It will select transactions from the
// lane respecting the selection logic of the prepareLaneHandler. It will then update the partial
// proposal with the selected transactions. If the proposal is unable to be updated, we return an
// error. The proposal will only be modified if it passes all of the invarient checks.
func (l *BaseLane) PrepareLane(
	ctx sdk.Context,
	proposal proposals.Proposal,
	next block.PrepareLanesHandler,
) (proposals.Proposal, error) {
	limit := proposal.GetLaneLimits(l.cfg.MaxBlockSpace)

	// Get the partial block constraints for this lane and create a partial proposal.
	txsToInclude, txsToRemove, err := l.prepareLaneHandler(ctx, proposal, limit)
	if err != nil {
		return proposal, err
	}

	// Remove all transactions that were invalid during the creation of the partial proposal.
	if err := utils.RemoveTxsFromLane(txsToRemove, l); err != nil {
		l.Logger().Error(
			"failed to remove transactions from lane",
			"lane", l.Name(),
			"err", err,
		)
	}

	// Update the proposal with the selected transactions. This fails if the lane attempted to add
	// more transactions than the max block space for the lane.
	if err := proposal.UpdateProposal(l.Name(), txsToInclude, limit); err != nil {
		return proposal, err
	}

	return next(ctx, proposal)
}

// ProcessLane verifies that the transactions included in the block proposal are valid respecting
// the verification logic of the lane (processLaneHandler). If the transactions are valid, we
// return the transactions that do not belong to this lane to the next lane. If the transactions
// are invalid, we return an error.
func (l *BaseLane) ProcessLane(
	ctx sdk.Context,
	proposal proposals.Proposal,
	txs [][]byte,
	next block.ProcessLanesHandler,
) (proposals.Proposal, error) {
	// Assume that this lane is processing sdk.Tx's and decode the transactions.
	decodedTxs, err := utils.GetDecodedTxs(l.TxDecoder(), txs)
	if err != nil {
		return proposal, err
	}

	limit := proposal.GetLaneLimits(l.cfg.MaxBlockSpace)

	// Optimistically update the proposal with the partial proposal. This fails if the lane attempted to add more
	// transactions than the allocated max block space for the lane.
	if err := proposal.UpdateProposal(l.Name(), decodedTxs, limit); err != nil {
		return proposal, err
	}

	// Verify the transactions that belong to this lane according to the verification logic of the lane.
	if err := l.processLaneHandler(ctx, decodedTxs); err != nil {
		return proposal, err
	}

	return next(ctx, proposal)
}

// AnteVerifyTx verifies that the transaction is valid respecting the ante verification logic of
// of the antehandler chain.
func (l *BaseLane) AnteVerifyTx(ctx sdk.Context, tx sdk.Tx, simulate bool) (sdk.Context, error) {
	if l.cfg.AnteHandler != nil {
		return l.cfg.AnteHandler(ctx, tx, simulate)
	}

	return ctx, nil
}
