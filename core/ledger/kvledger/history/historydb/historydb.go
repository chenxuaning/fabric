/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package historydb

import (
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/protos/common"
)

// HistoryDBProvider provides an instance of a history DB
type HistoryDBProvider interface {
	// GetDBHandle returns a handle to a HistoryDB
	GetDBHandle(id string) (HistoryDB, error)
	// Close closes all the HistoryDB instances and releases any resources held by HistoryDBProvider
	Close()
}

// HistoryDB - an interface that a history database should implement
type HistoryDB interface {
	NewHistoryQueryExecutor(blockStore blkstorage.BlockStore) (ledger.HistoryQueryExecutor, error)
	Commit(block *common.Block) error
	GetLastSavepoint() (*version.Height, error)
	ShouldRecover(lastAvailableBlock uint64) (bool, uint64, error)
	CommitLostBlock(blockAndPvtdata *ledger.BlockAndPvtData) error
	Name() string
}
