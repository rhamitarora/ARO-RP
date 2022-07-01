package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

const lenBearerTokenPrefix = len("Bearer ")

type appLensBearerTokenPolicy struct {
}

func (b *appLensBearerTokenPolicy) Do(req *policy.Request) (*http.Response, error) {
	currentAuthorization := req.Raw().Header.Get(headerAuthorization)
	if currentAuthorization == "" {
		return nil, errors.New("authorization header is missing")
	}

	token := currentAuthorization[lenBearerTokenPrefix:]
	req.Raw().Header.Set(headerAuthorization, fmt.Sprintf("type=aad&ver=1.0&sig=%v", token))
	return req.Next()
}
