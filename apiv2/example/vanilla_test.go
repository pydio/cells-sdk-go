package example

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	openapiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
	apiv2Client "github.com/pydio/cells-sdk-go/v5/apiv2/client"
	"github.com/pydio/cells-sdk-go/v5/apiv2/client/node_service"
	"github.com/pydio/cells-sdk-go/v5/apiv2/models"
	"github.com/pydio/cells-sdk-go/v5/apiv2/transport"
)

var DefaultConfig = &apiv2.SdkConfig{
	Url:        "https://localhost:8080",
	SkipVerify: true,
	IdToken:    "KAbNSWBGTP3G8DtWlxrT8w-a4Diw3Sl-712qYmZfkyM._OcK1Fc6wOB315SgTeFEoMXSKRZ-3HrLu6nLnnuluCU",
	CustomHeaders: map[string]string{
		transport.KeyUserAgent: transport.UserAgent(),
	},
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

func TestNodeService_SimpleCrudWithPat(t *testing.T) {
	// FIXME
	t.Skip("Skipping failing test for now")

	bearerTokenAuth := openapiClient.BearerToken(DefaultConfig.IdToken)
	rTransport, err := transport.GetRuntimeTransport(context.Background(), DefaultConfig)
	if err != nil {
		t.Fatal(err)
	}
	rClient := apiv2Client.New(rTransport, strfmt.Default)

	path := fmt.Sprintf("common-files/test-%s.txt", unique())

	// Create an empty file
	p1 := node_service.NewCreateParams()
	p1.Body = &models.RestCreateRequest{
		Inputs: []*models.RestIncomingNode{{
			Locator: &models.RestNodeLocator{Path: path},
			Type:    models.NewTreeNodeType("LEAF"),
		}},
		Recursive: false,
	}
	r1, err := rClient.NodeService.Create(p1, bearerTokenAuth)
	if err != nil {
		t.Fatalf("unable to create empty file: %s", err.Error())
	}
	if r1 == nil {
		t.Fatalf("expected response, got nil")
	}

	// Retrieve it by path
	p2 := &node_service.LookupParams{
		Body: &models.RestLookupRequest{
			Locators: &models.RestNodeLocators{Many: []*models.RestNodeLocator{{Path: path}}},
		},
	}
	r2, err := rClient.NodeService.Lookup(p2, bearerTokenAuth)
	if r2 == nil {
		t.Fatalf("expected response, got nil")
	} else if len(r2.Payload.Nodes) < 1 {
		t.Fatalf("no node found for %s", path)
	}

	// Retrieve it by UUID
	p3 := node_service.NewGetByUUIDParams().WithUUID(*r2.Payload.Nodes[0].UUID)
	r3, err := rClient.NodeService.GetByUUID(p3, bearerTokenAuth)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if r3 == nil {
		t.Fatalf("expected response, got nil")
	}
}

func TestNodeService_BasicAuth(t *testing.T) {

	// TODO basic Auth is not yet supported by API v2
	//   (we miss the frontend service to get a valid token from legacy credentials)
	t.SkipNow()

	ctx := context.Background()
	basicAuth := transport.BasicAuthWriter(ctx, DefaultConfig)
	rTransport, err := transport.GetRuntimeTransport(context.Background(), DefaultConfig)
	if err != nil {
		t.Fatal(err)
	}
	cli := apiv2Client.New(rTransport, strfmt.Default)

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
	_, err = cli.NodeService.GetByUUID(params, openapiClient.BearerToken(DefaultConfig.IdToken))
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
