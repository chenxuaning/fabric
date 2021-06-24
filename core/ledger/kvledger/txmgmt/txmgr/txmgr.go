/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package txmgr

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
)

// TxMgr - an interface that a transaction manager should implement
type TxMgr interface {
	NewQueryExecutor(txid string) (ledger.QueryExecutor, error)
	NewTxSimulator(txid string) (ledger.TxSimulator, error)
	ValidateAndPrepare(blockAndPvtdata *ledger.BlockAndPvtData, doMVCCValidation bool) ([]*TxStatInfo, []byte, error)
	RemoveStaleAndCommitPvtDataOfOldBlocks(blocksPvtData map[uint64][]*ledger.TxPvtData) error
	GetLastSavepoint() (*version.Height, error)
	ShouldRecover(lastAvailableBlock uint64) (bool, uint64, error)
	CommitLostBlock(blockAndPvtdata *ledger.BlockAndPvtData) error
	Commit() error
	Rollback()
	Shutdown()
	Name() string
}

// TxStatInfo encapsulates information about a transaction
type TxStatInfo struct {
	ValidationCode peer.TxValidationCode
	TxType         common.HeaderType
	ChaincodeID    *peer.ChaincodeID
	NumCollections int
}

// ErrUnsupportedTransaction is expected to be thrown if a unsupported query is performed in an update transaction
type ErrUnsupportedTransaction struct {
	Msg string
}

func (e *ErrUnsupportedTransaction) Error() string {
	return e.Msg
}

// ErrPvtdataNotAvailable is to be thrown when an application seeks a private data item
// during simulation and the simulator is not capable of returning the version of the
// private data item consistent with the snapshopt exposed to the simulation
type ErrPvtdataNotAvailable struct {
	Msg string
}

func (e *ErrPvtdataNotAvailable) Error() string {
	return e.Msg
}
