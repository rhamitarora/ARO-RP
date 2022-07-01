package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import "fmt"

// GetDetectorOptions includes options for GetDetector operation.
type GetDetectorOptions struct {
	ResourceID string
	DetectorID string
}

func (options *GetDetectorOptions) toHeaders() *map[string]string {
	if options.ResourceID == "" || options.DetectorID == "" {
		return nil
	}

	headers := make(map[string]string)
	headers[headerXmsPathQuery] = fmt.Sprintf("%s/detectors/%s", options.ResourceID, options.DetectorID)
	return &headers
}

// ListDetectorOptions includes options for ListDetector operation.
type ListDetectorsOptions struct {
	ResourceID string
}

func (options *ListDetectorsOptions) toHeaders() *map[string]string {
	if options.ResourceID == "" {
		return nil
	}

	headers := make(map[string]string)
	headers[headerXmsPathQuery] = fmt.Sprintf("%s/detectors", options.ResourceID)
	return &headers
}
