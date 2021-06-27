/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import (
	"crypto/rand"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"

	"github.com/cetcxinlian/cryptogm/sm2"
	"github.com/hyperledger/fabric/bccsp"
)

type SM2Signature struct {
	R, S *big.Int
}

func MarshalSM2Signature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(SM2Signature{r, s})
}

func UnmarshalSM2Signature(raw []byte) (*big.Int, *big.Int, error) {
	// Unmarshal
	sig := new(SM2Signature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, errors.New("Invalid signature. R must be different from nil.")
	}
	if sig.S == nil {
		return nil, nil, errors.New("Invalid signature. S must be different from nil.")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("Invalid signature. R must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("Invalid signature. S must be larger than zero")
	}

	return sig.R, sig.S, nil
}

func signSM2(k *sm2.PrivateKey, digest []byte, opts bccsp.SignerOpts) (signature []byte, err error) {
	r, s, err := sm2.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}
	return MarshalSM2Signature(r, s)
}

func verifySM2(k *sm2.PublicKey, signature, digest []byte, opts bccsp.SignerOpts) (valid bool, err error) {
	r, s, err := UnmarshalSM2Signature(signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}
	return sm2.Verify(k, digest, r, s), nil
}

type sm2Signer struct{}

func (s *sm2Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return signSM2(k.(*sm2PrivateKey).privKey, digest, opts)
}

type sm2PrivateKeyVerifier struct{}

func (v *sm2PrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySM2(&(k.(*sm2PrivateKey).privKey.PublicKey), signature, digest, opts)
}

type sm2PublicKeyKeyVerifier struct{}

func (v *sm2PublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return verifySM2(k.(*sm2PublicKey).pubKey, signature, digest, opts)
}
