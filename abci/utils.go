package abci

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/skip-mev/block-sdk/block"
	"github.com/skip-mev/block-sdk/block/proposals"
	"github.com/skip-mev/block-sdk/lanes/terminator"
)

const (
	// MaxUint64 is the maximum value of a uint64.
	MaxUint64 = 1<<64 - 1
)

// ChainPrepareLanes chains together the proposal preparation logic from each lane
// into a single function. The first lane in the chain is the first lane to be prepared and
// the last lane in the chain is the last lane to be prepared.
//
// In the case where any of the lanes fail to prepare the partial proposal, the lane that failed
// will be skipped and the next lane in the chain will be called to prepare the proposal.
func ChainPrepareLanes(chain []block.Lane) block.PrepareLanesHandler {
	if len(chain) == 0 {
		return nil
	}

	// Handle non-terminated decorators chain
	if (chain[len(chain)-1] != terminator.Terminator{}) {
		chain = append(chain, terminator.Terminator{})
	}

	return func(ctx sdk.Context, partialProposal proposals.Proposal) (finalProposal proposals.Proposal, err error) {
		lane := chain[0]
		lane.Logger().Info("preparing lane", "lane", lane.Name())

		// Cache the context in the case where any of the lanes fail to prepare the proposal.
		cacheCtx, write := ctx.CacheContext()

		// We utilize a recover to handle any panics or errors that occur during the preparation
		// of a lane's transactions. This defer will first check if there was a panic or error
		// thrown from the lane's preparation logic. If there was, we log the error, skip the lane,
		// and call the next lane in the chain to the prepare the proposal.
		defer func() {
			if rec := recover(); rec != nil || err != nil {
				lane.Logger().Error("failed to prepare lane", "lane", lane.Name(), "err", err, "recover_error", rec)
				lane.Logger().Info("skipping lane", "lane", lane.Name())

				if len(chain) <= 2 {
					// If there are only two lanes remaining, then the first lane in the chain
					// is the lane that failed to prepare the partial proposal and the second lane in the
					// chain is the terminator lane. We return the proposal as is.
					finalProposal, err = partialProposal, nil
				} else {
					// If there are more than two lanes remaining, then the first lane in the chain
					// is the lane that failed to prepare the proposal but the second lane in the
					// chain is not the terminator lane so there could potentially be more transactions
					// added to the proposal
					finalProposal, err = ChainPrepareLanes(chain[1:])(ctx, partialProposal)
				}
			} else {
				// Write the cache to the context since we know that the lane successfully prepared
				// the partial proposal. State is written to in a backwards, cascading fashion. This means
				// that the final context will only be updated after all other lanes have successfully
				// prepared the partial proposal.
				write()
			}
		}()

		// Get the maximum number of bytes that can be included in the proposal for this lane.
		metaData := partialProposal.GetMetaData()
		limit := proposals.GetLaneLimits(
			partialProposal.GetMaxTxBytes(), metaData.TotalTxBytes,
			partialProposal.GetMaxGasLimit(), metaData.TotalGasLimit,
			lane.GetMaxBlockSpace(),
		)

		return lane.PrepareLane(
			cacheCtx,
			partialProposal,
			limit,
			ChainPrepareLanes(chain[1:]),
		)
	}
}

// ChainProcessLanes chains together the proposal verification logic from each lane
// into a single function. The first lane in the chain is the first lane to be verified and
// the last lane in the chain is the last lane to be verified.
func ChainProcessLanes(partialProposals [][][]byte, chain []block.Lane) block.ProcessLanesHandler {
	if len(chain) == 0 {
		return nil
	}

	// Handle non-terminated decorators chain
	if (chain[len(chain)-1] != terminator.Terminator{}) {
		chain = append(chain, terminator.Terminator{})
		partialProposals = append(partialProposals, nil)
	}

	return func(ctx sdk.Context, proposal proposals.Proposal) (proposals.Proposal, error) {
		lane := chain[0]
		partialProposal := partialProposals[0]

		lane.Logger().Info(
			"processing lane",
			"lane", chain[0].Name(),
			"num_txs", len(partialProposal),
		)

		return lane.ProcessLane(ctx, proposal, partialProposal, ChainProcessLanes(partialProposals[1:], chain[1:]))
	}
}

func getBlockLimits(ctx sdk.Context) (int64, uint64) {
	blockParams := ctx.ConsensusParams().Block

	// If the max gas is set to 0, then the max gas limit for the block can be infinite.
	// Otherwise we use the max gas limit casted as a uint64 which is how gas limits are
	// extracted from sdk.Tx's.
	var maxGasLimit uint64
	if maxGas := blockParams.MaxGas; maxGas > 0 {
		maxGasLimit = uint64(maxGas)
	} else {
		maxGasLimit = MaxUint64
	}

	return blockParams.MaxBytes, maxGasLimit
}
