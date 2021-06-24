/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package couchdb

import (
	"time"

	"github.com/spf13/viper"
)

// CouchDBDef contains parameters
type CouchDBDef struct {
	URL                   string
	Username              string
	Password              string
	MaxRetries            int
	MaxRetriesOnStartup   int
	RequestTimeout        time.Duration
	CreateGlobalChangesDB bool
}

// GetCouchDBDefinition exposes the useCouchDB variable
func GetCouchDBDefinition() *CouchDBDef {

	couchDBAddress := viper.GetString("ledger.state.couchDBConfig.couchDBAddress")
	username := viper.GetString("ledger.state.couchDBConfig.username")
	password := viper.GetString("ledger.state.couchDBConfig.password")
	maxRetries := viper.GetInt("ledger.state.couchDBConfig.maxRetries")
	maxRetriesOnStartup := viper.GetInt("ledger.state.couchDBConfig.maxRetriesOnStartup")
	requestTimeout := viper.GetDuration("ledger.state.couchDBConfig.requestTimeout")
	createGlobalChangesDB := viper.GetBool("ledger.state.couchDBConfig.createGlobalChangesDB")

	return &CouchDBDef{couchDBAddress, username, password, maxRetries, maxRetriesOnStartup, requestTimeout, createGlobalChangesDB}
}
