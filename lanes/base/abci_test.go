package base_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/block-sdk/block"
	"github.com/skip-mev/block-sdk/block/base"
	"github.com/skip-mev/block-sdk/block/proposals"
	"github.com/skip-mev/block-sdk/block/utils"
	defaultlane "github.com/skip-mev/block-sdk/lanes/base"
	testutils "github.com/skip-mev/block-sdk/testutils"
)

func (s *BaseTestSuite) TestPrepareLane() {
	s.Run("should not build a proposal when amount configured to lane is too small", func() {
		// Create a basic transaction that should not in the proposal
		tx, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx: true,
		}
		lane := s.initLane(expectedExecution)

		s.Require().NoError(lane.Insert(sdk.Context{}, tx))

		txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz) - 1),
			MaxGas:     10000000000000,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is empty
		s.Require().Equal(0, stats.NumTxs)
		s.Require().Equal(int64(0), stats.TotalGasLimit)
		s.Require().Equal(uint64(0), stats.TotalGasLimit)
	})

	s.Run("should not build a proposal when gas configured to lane is too small", func() {
		// Create a basic transaction that should not in the proposal
		tx, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx: true,
		}
		lane := s.initLane(expectedExecution)

		// Insert the transaction into the lane
		s.Require().NoError(lane.Insert(sdk.Context{}, tx))

		txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz)),
			MaxGas:     9,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is empty
		s.Require().Equal(0, stats.NumTxs)
		s.Require().Equal(int64(0), stats.TotalTxBytes)
		s.Require().Equal(uint64(0), stats.TotalGasLimit)
	})

	s.Run("should not build a proposal when gas configured to lane is too small p2", func() {
		// Create a basic transaction that should not in the proposal
		tx, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx: true,
		}
		lane := s.initLane(expectedExecution)

		// Insert the transaction into the lane
		s.Require().NoError(lane.Insert(sdk.Context{}, tx))

		txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		// Create a proposal
		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz)) * 10, // have enough space for 10 of these txs
			MaxGas:     9,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is empty
		s.Require().Equal(0, stats.NumTxs)
		s.Require().Equal(int64(0), stats.TotalTxBytes)
		s.Require().Equal(uint64(0), stats.TotalGasLimit)
	})

	s.Run("should be able to build a proposal with a tx that just fits in", func() {
		// Create a basic transaction that should just fit in with the gas limit
		// and max size
		tx, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx: true,
		}
		lane := s.initLane(expectedExecution)

		s.Require().NoError(lane.Insert(sdk.Context{}, tx))

		txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz)),
			MaxGas:     10,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is not empty and contains the transaction
		s.Require().Equal(1, stats.NumTxs)
		s.Require().Equal(limit.MaxTxBytes, stats.TotalTxBytes)
		s.Require().Equal(uint64(10), stats.TotalGasLimit)
		s.Require().Equal(txBz, finalProposal.GetTxs()[0])
	})

	s.Run("should not build a proposal with a that fails verify tx", func() {
		// Create a basic transaction that should not in the proposal
		tx, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		// We expect the transaction to fail verify tx
		expectedExecution := map[sdk.Tx]bool{
			tx: false,
		}
		lane := s.initLane(expectedExecution)

		s.Require().NoError(lane.Insert(sdk.Context{}, tx))

		txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz)),
			MaxGas:     10,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is empty
		s.Require().Equal(0, stats.NumTxs)
		s.Require().Equal(int64(0), stats.TotalTxBytes)
		s.Require().Equal(uint64(0), stats.TotalGasLimit)

		// Ensure the transaction is removed from the lane
		s.Require().False(lane.Contains(tx))
		s.Require().Equal(0, lane.CountTx())
	})

	s.Run("should order transactions correctly in the proposal", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(2)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			10,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		}
		lane := s.initLane(expectedExecution)

		s.Require().NoError(lane.Insert(sdk.Context{}, tx1))
		s.Require().NoError(lane.Insert(sdk.Context{}, tx2))

		txBz1, err := s.encodingConfig.TxConfig.TxEncoder()(tx1)
		s.Require().NoError(err)

		txBz2, err := s.encodingConfig.TxConfig.TxEncoder()(tx2)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz1)) + int64(len(txBz2)),
			MaxGas:     20,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is ordered correctly
		s.Require().Equal(2, stats.NumTxs)
		s.Require().Equal(limit.MaxTxBytes, stats.TotalTxBytes)
		s.Require().Equal(uint64(20), stats.TotalGasLimit)
		s.Require().Equal([][]byte{txBz1, txBz2}, finalProposal.GetTxs())
	})

	s.Run("should order transactions correctly in the proposal (with different insertion)", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(2)),
		)
		s.Require().NoError(err)

		expectedExecution := map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		}
		lane := s.initLane(expectedExecution)

		s.Require().NoError(lane.Insert(sdk.Context{}, tx1))
		s.Require().NoError(lane.Insert(sdk.Context{}, tx2))

		txBz1, err := s.encodingConfig.TxConfig.TxEncoder()(tx1)
		s.Require().NoError(err)

		txBz2, err := s.encodingConfig.TxConfig.TxEncoder()(tx2)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz1)) + int64(len(txBz2)),
			MaxGas:     2,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is ordered correctly
		s.Require().Equal(2, stats.NumTxs)
		s.Require().Equal(limit.MaxTxBytes, stats.TotalTxBytes)
		s.Require().Equal(uint64(2), stats.TotalGasLimit)
		s.Require().Equal([][]byte{txBz2, txBz1}, finalProposal.GetTxs())
	})

	s.Run("should include tx that fits in proposal when other does not", func() {
		// Create a basic transaction that should not in the proposal
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			2,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			10, // This tx is too large to fit in the proposal
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		// Create a lane with a max block space of 1 but a proposal that is smaller than the tx
		expectedExecution := map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		}
		lane := s.initLane(expectedExecution)

		// Insert the transaction into the lane
		s.Require().NoError(lane.Insert(sdk.Context{}.WithPriority(10), tx1))
		s.Require().NoError(lane.Insert(sdk.Context{}.WithPriority(5), tx2))

		txBz1, err := s.encodingConfig.TxConfig.TxEncoder()(tx1)
		s.Require().NoError(err)

		txBz2, err := s.encodingConfig.TxConfig.TxEncoder()(tx2)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz1)) + int64(len(txBz2)) - 1,
			MaxGas:     3,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is ordered correctly
		s.Require().Equal(1, stats.NumTxs)
		s.Require().Equal(int64(len(txBz1)), stats.TotalTxBytes)
		s.Require().Equal(uint64(2), stats.TotalGasLimit)
		s.Require().Equal([][]byte{txBz1}, finalProposal.GetTxs())
	})

	s.Run("should include tx that consumes all gas in proposal while other cannot", func() {
		// Create a basic transaction that should not in the proposal
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			2,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			10, // This tx is too large to fit in the proposal
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		// Create a lane with a max block space of 1 but a proposal that is smaller than the tx
		expectedExecution := map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		}
		lane := s.initLane(expectedExecution)

		// Insert the transaction into the lane
		s.Require().NoError(lane.Insert(sdk.Context{}.WithPriority(10), tx1))
		s.Require().NoError(lane.Insert(sdk.Context{}.WithPriority(5), tx2))

		txBz1, err := s.encodingConfig.TxConfig.TxEncoder()(tx1)
		s.Require().NoError(err)

		txBz2, err := s.encodingConfig.TxConfig.TxEncoder()(tx2)
		s.Require().NoError(err)

		limit := proposals.LaneLimits{
			MaxTxBytes: int64(len(txBz1)) + int64(len(txBz2)) - 1,
			MaxGas:     1,
		}
		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		finalProposal, err := lane.PrepareLane(sdk.Context{}, emptyProposal, limit, block.NoOpPrepareLanesHandler())
		s.Require().NoError(err)

		stats := finalProposal.GetMetaData()

		// Ensure the proposal is ordered correctly
		s.Require().Equal(1, stats.NumTxs)
		s.Require().Equal(int64(len(txBz2)), stats.TotalTxBytes)
		s.Require().Equal(uint64(1), stats.TotalGasLimit)
		s.Require().Equal([][]byte{txBz2}, finalProposal.GetTxs())
	})
}

func (s *BaseTestSuite) TestProcessLane() {
	s.Run("should accept a proposal with valid transactions", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
		})

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().NoError(err)
	})

	s.Run("should not accept a proposal with invalid transactions", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: false,
		})

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal with some invalid transactions", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		tx3, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[2],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
			tx2,
			tx3,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: false,
			tx3: true,
		})

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should accept proposal with transactions in correct order", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(2)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
			tx2,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		})

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().NoError(err)
	})

	s.Run("should not accept a proposal with transactions that are not in the correct order", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(2)),
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
			tx2,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		})

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal where transactions are out of order relative to other lanes", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			2,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(1)),
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			1,
			sdk.NewCoin(s.gasTokenDenom, math.NewInt(2)),
		)
		s.Require().NoError(err)

		otherLane := s.initLane(nil)

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: false,
		})
		lane.SetIgnoreList([]block.Lane{otherLane})

		proposal := []sdk.Tx{
			tx1,
			tx2,
		}

		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			100000,
			100000,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal that builds too large of a partial block", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			1,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
		})

		maxSize := s.getTxSize(tx1) - 1
		limit := proposals.NewLaneLimits(maxSize, 100000)
		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal that builds a partial block that is too gas consumptive", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
		})

		maxSize := s.getTxSize(tx1)
		limit := proposals.NewLaneLimits(maxSize, 9)
		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal that builds a partial block that is too gas consumptive p2", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			10,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
			tx2,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		})

		maxSize := s.getTxSize(tx1) + s.getTxSize(tx2)
		limit := proposals.NewLaneLimits(maxSize, 19)
		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})

	s.Run("should not accept a proposal that builds a partial block that is too large p2", func() {
		tx1, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[0],
			0,
			1,
			0,
			10,
		)
		s.Require().NoError(err)

		tx2, err := testutils.CreateRandomTx(
			s.encodingConfig.TxConfig,
			s.accounts[1],
			0,
			1,
			0,
			10,
		)
		s.Require().NoError(err)

		proposal := []sdk.Tx{
			tx1,
			tx2,
		}

		lane := s.initLane(map[sdk.Tx]bool{
			tx1: true,
			tx2: true,
		})

		maxSize := s.getTxSize(tx1) + s.getTxSize(tx2) - 1
		limit := proposals.NewLaneLimits(maxSize, 20)
		partialProposal, err := utils.GetEncodedTxs(s.encodingConfig.TxConfig.TxEncoder(), proposal)
		s.Require().NoError(err)

		emptyProposal := proposals.NewProposal(
			s.encodingConfig.TxConfig.TxEncoder(),
			limit.MaxTxBytes,
			limit.MaxGas,
		)

		_, err = lane.ProcessLane(sdk.Context{}, emptyProposal, partialProposal, block.NoOpProcessLanesHandler())
		s.Require().Error(err)
	})
}

func (s *BaseTestSuite) initLane(
	expectedExecution map[sdk.Tx]bool,
) *defaultlane.DefaultLane {
	config := base.NewLaneConfig(
		log.NewTestLogger(s.T()),
		s.encodingConfig.TxConfig.TxEncoder(),
		s.encodingConfig.TxConfig.TxDecoder(),
		s.setUpAnteHandler(expectedExecution),
		math.LegacyNewDec(1),
	)

	return defaultlane.NewDefaultLane(config)
}

func (s *BaseTestSuite) setUpAnteHandler(expectedExecution map[sdk.Tx]bool) sdk.AnteHandler {
	txCache := make(map[string]bool)
	for tx, pass := range expectedExecution {
		bz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		hash := sha256.Sum256(bz)
		hashStr := hex.EncodeToString(hash[:])
		txCache[hashStr] = pass
	}

	anteHandler := func(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, err error) {
		bz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
		s.Require().NoError(err)

		hash := sha256.Sum256(bz)
		hashStr := hex.EncodeToString(hash[:])

		pass, found := txCache[hashStr]
		if !found {
			return ctx, fmt.Errorf("tx not found")
		}

		if pass {
			return ctx, nil
		}

		return ctx, fmt.Errorf("tx failed")
	}

	return anteHandler
}

func (s *BaseTestSuite) getTxSize(tx sdk.Tx) int64 {
	txBz, err := s.encodingConfig.TxConfig.TxEncoder()(tx)
	s.Require().NoError(err)

	return int64(len(txBz))
}
