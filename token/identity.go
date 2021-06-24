/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package token

// Identity refers to the creator of a tx;
type Identity interface {
	Serialize() ([]byte, error)
}

// SigningIdentity defines the functions necessary to sign an
// array of bytes; it is needed to sign the commands transmitted to
// the prover peer service.
type SigningIdentity interface {
	Identity // extends Identity

	Sign(msg []byte) ([]byte, error)

	GetPublicVersion() Identity
}
