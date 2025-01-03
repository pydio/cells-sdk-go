package example

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"testing"
	"time"

	cellsSdk "github.com/pydio/cells-sdk-go/v5"
	"github.com/pydio/cells-sdk-go/v5/client/node_service"
	"github.com/pydio/cells-sdk-go/v5/models"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

var DefaultConfig = &cellsSdk.SdkConfig{
	Url:        "https://localhost:8080",
	SkipVerify: true,
	IdToken:    "KAbNSWBGTP3G8DtWlxrT8w-a4Diw3Sl-712qYmZfkyM._OcK1Fc6wOB315SgTeFEoMXSKRZ-3HrLu6nLnnuluCU",
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

	bearerTokenAuth := httptransport.BearerToken(DefaultConfig.IdToken)
	cli, err := transport.GetRestClient(DefaultConfig, false)
	if err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("common-files/test-%s.txt", unique())

	p1 := node_service.NewCreateParams()
	p1.Body = &models.RestCreateRequest{
		Inputs: []*models.RestIncomingNode{{
			Locator: &models.RestNodeLocator{Path: path},
			Type:    models.NewTreeNodeType("LEAF"),
		}},
		Recursive: false,
	}
	r1, err := cli.NodeService.Create(p1, bearerTokenAuth)
	if err != nil {
		t.Fatalf("unable to create empty file: %s", err.Error())
	}
	if r1 == nil {
		t.Fatalf("expected response, got nil")
	}

	p2 := &node_service.LookupParams{
		Body: &models.RestLookupRequest{
			Locators: &models.RestNodeLocators{Many: []*models.RestNodeLocator{{Path: path}}},
		},
	}
	r2, err := cli.NodeService.Lookup(p2, bearerTokenAuth)
	if r2 == nil {
		t.Fatalf("expected response, got nil")
	} else if len(r2.Payload.Nodes) < 1 {
		t.Fatalf("no node found for %s", path)
	}

	p3 := node_service.NewGetByUUIDParams().WithUUID(*r2.Payload.Nodes[0].UUID)
	r3, err := cli.NodeService.GetByUUID(p3, bearerTokenAuth)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if r3 == nil {
		t.Fatalf("expected response, got nil")
	}
}

func TestNodeService_BasicAuth(t *testing.T) {

	basicAuth := httptransport.BasicAuth("bob", "bob")
	cli, err := transport.GetRestClient(DefaultConfig, true)
	if err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("common-files/with-basic-%s.txt", unique())

	p1 := node_service.NewCreateParams()
	p1.Body = &models.RestCreateRequest{
		Inputs: []*models.RestIncomingNode{{
			Locator: &models.RestNodeLocator{Path: path},
			Type:    models.NewTreeNodeType("LEAF"),
		}},
		Recursive: false,
	}
	r1, err := cli.NodeService.Create(p1, basicAuth)
	if err != nil {
		t.Fatalf("unable to create empty file with basic auth: %s", err.Error())
	}
	if r1 == nil {
		t.Fatalf("expected response, got nil")
	}

	p2 := &node_service.LookupParams{
		Body: &models.RestLookupRequest{
			Locators: &models.RestNodeLocators{Many: []*models.RestNodeLocator{{Path: path}}},
		},
	}
	r2, err := cli.NodeService.Lookup(p2, basicAuth)
	if r2 == nil {
		t.Fatalf("expected response, got nil")
	} else if len(r2.Payload.Nodes) < 1 {
		t.Fatalf("no node found for %s", path)
	}
}

func TestNodeService_GetByUUID_Error(t *testing.T) {
	cli, err := transport.GetRestClient(DefaultConfig, false)
	if err != nil {
		t.Fatal(err)
	}

	params := node_service.NewGetByUUIDParams().WithUUID("invalid-uuid")
	_, err = cli.NodeService.GetByUUID(params, httptransport.BearerToken(DefaultConfig.IdToken))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func unique() string {
	rand := fmt.Sprintf("%d", time.Now().Nanosecond())
	hashProvider := md5.New()
	hashProvider.Write([]byte(rand))
	return hex.EncodeToString(hashProvider.Sum(nil))[:6]
}
