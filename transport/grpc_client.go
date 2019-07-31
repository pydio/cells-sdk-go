package transport

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/go-openapi/strfmt"
	sdk "github.com/pydio/cells-sdk-go"
	"github.com/pydio/cells-sdk-go/client/config_service"
)

var (
	detectedGrpcPorts map[string]string
)

// DetectGrpcPort contacts the discovery service to find the grpc port and loads a valid JWT.
func DetectGrpcPort(config *sdk.SdkConfig, reload bool) (host string, port string, err error) {

	u, e := url.Parse(config.Url)
	if e != nil {
		err = e
		return
	}

	if strings.Contains(u.Host, ":") {
		host, _, err = net.SplitHostPort(u.Host)
		if err != nil {
			return "", "", err
		}
	} else {
		host = u.Host
	}
	var ok bool
	if detectedGrpcPorts == nil {
		detectedGrpcPorts = make(map[string]string)
	}
	if port, ok = detectedGrpcPorts[host]; !ok || reload {
		_, t, e := GetRestClientTransport(config, true)
		if e != nil {
			err = e
			return
		}
		restClient := config_service.New(t, strfmt.Default)
		res, ok := restClient.EndpointsDiscovery(nil)
		if ok != nil {
			err = ok
			return
		}
		if pp, ok := res.Payload.Endpoints["grpc"]; ok && len(pp) > 0 {
			port = strings.Split(pp, ",")[0]
			detectedGrpcPorts[host] = port
		} else {
			err = fmt.Errorf("no port declared for GRPC endpoint")
			return
		}
	}
	return

}
