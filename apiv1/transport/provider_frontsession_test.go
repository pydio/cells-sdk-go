package transport

import (
    "reflect"
    "testing"

    "github.com/pydio/cells-sdk-go/v5/apiv1"
)

// Test that the FrontSessionTokenProvider defaults to CellsApiResourcePath when ApiResourcePath is empty.
func TestNewFrontSessionTokenProvider_DefaultApiPath(t *testing.T) {
    cfg := &apiv1.SdkConfig{
        Url: "http://localhost",
    }
    prov, err := NewFrontSessionTokenProvider(cfg)
    if err != nil {
        t.Fatalf("NewFrontSessionTokenProvider returned error: %v", err)
    }
    v := reflect.ValueOf(prov).Elem()
    apiPath := v.FieldByName("apiPath").String()
    if apiPath != apiv1.CellsApiResourcePath {
        t.Errorf("expected apiPath %q, got %q", apiv1.CellsApiResourcePath, apiPath)
    }
}

// Test that the FrontSessionTokenProvider uses a custom ApiResourcePath when set.
func TestNewFrontSessionTokenProvider_CustomApiPath(t *testing.T) {
    custom := "/pub/a"
    cfg := &apiv1.SdkConfig{
        Url:             "http://localhost",
        ApiResourcePath: custom,
    }
    prov, err := NewFrontSessionTokenProvider(cfg)
    if err != nil {
        t.Fatalf("NewFrontSessionTokenProvider returned error: %v", err)
    }
    v := reflect.ValueOf(prov).Elem()
    apiPath := v.FieldByName("apiPath").String()
    if apiPath != custom {
        t.Errorf("expected apiPath %q, got %q", custom, apiPath)
    }
}