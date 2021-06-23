/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

// SubChannelSelectionCriteria describes a way of selecting peers from a sub-channel
// given their signatures
type SubChannelSelectionCriteria func(signature PeerSignature) bool
