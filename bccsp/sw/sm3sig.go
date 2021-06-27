/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sw

import "hash"

type sm3sig struct {
	msg []byte
}

func NewSM3Sig() hash.Hash {
	return &sm3sig{}
}

func (d *sm3sig) Write(p []byte) (n int, err error) {
	d.msg = append(d.msg, p...)
	return len(d.msg), nil
}

func (d *sm3sig) Sum(b []byte) []byte {
	if b != nil {
		panic("sm3sig fail: b must be nil")
	}

	return d.msg
}

func (d *sm3sig) Reset() {
	d.msg = d.msg[:0]
}

func (d *sm3sig) Size() int {
	return 0
}

func (d *sm3sig) BlockSize() int {
	return 0
}
