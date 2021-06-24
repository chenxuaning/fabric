/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kvledger

import "github.com/hyperledger/fabric/core/ledger"

type recoverable interface {
	// ShouldRecover return whether recovery is need.
	// If the recovery is needed, this method also returns the block number to start recovery from.
	// lastAvailableBlock is the max block number that has been committed to the block storage
	ShouldRecover(lastAvailableBlock uint64) (bool, uint64, error)
	// CommitLostBlock recommits the block
	CommitLostBlock(block *ledger.BlockAndPvtData) error
	// Name returns the name of the database: either state or history
	Name() string
}

type recoverer struct {
	firstBlockNum uint64
	recoverable   recoverable
}
