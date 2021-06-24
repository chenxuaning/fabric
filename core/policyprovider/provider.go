/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package policyprovider

import (
	"github.com/hyperledger/fabric/core/peer"
	"github.com/hyperledger/fabric/core/policy"
	"github.com/hyperledger/fabric/msp/mgmt"
)

// init is called when this package is loaded. This implementation registers the factory
func init() {
	policy.RegisterPolicyCheckerFactory(&defaultFactory{})
}

type defaultFactory struct{}

func (f *defaultFactory) NewPolicyChecker() policy.PolicyChecker {
	return policy.NewPolicyChecker(
		peer.NewChannelPolicyManagerGetter(),
		mgmt.GetLocalMSP(),
		mgmt.NewLocalMSPPrincipalGetter(),
	)
}

// GetPolicyChecker returns instances of PolicyChecker;
func GetPolicyChecker() policy.PolicyChecker {
	return policy.GetPolicyChecker()
}
