/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"crypto/rand"
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/common/metadata"
)

// ComputeSHA256 returns SHA2-256 on data
func ComputeSHA256(data []byte) (hash []byte) {
	hash, err := factory.GetDefault().Hash(data, &bccsp.SHA256Opts{})
	if err != nil {
		panic(fmt.Errorf("Failed computing SHA256 on [% x]", data))
	}
	return
}

// ComputeSHA3256 returns SHA3-256 on data
func ComputeSHA3256(data []byte) (hash []byte) {
	hash, err := factory.GetDefault().Hash(data, &bccsp.SHA3_256Opts{})
	if err != nil {
		panic(fmt.Errorf("Failed computing SHA3_256 on [% x]", data))
	}
	return
}

// GenerateBytesUUID returns a UUID based on RFC 4122 returning the generated bytes
func GenerateBytesUUID() []byte {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		panic(fmt.Sprintf("Error generating UUID: %s", err))
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return uuid
}

// GenerateUUID returns a UUID based on RFC 4122
func GenerateUUID() string {
	uuid := GenerateBytesUUID()
	return idBytesToStr(uuid)
}

// CreateUtcTimestamp returns a google/protobuf/Timestamp in UTC
func CreateUtcTimestamp() *timestamp.Timestamp {
	now := time.Now().UTC()
	secs := now.Unix()
	nanos := int32(now.UnixNano() - (secs * 1000000000))
	return &(timestamp.Timestamp{Seconds: secs, Nanos: nanos})
}

func idBytesToStr(id []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
}

// ToChaincodeArgs converts string args to []byte args
func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// GetSysCCVersion returns the version of all system chaincodes
// This needs to be revisited on policies around system chaincode
// "upgrades" from user and relationship with "fabric" upgrade. For
// now keep it simple and use the fabric's version stamp
func GetSysCCVersion() string {
	return metadata.Version
}

// ConcatenateBytes is useful for combining multiple arrays of bytes, especially for
// signatures or digests over multiple fields
func ConcatenateBytes(data ...[]byte) []byte {
	finalLength := 0
	for _, slice := range data {
		finalLength += len(slice)
	}
	result := make([]byte, finalLength)
	last := 0
	for _, slice := range data {
		for i := range slice {
			result[i+last] = slice[i]
		}
		last += len(slice)
	}
	return result
}

// `flatten` recursively retrieves every leaf node in a struct in depth-first fashion
// and aggregate the results into given string slice with format: "path.to.leaf = value"
// in the order of definition. Root name is ignored in the path. This helper function is
// useful to pretty-print a struct, such as configs.
// for example, given data structure:
// A{
//   B{
//     C: "foo",
//     D: 42,
//   },
//   E: nil,
// }
// it should yield a slice of string containing following items:
// [
//   "B.C = \"foo\"",
//   "B.D = 42",
//   "E =",
// ]
func Flatten(i interface{}) []string {
	var res []string
	flatten("", &res, reflect.ValueOf(i))
	return res
}

const DELIMITER = "."

func flatten(k string, m *[]string, v reflect.Value) {
	delimiter := DELIMITER
	if k == "" {
		delimiter = ""
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			*m = append(*m, fmt.Sprintf("%s =", k))
			return
		}
		flatten(k, m, v.Elem())
	case reflect.Struct:
		if x, ok := v.Interface().(fmt.Stringer); ok {
			*m = append(*m, fmt.Sprintf("%s = %v", k, x))
			return
		}

		for i := 0; i < v.NumField(); i++ {
			flatten(k+delimiter+v.Type().Field(i).Name, m, v.Field(i))
		}
	case reflect.String:
		// It is useful to quote string values
		*m = append(*m, fmt.Sprintf("%s = \"%s\"", k, v))
	default:
		*m = append(*m, fmt.Sprintf("%s = %v", k, v))
	}
}