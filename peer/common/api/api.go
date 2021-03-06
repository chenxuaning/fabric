/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"context"

	"github.com/hyperledger/fabric/peer/chaincode/api"
	cb "github.com/hyperledger/fabric/protos/common"
	ab "github.com/hyperledger/fabric/protos/orderer"
	"google.golang.org/grpc"
)

// DeliverClient defines the interface for a deliver client
type DeliverClient interface {
	Deliver(ctx context.Context, opts ...grpc.CallOption) (DeliverService, error)
}

// DeliverService defines the interface for delivering blocks
type DeliverService interface {
	Send(*cb.Envelope) error
	Recv() (*ab.DeliverResponse, error)
	CloseSend() error
}

// PeerDeliverClient defines the interface for a peer deliver client
type PeerDeliverClient interface {
	Deliver(ctx context.Context, opts ...grpc.CallOption) (api.Deliver, error)
	DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (api.Deliver, error)
}
