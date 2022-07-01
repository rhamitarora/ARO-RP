package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.
// AppLens Client created from CosmosDB Client
// (https://github.com/Azure/azure-sdk-for-go/blob/3f7acd20691214ef2cb1f0132f82115f1df01a8c/sdk/data/azcosmos/cosmos_client.go)

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/gofrs/uuid"
)

// AppLens client is used to interact with the Azure AppLens service.
type Client struct {
	endpoint string
	pipeline runtime.Pipeline
}

// Endpoint used to create the client.
func (c *Client) Endpoint() string {
	return c.endpoint
}

// NewClient creates a new instance of AppLens client with Azure AD access token authentication. It uses the default pipeline configuration.
// endpoint - The applens service endpoint to use.
// cred - The credential used to authenticate with the applens service.
// options - Optional AppLens client options.  Pass nil to accept default values.
func NewClient(endpoint string, cred azcore.TokenCredential, o *ClientOptions) (*Client, error) {
	scope, err := createScopeFromEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{endpoint: endpoint, pipeline: newPipeline([]policy.Policy{runtime.NewBearerTokenPolicy(cred, scope, nil), &appLensBearerTokenPolicy{}}, o)}, nil
}

func newPipeline(authPolicy []policy.Policy, options *ClientOptions) runtime.Pipeline {
	if options == nil {
		options = &ClientOptions{}
	}

	return runtime.NewPipeline("applens", serviceLibVersion,
		runtime.PipelineOptions{
			PerCall:  []policy.Policy{},
			PerRetry: authPolicy,
		},
		&options.ClientOptions)
}

func createScopeFromEndpoint(endpoint string) ([]string, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return []string{fmt.Sprintf("%s://%s/.default", u.Scheme, u.Hostname())}, nil
}

//  ListDetectors obtains the list of detectors for a service from AppLens.
// ctx - The context for the request.
// o - Options for Read operation.
func (c *Client) ListDetectors(
	ctx context.Context,
	o *ListDetectorsOptions) (*http.Response, error) {
	if o == nil {
		o = &ListDetectorsOptions{}
	}

	azResponse, err := c.sendGetRequest(
		ctx,
		o,
		nil)
	if err != nil {
		return nil, err
	}

	return azResponse, nil
}

// GetDetector obtains detector information from AppLens.
// ctx - The context for the request.
// o - Options for Read operation.
func (c *Client) GetDetector(
	ctx context.Context,
	o *GetDetectorOptions) (*http.Response, error) {
	if o == nil {
		o = &GetDetectorOptions{}
	}

	azResponse, err := c.sendGetRequest(
		ctx,
		o,
		nil)
	if err != nil {
		return nil, err
	}

	return azResponse, nil
}

func (c *Client) sendGetRequest(
	ctx context.Context,
	requestOptions appLensRequestOptions,
	requestEnricher func(*policy.Request)) (*http.Response, error) {
	req, err := c.createRequest(ctx, http.MethodGet, requestOptions, requestEnricher)
	if err != nil {
		return nil, err
	}

	return c.executeAndEnsureSuccessResponse(req)
}

func (c *Client) createRequest(
	ctx context.Context,
	method string,
	requestOptions appLensRequestOptions,
	requestEnricher func(*policy.Request)) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, method, c.endpoint)
	if err != nil {
		return nil, err
	}

	if requestOptions != nil {
		headers := requestOptions.toHeaders()
		if headers != nil {
			for k, v := range *headers {
				req.Raw().Header.Set(k, v)
			}
		}
	}

	id := uuid.Must(uuid.NewV4()).String()
	req.Raw().Header.Set(headerXmsClientRequestId, id)
	req.Raw().Header.Set(headerXmsDate, time.Now().UTC().Format(http.TimeFormat))
	req.Raw().Header.Set(headerXmsRequestId, id)

	if requestEnricher != nil {
		requestEnricher(req)
	}

	return req, nil
}

func (c *Client) executeAndEnsureSuccessResponse(request *policy.Request) (*http.Response, error) {
	response, err := c.pipeline.Do(request)
	if err != nil {
		return nil, err
	}

	successResponse := (response.StatusCode >= 200 && response.StatusCode < 300) || response.StatusCode == 304
	if successResponse {
		return response, nil
	}

	return nil, newAppLensError(response)
}
