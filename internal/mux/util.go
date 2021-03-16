package mux

import (
	"context"
	"net/http"
)

// ----------------------------------------------------------------------------
// Request context
// ----------------------------------------------------------------------------

type paramsKey struct{}

func PathParams(req *http.Request) Params {
	return req.Context().Value(paramsKey{}).(Params)
}

// Adds variables to the request
func requestWithPathParams(req *http.Request, params Params) *http.Request {
	if len(params) == 0 {
		params = Params{}
	}

	return req.Clone(context.WithValue(req.Context(), paramsKey{}, params))
}