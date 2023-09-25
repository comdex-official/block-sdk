package block

import (
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmempool "github.com/cosmos/cosmos-sdk/types/mempool"
	"github.com/skip-mev/block-sdk/block/proposals"
)

// LaneMempool defines the interface a lane's mempool should implement. The basic API
// is the same as the sdk.Mempool, but it also includes a Compare function that is used
// to determine the relative priority of two transactions belonging in the same lane.
//
//go:generate mockery --name LaneMempool --output ./mocks --outpkg mocks --case underscore
type LaneMempool interface {
	sdkmempool.Mempool

	// Compare determines the relative priority of two transactions belonging in the same lane. Compare
	// will return -1 if this transaction has a lower priority than the other transaction, 0 if they have
	// the same priority, and 1 if this transaction has a higher priority than the other transaction.
	Compare(ctx sdk.Context, this, other sdk.Tx) int

	// Contains returns true if the transaction is contained in the mempool.
	Contains(tx sdk.Tx) bool
}

// Lane defines an interface used for matching transactions to lanes, storing transactions,
// and constructing partial blocks.
//
//go:generate mockery --name Lane --output ./mocks --outpkg mocks --case underscore
type Lane interface {
	LaneMempool

	// PrepareLane builds a portion of the block. It inputs the lane limits - maximum number of bytes and
	// gas - that can be consumed by a given lane, the partial proposal, and a function to call the next
	// lane in the chain. The next lane in the chain will be called with the updated proposal and context.
	PrepareLane(
		ctx sdk.Context,
		proposal proposals.Proposal,
		limit proposals.LaneLimits,
		next PrepareLanesHandler,
	) (proposals.Proposal, error)

	// ProcessLane verifies this lane's portion of a proposed block. It inputs the transactions
	// that may belong to this lane and a function to call the next lane in the chain. The next
	// lane in the chain will be called with the updated context and filtered down transactions.
	ProcessLane(
		ctx sdk.Context,
		proposal proposals.Proposal,
		partialProposal [][]byte,
		next ProcessLanesHandler,
	) (proposals.Proposal, error)

	// GetMaxBlockSpace returns the max block space for the lane as a relative percentage.
	GetMaxBlockSpace() math.LegacyDec

	// Logger returns the lane's logger.
	Logger() log.Logger

	// Name returns the name of the lane.
	Name() string

	// SetAnteHandler sets the lane's antehandler.
	SetAnteHandler(antehander sdk.AnteHandler)

	// SetIgnoreList sets the lanes that should be ignored by this lane.
	SetIgnoreList(ignoreList []Lane)

	// Match determines if a transaction belongs to this lane.
	Match(ctx sdk.Context, tx sdk.Tx) bool
}
