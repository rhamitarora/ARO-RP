package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

type appLensRequestOptions interface {
	toHeaders() *map[string]string
}
