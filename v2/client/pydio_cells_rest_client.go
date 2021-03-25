// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v2/client/acl_service"
	"github.com/pydio/cells-sdk-go/v2/client/activity_service"
	"github.com/pydio/cells-sdk-go/v2/client/admin_tree_service"
	"github.com/pydio/cells-sdk-go/v2/client/config_service"
	"github.com/pydio/cells-sdk-go/v2/client/frontend_service"
	"github.com/pydio/cells-sdk-go/v2/client/graph_service"
	"github.com/pydio/cells-sdk-go/v2/client/install_service"
	"github.com/pydio/cells-sdk-go/v2/client/jobs_service"
	"github.com/pydio/cells-sdk-go/v2/client/log_service"
	"github.com/pydio/cells-sdk-go/v2/client/mailer_service"
	"github.com/pydio/cells-sdk-go/v2/client/meta_service"
	"github.com/pydio/cells-sdk-go/v2/client/policy_service"
	"github.com/pydio/cells-sdk-go/v2/client/role_service"
	"github.com/pydio/cells-sdk-go/v2/client/search_service"
	"github.com/pydio/cells-sdk-go/v2/client/share_service"
	"github.com/pydio/cells-sdk-go/v2/client/templates_service"
	"github.com/pydio/cells-sdk-go/v2/client/token_service"
	"github.com/pydio/cells-sdk-go/v2/client/tree_service"
	"github.com/pydio/cells-sdk-go/v2/client/update_service"
	"github.com/pydio/cells-sdk-go/v2/client/user_meta_service"
	"github.com/pydio/cells-sdk-go/v2/client/user_service"
	"github.com/pydio/cells-sdk-go/v2/client/workspace_service"
)

// Default pydio cells rest HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https", "wss"}

// NewHTTPClient creates a new pydio cells rest HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PydioCellsRest {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new pydio cells rest HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PydioCellsRest {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new pydio cells rest client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PydioCellsRest {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PydioCellsRest)
	cli.Transport = transport

	cli.ACLService = acl_service.New(transport, formats)

	cli.ActivityService = activity_service.New(transport, formats)

	cli.AdminTreeService = admin_tree_service.New(transport, formats)

	cli.ConfigService = config_service.New(transport, formats)

	cli.FrontendService = frontend_service.New(transport, formats)

	cli.GraphService = graph_service.New(transport, formats)

	cli.InstallService = install_service.New(transport, formats)

	cli.JobsService = jobs_service.New(transport, formats)

	cli.LogService = log_service.New(transport, formats)

	cli.MailerService = mailer_service.New(transport, formats)

	cli.MetaService = meta_service.New(transport, formats)

	cli.PolicyService = policy_service.New(transport, formats)

	cli.RoleService = role_service.New(transport, formats)

	cli.SearchService = search_service.New(transport, formats)

	cli.ShareService = share_service.New(transport, formats)

	cli.TemplatesService = templates_service.New(transport, formats)

	cli.TokenService = token_service.New(transport, formats)

	cli.TreeService = tree_service.New(transport, formats)

	cli.UpdateService = update_service.New(transport, formats)

	cli.UserMetaService = user_meta_service.New(transport, formats)

	cli.UserService = user_service.New(transport, formats)

	cli.WorkspaceService = workspace_service.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PydioCellsRest is a client for pydio cells rest
type PydioCellsRest struct {
	ACLService *acl_service.Client

	ActivityService *activity_service.Client

	AdminTreeService *admin_tree_service.Client

	ConfigService *config_service.Client

	FrontendService *frontend_service.Client

	GraphService *graph_service.Client

	InstallService *install_service.Client

	JobsService *jobs_service.Client

	LogService *log_service.Client

	MailerService *mailer_service.Client

	MetaService *meta_service.Client

	PolicyService *policy_service.Client

	RoleService *role_service.Client

	SearchService *search_service.Client

	ShareService *share_service.Client

	TemplatesService *templates_service.Client

	TokenService *token_service.Client

	TreeService *tree_service.Client

	UpdateService *update_service.Client

	UserMetaService *user_meta_service.Client

	UserService *user_service.Client

	WorkspaceService *workspace_service.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PydioCellsRest) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.ACLService.SetTransport(transport)

	c.ActivityService.SetTransport(transport)

	c.AdminTreeService.SetTransport(transport)

	c.ConfigService.SetTransport(transport)

	c.FrontendService.SetTransport(transport)

	c.GraphService.SetTransport(transport)

	c.InstallService.SetTransport(transport)

	c.JobsService.SetTransport(transport)

	c.LogService.SetTransport(transport)

	c.MailerService.SetTransport(transport)

	c.MetaService.SetTransport(transport)

	c.PolicyService.SetTransport(transport)

	c.RoleService.SetTransport(transport)

	c.SearchService.SetTransport(transport)

	c.ShareService.SetTransport(transport)

	c.TemplatesService.SetTransport(transport)

	c.TokenService.SetTransport(transport)

	c.TreeService.SetTransport(transport)

	c.UpdateService.SetTransport(transport)

	c.UserMetaService.SetTransport(transport)

	c.UserService.SetTransport(transport)

	c.WorkspaceService.SetTransport(transport)

}
