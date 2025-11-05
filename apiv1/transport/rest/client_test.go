package rest

import (
	"reflect"
	"testing"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
)

// TestGetApiRuntime exercises both default and custom ApiResourcePath behavior.
func TestGetApiRuntime(t *testing.T) {
	t.Run("Default BasePath", func(t *testing.T) {
		cfg := &apiv1.SdkConfig{Url: "https://example.com"}
		rt, err := GetApiRuntime(cfg, true)
		if err != nil {
			t.Fatalf("GetApiRuntime returned error: %v", err)
		}
		v := reflect.ValueOf(rt).Elem()

		t.Run("Host is example.com", func(t *testing.T) {
			host := v.FieldByName("Host").String()
			if host != "example.com" {
				t.Errorf("expected host 'example.com', got %q", host)
			}
		})

		t.Run("BasePath is default", func(t *testing.T) {
			basePath := v.FieldByName("BasePath").String()
			if basePath != apiv1.CellsApiResourcePath {
				t.Errorf("expected BasePath %q, got %q", apiv1.CellsApiResourcePath, basePath)
			}
		})

		t.Run("Schemes contains https", func(t *testing.T) {
			schemes := v.FieldByName("schemes")
			if schemes.Len() != 1 || schemes.Index(0).String() != "https" {
				t.Errorf("expected single scheme 'https', got %v", schemes.Interface())
			}
		})
	})

	t.Run("Custom BasePath", func(t *testing.T) {
		custom := "/pub/a"
		cfg := &apiv1.SdkConfig{
			Url:               "http://localhost:1234",
			ApiResourcePrefix: custom,
		}
		rt, err := GetApiRuntime(cfg, true)
		if err != nil {
			t.Fatalf("GetApiRuntime returned error: %v", err)
		}
		basePath := reflect.ValueOf(rt).Elem().FieldByName("BasePath").String()
		if basePath != custom {
			t.Errorf("expected BasePath %q, got %q", custom, basePath)
		}
	})
}
