package proposals_test

import (
	"math/rand"
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/skip-mev/block-sdk/block/proposals"
	"github.com/skip-mev/block-sdk/block/utils"
	"github.com/skip-mev/block-sdk/testutils"
	"github.com/stretchr/testify/require"
)

func TestUpdateProposal(t *testing.T) {

	encodingConfig := testutils.CreateTestEncodingConfig()

	// Create a few random accounts
	random := rand.New(rand.NewSource(1))
	accounts := testutils.RandomAccounts(random, 5)

	t.Run("can update with no transactions", func(t *testing.T) {
		proposal := proposals.NewProposal(nil, 100, 100)

		limit := proposals.LaneLimits{
			MaxTxBytes:  100,
			MaxGasLimit: 100,
		}

		err := proposal.UpdateProposal("test", nil, limit)
		require.NoError(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 0, len(proposal.Txs))
		require.Equal(t, int64(0), proposal.BlockSize)
		require.Equal(t, uint64(0), proposal.GasLimt)
		require.Equal(t, 0, len(proposal.Info.Lanes))

		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 1, len(block))
	})

	t.Run("can update with a single transaction", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0])
		gasLimit := 100

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.NoError(t, err)

		// Ensure that the proposal is not empty.
		require.Equal(t, 1, len(proposal.Txs))
		require.Equal(t, int64(size), proposal.BlockSize)
		require.Equal(t, uint64(gasLimit), proposal.GasLimt)
		require.Equal(t, 1, len(proposal.Info.Lanes))
		require.Equal(t, uint64(1), proposal.Info.Lanes["test"].NumTxs)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 2, len(block))
		require.Equal(t, txBzs[0], block[1])
	})

	t.Run("can update with multiple transactions", func(t *testing.T) {
		txs := make([]sdk.Tx, 0)

		for i := 0; i < 10; i++ {
			tx, err := testutils.CreateRandomTx(
				encodingConfig.TxConfig,
				accounts[0],
				0,
				uint64(i),
				0,
				100,
			)
			require.NoError(t, err)

			txs = append(txs, tx)
		}

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), txs)
		require.NoError(t, err)

		size := 0
		gasLimit := uint64(0)

		for _, txBz := range txBzs {
			size += len(txBz)
			gasLimit += 100
		}

		// This should be just enough to store all of the txs
		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: gasLimit,
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 1000000, 10000000)

		err = proposal.UpdateProposal("test", txs, limit)
		require.NoError(t, err)

		// Ensure that the proposal is not empty.
		require.Equal(t, len(txs), len(proposal.Txs))
		require.Equal(t, int64(size), proposal.BlockSize)
		require.Equal(t, gasLimit, proposal.GasLimt)
		require.Equal(t, uint64(10), proposal.Info.Lanes["test"].NumTxs)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 11, len(block))

		for i := 0; i < 10; i++ {
			require.Equal(t, txBzs[i], block[i+1])
		}
	})

	t.Run("rejects an update with duplicate transactions", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0])
		gasLimit := 100

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.NoError(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 1, len(proposal.Txs))
		require.Equal(t, int64(size), proposal.BlockSize)
		require.Equal(t, uint64(gasLimit), proposal.GasLimt)
		require.Equal(t, 1, len(proposal.Info.Lanes))
		require.Equal(t, uint64(1), proposal.Info.Lanes["test"].NumTxs)

		// Attempt to add the same transaction again.
		err = proposal.UpdateProposal("test2", []sdk.Tx{tx}, limit)
		require.Error(t, err)

		require.Equal(t, 1, len(proposal.Txs))
		require.Equal(t, int64(size), proposal.BlockSize)
		require.Equal(t, uint64(gasLimit), proposal.GasLimt)
		require.Equal(t, 1, len(proposal.Info.Lanes))
		require.Equal(t, uint64(1), proposal.Info.Lanes["test"].NumTxs)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 2, len(block))
		require.Equal(t, txBzs[0], block[1])
	})

	t.Run("rejects an update with duplicate lane updates", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		tx2, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[1],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx, tx2})
		require.NoError(t, err)

		size := len(txBzs[0]) + len(txBzs[1])
		gasLimit := 200

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.NoError(t, err)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx2}, limit)
		require.Error(t, err)

		// Ensure that the proposal is not empty.
		require.Equal(t, 1, len(proposal.Txs))
		require.Equal(t, int64(len(txBzs[0])), proposal.BlockSize)
		require.Equal(t, uint64(100), proposal.GasLimt)
		require.Equal(t, 1, len(proposal.Info.Lanes))
		require.Equal(t, uint64(1), proposal.Info.Lanes["test"].NumTxs)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 2, len(block))
		require.Equal(t, txBzs[0], block[1])
	})

	t.Run("rejects an update where lane limit is smaller (block size)", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0]) - 1
		gasLimit := 100

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.Error(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 0, len(proposal.Txs))
		require.Equal(t, int64(0), proposal.BlockSize)
		require.Equal(t, uint64(0), proposal.GasLimt)
		require.Equal(t, 0, len(proposal.Info.Lanes))

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 1, len(block))
	})

	t.Run("rejects an update where the lane limit is smaller (gas limit)", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0])
		gasLimit := 100 - 1

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.Error(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 0, len(proposal.Txs))
		require.Equal(t, int64(0), proposal.BlockSize)
		require.Equal(t, 0, len(proposal.Info.Lanes))
		require.Equal(t, uint64(0), proposal.GasLimt)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 1, len(block))
	})

	t.Run("rejects an update where the proposal exceeds max block size", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0])
		gasLimit := 100

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), limit.MaxTxBytes-1, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.Error(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 0, len(proposal.Txs))
		require.Equal(t, int64(0), proposal.BlockSize)
		require.Equal(t, uint64(0), proposal.GasLimt)
		require.Equal(t, 0, len(proposal.Info.Lanes))

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 1, len(block))
	})

	t.Run("rejects an update where the proposal exceeds max gas limit", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx})
		require.NoError(t, err)

		size := len(txBzs[0])
		gasLimit := 100

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, limit.MaxGasLimit-1)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.Error(t, err)

		// Ensure that the proposal is empty.
		require.Equal(t, 0, len(proposal.Txs))
		require.Equal(t, int64(0), proposal.BlockSize)
		require.Equal(t, uint64(0), proposal.GasLimt)
		require.Equal(t, 0, len(proposal.Info.Lanes))

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 1, len(block))
	})

	t.Run("can add transactions from multiple lanes", func(t *testing.T) {
		tx, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[0],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		tx2, err := testutils.CreateRandomTx(
			encodingConfig.TxConfig,
			accounts[1],
			0,
			1,
			0,
			100,
		)
		require.NoError(t, err)

		txBzs, err := utils.GetEncodedTxs(encodingConfig.TxConfig.TxEncoder(), []sdk.Tx{tx, tx2})
		require.NoError(t, err)

		size := len(txBzs[0]) + len(txBzs[1])
		gasLimit := 200

		limit := proposals.LaneLimits{
			MaxTxBytes:  int64(size),
			MaxGasLimit: uint64(gasLimit),
		}

		proposal := proposals.NewProposal(encodingConfig.TxConfig.TxEncoder(), 10000, 10000)

		err = proposal.UpdateProposal("test", []sdk.Tx{tx}, limit)
		require.NoError(t, err)

		err = proposal.UpdateProposal("test2", []sdk.Tx{tx2}, limit)
		require.NoError(t, err)

		// Ensure that the proposal is not empty.
		require.Equal(t, 2, len(proposal.Txs))
		require.Equal(t, int64(size), proposal.BlockSize)
		require.Equal(t, uint64(gasLimit), proposal.GasLimt)
		require.Equal(t, 2, len(proposal.Info.Lanes))
		require.Equal(t, uint64(1), proposal.Info.Lanes["test"].NumTxs)
		require.Equal(t, uint64(1), proposal.Info.Lanes["test2"].NumTxs)

		// Ensure that the proposal can be marshalled.
		block, err := proposal.GetProposalWithInfo()
		require.NoError(t, err)
		require.Equal(t, 3, len(block))
		require.Equal(t, txBzs[0], block[1])
		require.Equal(t, txBzs[1], block[2])
	})
}

func TestGetLaneLimits(t *testing.T) {
	testCases := []struct {
		name              string
		maxTxBytes        int64
		totalTxBytesUsed  int64
		maxGasLimit       uint64
		totalGasLimitUsed uint64
		ratio             math.LegacyDec
		expectedTxBytes   int64
		expectedGasLimit  uint64
	}{
		{
			"ratio is zero",
			100,
			50,
			100,
			50,
			math.LegacyZeroDec(),
			50,
			50,
		},
		{
			"ratio is zero",
			100,
			100,
			50,
			25,
			math.LegacyZeroDec(),
			0,
			25,
		},
		{
			"ratio is zero",
			100,
			150,
			100,
			150,
			math.LegacyZeroDec(),
			0,
			0,
		},
		{
			"ratio is 1",
			100,
			0,
			75,
			0,
			math.LegacyOneDec(),
			100,
			75,
		},
		{
			"ratio is 10%",
			100,
			0,
			75,
			0,
			math.LegacyMustNewDecFromStr("0.1"),
			10,
			7,
		},
		{
			"ratio is 25%",
			100,
			0,
			80,
			0,
			math.LegacyMustNewDecFromStr("0.25"),
			25,
			20,
		},
		{
			"ratio is 50%",
			101,
			0,
			75,
			0,
			math.LegacyMustNewDecFromStr("0.5"),
			50,
			37,
		},
		{
			"ratio is 33%",
			100,
			0,
			75,
			0,
			math.LegacyMustNewDecFromStr("0.33"),
			33,
			24,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			proposal := proposals.Proposal{
				MaxBlockSize: tc.maxTxBytes,
				BlockSize:    tc.totalTxBytesUsed,
				MaxGasLimit:  tc.maxGasLimit,
				GasLimt:      tc.totalGasLimitUsed,
			}

			res := proposal.GetLaneLimits(tc.ratio)

			if res.MaxTxBytes != tc.expectedTxBytes {
				t.Errorf("expected tx bytes %d, got %d", tc.expectedTxBytes, res.MaxTxBytes)
			}

			if res.MaxGasLimit != tc.expectedGasLimit {
				t.Errorf("expected gas limit %d, got %d", tc.expectedGasLimit, res.MaxGasLimit)
			}
		})
	}
}
