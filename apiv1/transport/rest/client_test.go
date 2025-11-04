package rest

import (
	"reflect"
	"testing"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
)

// Test that GetApiRuntime uses the default CellsApiResourcePath when ApiResourcePath is empty.
func TestGetApiRuntime_DefaultBasePath(t *testing.T) {
	cfg := &apiv1.SdkConfig{
		Url: "https://example.com",
	}
	rt, err := GetApiRuntime(cfg, true)
	if err != nil {
		t.Fatalf("GetApiRuntime returned error: %v", err)
	}
	v := reflect.ValueOf(rt).Elem()
	// Verify the Host field
	hostVal := v.FieldByName("Host")
	if !hostVal.IsValid() {
		t.Fatalf("Runtime struct missing Host field")
	}
	host := hostVal.String()
	if host != "example.com" {
		t.Errorf("expected host 'example.com', got %q", host)
	}
	// Verify the BasePath field
	bpVal := v.FieldByName("BasePath")
	if !bpVal.IsValid() {
		t.Fatalf("Runtime struct missing BasePath field")
	}
	basePath := bpVal.String()
	if basePath != apiv1.CellsApiResourcePath {
		t.Errorf("expected basePath %q, got %q", apiv1.CellsApiResourcePath, basePath)
	}
	// Verify the schemes slice (unexported field)
	schemesVal := v.FieldByName("schemes")
	if !schemesVal.IsValid() {
		t.Fatalf("Runtime struct missing unexported schemes field")
	}
	if schemesVal.Len() != 1 {
		t.Errorf("expected 1 scheme, got %d", schemesVal.Len())
	} else {
		scheme := schemesVal.Index(0).String()
		if scheme != "https" {
			t.Errorf("expected scheme 'https', got %q", scheme)
		}
	}
}

// Test that GetApiRuntime respects a custom ApiResourcePath.
func TestGetApiRuntime_CustomBasePath(t *testing.T) {
	custom := "/pub/a"
	cfg := &apiv1.SdkConfig{
		Url:               "http://localhost:1234",
		ApiResourcePrefix: custom,
	}
	rt, err := GetApiRuntime(cfg, true)
	if err != nil {
		t.Fatalf("GetApiRuntime returned error: %v", err)
	}
	v := reflect.ValueOf(rt).Elem()
	// Verify custom BasePath
	bpVal := v.FieldByName("BasePath")
	if !bpVal.IsValid() {
		t.Fatalf("Runtime struct missing BasePath field")
	}
	basePath := bpVal.String()
	if basePath != custom {
		t.Errorf("expected basePath %q, got %q", custom, basePath)
	}
}
