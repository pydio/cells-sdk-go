package transport

import (
	"context"
	"reflect"
	"testing"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
)

// TestGetRuntimeTransport verifies the transport returned by GetRuntimeTransport has the correct settings.
func TestGetRuntimeTransport(t *testing.T) {
	ctx := context.Background()
	cfg := &apiv2.SdkConfig{Url: "https://example.com"}
	rt, err := GetRuntimeTransport(ctx, cfg)
	if err != nil {
		t.Fatalf("GetRuntimeTransport returned error: %v", err)
	}
	v := reflect.ValueOf(rt).Elem()
	// Host should match the URL host
	if host := v.FieldByName("Host").String(); host != "example.com" {
		t.Errorf("expected host 'example.com', got %q", host)
	}
	// BasePath should use the default CellsApiPrefix
	if base := v.FieldByName("BasePath").String(); base != CellsApiPrefix {
		t.Errorf("expected BasePath %q, got %q", CellsApiPrefix, base)
	}
	// Schemes should contain the URL scheme
	schemes := v.FieldByName("schemes")
	if schemes.Len() != 1 || schemes.Index(0).String() != "https" {
		t.Errorf("expected schemes [\"https\"], got %v", schemes.Interface())
	}
	// Context should be preserved
	if gotCtx, ok := v.FieldByName("Context").Interface().(context.Context); !ok || gotCtx != ctx {
		t.Errorf("expected Context to be %v, got %v", ctx, gotCtx)
	}
}

// TestGetClientTransport checks both default and custom ApiResourcePrefix behavior.
func TestGetClientTransport(t *testing.T) {
	t.Run("Default BasePath", func(t *testing.T) {
		cfg := &apiv2.SdkConfig{Url: "http://example.com"}
		rt, err := GetClientTransport(cfg, true)
		if err != nil {
			t.Fatalf("GetClientTransport returned error: %v", err)
		}
		v := reflect.ValueOf(rt).Elem()
		// Host should match the URL host
		if host := v.FieldByName("Host").String(); host != "example.com" {
			t.Errorf("expected host 'example.com', got %q", host)
		}
		// BasePath should use the default CellsApiPrefix
		if base := v.FieldByName("BasePath").String(); base != CellsApiPrefix {
			t.Errorf("expected BasePath %q, got %q", CellsApiPrefix, base)
		}
		// Schemes should contain the URL scheme
		schemes := v.FieldByName("schemes")
		if schemes.Len() != 1 || schemes.Index(0).String() != "http" {
			t.Errorf("expected schemes [\"http\"], got %v", schemes.Interface())
		}
		// Context should be background
		if gotCtx, ok := v.FieldByName("Context").Interface().(context.Context); !ok {
			t.Errorf("expected Context to be context.Context, got %v", v.FieldByName("Context").Kind())
		} else if gotCtx != context.Background() {
			t.Errorf("expected Context to be Background, got %v", gotCtx)
		}
	})

	t.Run("Custom BasePath", func(t *testing.T) {
		custom := "/myapi/a"
		cfg := &apiv2.SdkConfig{Url: "http://localhost:1234", ApiResourcePrefix: custom}
		rt, err := GetClientTransport(cfg, true)
		if err != nil {
			t.Fatalf("GetClientTransport returned error: %v", err)
		}
		v := reflect.ValueOf(rt).Elem()
		// BasePath should match the custom prefix
		if base := v.FieldByName("BasePath").String(); base != custom {
			t.Errorf("expected BasePath %q, got %q", custom, base)
		}
	})
}
