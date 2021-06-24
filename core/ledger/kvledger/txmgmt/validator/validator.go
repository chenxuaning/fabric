/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package validator

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/txmgr"
)

// Validator validates the transactions present in a block and returns a batch that should be used to update the state
type Validator interface {
	ValidateAndPrepareBatch(blockAndPvtdata *ledger.BlockAndPvtData, doMVCCValidation bool) (
		*privacyenabledstate.UpdateBatch, []*txmgr.TxStatInfo, error,
	)
}

// ErrPvtdataHashMissmatch is to be thrown if the hash of a collection present in the public read-write set
// does not match with the corresponding pvt data  supplied with the block for validation
type ErrPvtdataHashMissmatch struct {
	Msg string
}

func (e *ErrPvtdataHashMissmatch) Error() string {
	return e.Msg
}
