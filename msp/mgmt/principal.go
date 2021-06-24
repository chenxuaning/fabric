/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package mgmt

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

const (
	// Admins is the label for the local MSP admins
	Admins = "Admins"

	// Members is the label for the local MSP members
	Members = "Members"
)

type MSPPrincipalGetter interface {
	// Get returns an MSP principal for the given role
	Get(role string) (*msp.MSPPrincipal, error)
}

func NewLocalMSPPrincipalGetter() MSPPrincipalGetter {
	return &localMSPPrincipalGetter{}
}

type localMSPPrincipalGetter struct{}

func (m *localMSPPrincipalGetter) Get(role string) (*msp.MSPPrincipal, error) {
	mspid, err := GetLocalMSP().GetIdentifier()
	if err != nil {
		return nil, errors.WithMessage(err, "could not extract local msp identifier")
	}

	switch role {
	case Admins:
		principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &msp.MSPPrincipal{
			PrincipalClassification: msp.MSPPrincipal_ROLE,
			Principal:               principalBytes}, nil
	case Members:
		principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &msp.MSPPrincipal{
			PrincipalClassification: msp.MSPPrincipal_ROLE,
			Principal:               principalBytes}, nil
	default:
		return nil, errors.Errorf("MSP Principal role [%s] not recognized", role)
	}
}
