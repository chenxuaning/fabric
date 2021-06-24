/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package utils

// ErrToString converts and error to a string. If the error is nil, it returns the string "<clean>"
func ErrToString(err error) string {
	if err != nil {
		return err.Error()
	}

	return "<clean>"
}
