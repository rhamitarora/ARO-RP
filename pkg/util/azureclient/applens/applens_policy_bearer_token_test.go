package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

func TestConvertBearerToken(t *testing.T) {
	srv, close := NewTLSServer()
	defer close()
	srv.SetResponse(WithStatusCode(http.StatusOK))

	verifier := bearerTokenVerify{}
	pl := runtime.NewPipeline("applenstest", "v1.0.0", runtime.PipelineOptions{PerCall: []policy.Policy{&mockAuthPolicy{}, &appLensBearerTokenPolicy{}, &verifier}}, &policy.ClientOptions{Transport: srv})
	req, err := runtime.NewRequest(context.Background(), http.MethodGet, srv.URL())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = pl.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if verifier.authHeaderContent != "type=aad&ver=1.0&sig=this is a test token" {
		t.Fatalf("Expected auth header content to be 'type=aad&ver=1.0&sig=this is a test token', got %s", verifier.authHeaderContent)
	}
}

type bearerTokenVerify struct {
	authHeaderContent string
}

func (p *bearerTokenVerify) Do(req *policy.Request) (*http.Response, error) {
	p.authHeaderContent = req.Raw().Header.Get(headerAuthorization)

	return req.Next()
}

type mockAuthPolicy struct{}

func (p *mockAuthPolicy) Do(req *policy.Request) (*http.Response, error) {
	req.Raw().Header.Set(headerAuthorization, "Bearer this is a test token")

	return req.Next()
}
