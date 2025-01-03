package example

import (
	"testing"

	cellsSdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/client/node_service"
	"github.com/pydio/cells-sdk-go/v5/models"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

var DefaultConfig = &cellsSdk.SdkConfig{
	Url:        "https://localhost:8080",
	SkipVerify: true,
	IdToken:    "0XmNa4mhWfLhURwFP6_XJVuEDgPDnoa8EKvJB55fY8g.yddonfgcR4oogCzRqIfimzN7dRNk7si3X3TFRL7hYjo",
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

	path := "common-files/test.txt"

	p1 := node_service.NewCreateParams()
	p1.Body = &models.RestCreateRequest{
		Inputs: []*models.RestIncomingNode{&models.RestIncomingNode{
			//			ContentType: "text/plain",
			Locator: &models.RestNodeLocator{Path: path},
			Type:    models.NewTreeNodeType("LEAF"),
		}},
		Recursive: false,
	}
	resp, err := cli.NodeService.Create(p1)
	if err != nil {
		t.Fatalf("unable to create empty file: %s", err.Error())
	}
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}

	params2 := node_service.NewGetByPathParams().WithPath("common-files/test.txt")
	resp2, err := cli.NodeService.GetByPath(params2)
	if resp2 == nil {
		t.Fatalf("expected response, got nil")
	}

	params3 := node_service.NewGetByUUIDParams().WithUUID(resp2.Payload.UUID)
	resp3, err := cli.NodeService.GetByUUID(params3)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if resp3 == nil {
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
