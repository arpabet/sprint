/*
 * Copyright (c) 2025-2026 Karagatan LLC.
 * SPDX-License-Identifier: Apache-2.0
 */

package sprintutils

import (
	"net/http"
	"net/url"

	rt "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/xerrors"
)

func FindGatewayHandler(srv *http.Server, pattern string) (*rt.ServeMux, error) {
	handler := srv.Handler

	switch mux := handler.(type) {
	case *rt.ServeMux:
		return mux, nil
	case *http.ServeMux:
		return findGatewayAPIHandler(mux, pattern)
	default:
		return nil, xerrors.Errorf("unknown server handler '%v'", handler)
	}
}

func findGatewayAPIHandler(mux *http.ServeMux, pattern string) (*rt.ServeMux, error) {

	u, err := url.Parse("http://localhost:/api/")
	if err != nil {
		return nil, xerrors.Errorf("parsing configuration URL error, %v", err)
	}
	req := &http.Request{
		Method:     "GET",
		URL:        u,
		Host:       "localhost",
		RequestURI: pattern,
	}

	handler, foundPattern := mux.Handler(req)
	if foundPattern != pattern {
		return nil, xerrors.Errorf("invalid configuration of http mux, found pattern '%s' whereas expected '%s'", foundPattern, pattern)
	}

	if handler == nil {
		return nil, xerrors.Errorf("handler not found for pattern '%s'", pattern)
	}

	rtMux, ok := handler.(*rt.ServeMux)
	if !ok {
		return nil, xerrors.Errorf("non gateway mux '%v' found on pattern '%s'", handler, pattern)
	}

	return rtMux, nil
}
