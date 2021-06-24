/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ccprovider

import (
	"fmt"
	"sync"
)

// ccInfoCacheImpl implements in-memory cache for ChaincodeData
// needed by endorser to verify if the local instantiation policy
// matches the instantiation policy on a channel before honoring
// an invoke
type ccInfoCacheImpl struct {
	sync.RWMutex

	cache        map[string]*ChaincodeData
	cacheSupport CCCacheSupport
}

// NewCCInfoCache returns a new cache on top of the supplied CCInfoProvider instance
func NewCCInfoCache(cs CCCacheSupport) *ccInfoCacheImpl {
	return &ccInfoCacheImpl{
		cache:        make(map[string]*ChaincodeData),
		cacheSupport: cs,
	}
}

func (c *ccInfoCacheImpl) GetChaincodeData(ccname string, ccversion string) (*ChaincodeData, error) {
	// c.cache is guaranteed to be non-nil

	key := ccname + "/" + ccversion

	c.RLock()
	ccdata, in := c.cache[key]
	c.RUnlock()

	if !in {
		var err error

		// the chaincode data is not in the cache
		// try to look it up from the file system
		ccpack, err := c.cacheSupport.GetChaincode(ccname, ccversion)
		if err != nil || ccpack == nil {
			return nil, fmt.Errorf("cannot retrieve package for chaincode %s/%s, error %s", ccname, ccversion, err)
		}

		// we have a non-nil ChaincodeData, put it in the cache
		c.Lock()
		ccdata = ccpack.GetChaincodeData()
		c.cache[key] = ccdata
		c.Unlock()
	}

	return ccdata, nil
}
