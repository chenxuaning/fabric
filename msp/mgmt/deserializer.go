/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mgmt

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/msp"
	mspproto "github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

// DeserializersManager is a support interface to
// access the local and channel deserializers
type DeserializersManager interface {

	// Deserialize receives SerializedIdentity bytes and returns the unmarshaled form
	// of the SerializedIdentity, or error on failure
	Deserialize(raw []byte) (*mspproto.SerializedIdentity, error)

	// GetLocalMSPIdentifier returns the local MSP identifier
	GetLocalMSPIdentifier() string

	// GetLocalDeserializer returns the local identity deserializer
	GetLocalDeserializer() msp.IdentityDeserializer

	// GetChannelDeserializers returns a map of the channel deserializers
	GetChannelDeserializers() map[string]msp.IdentityDeserializer
}

// DeserializersManager returns a new instance of DeserializersManager
func NewDeserializersManager() DeserializersManager {
	return &mspDeserializersManager{}
}

type mspDeserializersManager struct{}

func (m *mspDeserializersManager) Deserialize(raw []byte) (*mspproto.SerializedIdentity, error) {
	sId := &mspproto.SerializedIdentity{}
	err := proto.Unmarshal(raw, sId)
	if err != nil {
		return nil, errors.Wrap(err, "could not deserialize a SerializedIdentity")
	}
	return sId, nil
}

func (m *mspDeserializersManager) GetLocalMSPIdentifier() string {
	id, _ := GetLocalMSP().GetIdentifier()
	return id
}

func (m *mspDeserializersManager) GetLocalDeserializer() msp.IdentityDeserializer {
	return GetLocalMSP()
}

func (m *mspDeserializersManager) GetChannelDeserializers() map[string]msp.IdentityDeserializer {
	return GetDeserializers()
}
