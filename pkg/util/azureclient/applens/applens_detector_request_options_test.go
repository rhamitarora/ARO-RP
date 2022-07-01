package applens

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.s

import (
	"testing"
)

func TestGetDetectorOptionsToHeaders(t *testing.T) {
	options := &GetDetectorOptions{}
	if options.toHeaders() != nil {
		t.Error("toHeaders should return nil")
	}

	options.ResourceID = ""
	options.DetectorID = ""
	header := options.toHeaders()
	if header != nil {
		t.Fatal("toHeaders should return nil")
	}

	options.DetectorID = "testdetector"
	header = options.toHeaders()
	if header != nil {
		t.Fatal("toHeaders should return nil")
	}

	options.ResourceID = "testresourceid"
	header = options.toHeaders()
	if header == nil {
		t.Fatal("toHeaders should return non-nil")
	}

	options.DetectorID = ""
	header = options.toHeaders()
	if header != nil {
		t.Fatal("toHeaders should return nil")
	}
}

func TestListDetectorsOptionsToHeaders(t *testing.T) {
	options := &ListDetectorsOptions{}
	if options.toHeaders() != nil {
		t.Error("toHeaders should return nil")
	}

	options.ResourceID = ""
	header := options.toHeaders()
	if header != nil {
		t.Fatal("toHeaders should return nil")
	}

	options.ResourceID = "testresourceid"
	header = options.toHeaders()
	if header == nil {
		t.Fatal("toHeaders should return non-nil")
	}
}
