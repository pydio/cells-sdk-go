# Use of OpenAPI within Pydio Cells

_DISCLAIMER: this documentation's aim is to enable the developper team to share information and notes while working on the current project. It is neither a reference documentation nor something you should rely on. It is rather a source of interesting resources and articles that we gather here to be able to find them again quickly and share them._

## Tips and tricks

A list of web resources of interest while working on the go SDK.

- a short [overview and example](https://medium.com/@marcus.olsson/writing-a-go-client-for-your-restful-api-c193a2f4998c) with go as API client.

## About Pydio Cells, Protobuff and OpenAPI

Pydio Cells use proto3 to describe model and relevant services.

All proto3 sources files can be found under `cells/common/proto/...`

We use a custom micro version of protoc (with a specific micro plugin) to generate the go files.

We then apply a specific version of the protoc compiler with a swagger plugin and some google specific options to generate the corresponding swagger file in JSON format from the protobuf file.
This enables us to generate a file that we use to describe our REST api.

The swagger file retrieves all comments from the proto3 file. It also describes the various models.

As a last step of our process, we also store the resulting swagger in a simple string that we store in a go file.

## A few notes about the use of the resulting API

### WebServices

- The blue print of our web services is specified and implemented in cells/common/service.
- Each service is then configured and declared in the corresponding `.../rest/plugins.go` file. The service are then instantiated and managed by the Pydio Cells core framework.
- In go, we use a specific Json marshaller (JSonPB) to encode and decode messages and provide communication with the REST service.  
    This, among others, to be able to use the string representation of the values rather than the corresponding number for the various enums that are then serialized in the JSON messages.  
    See [rest-entity-rw.go](https://github.com/pydio/cells/blob/master/common/service/rest-entity-rw.go) and [service-web.go](https://github.com/pydio/cells/blob/master/common/service/service-web.go)

- Base path in the swagger is defined by (and must fit with) the suffix of rest service's name. Typically service pydio.rest.idm can be found at `<host>/a/idm`.
- We also need to configure the rest service by implementing 2 methods in each rest package, typically:

```go
// SwaggerTags desclares the name of the services that are implemented by this struct.
func (s *RoleHandler) SwaggerTags() []string {
    return []string{"RoleService"}
}

// Filter can be used to change pre-defined paths.
func (s *RoleHandler) Filter() func(string) string {
    return nil
}
```

#### JWT

JWT mechanisms are provided by Dex (based on OpenID connect)

#### Micro API

We furthermore use micro API to provide the gateway.

This service is declared in gateway/micro/plugins.go

It receives the REST request and dispach them.

For more info, see [micro repository on Github](https://github.com/micro/micro)

- To define the root path one should edit the `pydio.json` file that is in the root folder of Pydio Cells data folder (which is typically located at `~/.config/pydio/cells`). We do this to ease enablement and disablement of TLS communication.
