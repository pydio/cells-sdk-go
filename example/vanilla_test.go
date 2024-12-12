package example

import (
	"testing"

	cellsSdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/client/node_service"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

var DefaultConfig = &cellsSdk.SdkConfig{
	Url:        "https://localhost:8080",
	SkipVerify: true,
	IdToken:    "b06mpsmXzz79Z_HucfFtTeK3VsFQX0_fkPKp-elJLZQ.kM0jaI98v2zDHOa4rh7JJXjj3TjcIT04aCgfohwE2eE",
}

func TestNewHTTPClient(t *testing.T) {
	cli, err := transport.GetRestClient(DefaultConfig, false)
	if err != nil {
		t.Fatal(err)
	}

	if cli.NodeService == nil {
		t.Fatalf("expected NodeService to be initialized, got nil")
	}
}

func TestNodeService_GetByUUID(t *testing.T) {
	cli, err := transport.GetRestClient(DefaultConfig, false)
	if err != nil {
		t.Fatal(err)
	}

	params := node_service.NewGetByPathParams().WithPath("common-files/test.txt")
	resp, err := cli.NodeService.GetByPath(params)
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}

	params2 := node_service.NewGetByUUIDParams().WithUUID(resp.Payload.UUID)
	resp2, err := cli.NodeService.GetByUUID(params2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp2 == nil {
		t.Fatalf("expected response, got nil")
	}
}

func TestNodeService_GetByUUID_Error(t *testing.T) {
	cli, err := transport.GetRestClient(DefaultConfig, false)
	if err != nil {
		t.Fatal(err)
	}

	params := node_service.NewGetByUUIDParams().WithUUID("invalid-uuid")
	_, err = cli.NodeService.GetByUUID(params)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

//func TestNodeService_Create(t *testing.T) {
//	cli, err := transport.GetRestClient(DefaultConfig, false)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	params := node_service.NewCreateParams()
//	params.Body.Inputs
//	resp, err := cli.NodeService.Create(params)
//
//	if err != nil {
//		t.Fatalf("expected no error, got %v", err)
//	}
//	if resp == nil {
//		t.Fatalf("expected response, got nil")
//	}
//}
