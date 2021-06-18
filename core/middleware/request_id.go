/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package middleware

import (
	"context"
	"net/http"
)

var requestIDKey = requestIDKeyType{}

type requestIDKeyType struct{}

type GenerateIDFunc func() string

type requestID struct {
	generateID GenerateIDFunc
	next       http.Handler
}

func withRequestID(generator GenerateIDFunc) Middleware {
	return func(next http.Handler) http.Handler {
		return &requestID{next: next, generateID: generator}
	}
}

func (r *requestID) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqID := req.Header.Get("X-Request-Id")
	if reqID == "" {
		reqID = r.generateID()
		req.Header.Set("X-Request-Id", reqID)
	}

	ctx := context.WithValue(req.Context(), requestIDKey, reqID)
	req = req.WithContext(ctx)

	w.Header().Add("X-Request-Id", reqID)

	r.next.ServeHTTP(w, req)
}
