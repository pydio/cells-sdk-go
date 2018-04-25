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

As a last step of our process, we also store the resulting swager in a simple string that we store in a go file.

## A few notes about the use of the resulting API

### WebServices

- The blue print of our web services is specified and implemented in cells/common/service.
- Each service is then configured and declared in the corresponding `.../rest/plugins.go` file. The service are then instantiated and managed by the Pydio Cells core framework.
- In go, we use a specific Json marshaller: JSonPB, among others to be able to use TEXT value of the various enums than the corresponding number in the JSON serialised messages.
- Base path in the swagger is defined by (and must fit with) the suffix of rest service's name. Typically service pydio.rest.idm can be found at `<host>/a/idm`.
- We also need to configure the rest service by implementing 2 methods in each rest package, typically:

dans les rest.go , on a ajouter un decodre et un encoder pour serialiser avec JSONPB a la lib qui nous permet de manipuler les service rest --> 
    cf rest-entity-rw.go & service-web.go


```go
// SwaggerTags desclares the name of the method that are implemented by this struct.
func (s *RoleHandler) SwaggerTags() []string {
    return []string{"RoleService"}
}

// Filter can be use to change defined paths.
func (s *RoleHandler) Filter() func(string) string {
    return nil
}
```




en plus: micro api 

on démarre un service micro api -> gateway/micro/rest
il fait le gateway on attaque micro api et lui dispatch

cf github de micro


Pour définir son root path maintenant on édit pydio.json

(on a fait ca a cause des histoire d'enable /disable du SSL)


pour generer les clients php et js, cf front/core/generate-api.sh

Le JWT est entièrement géré dans le PHP
-> pour l'integration gui.ajax/res/js/core/http/restClient.es6


dex -> openid connect



penser à bien toujours utiliser le plugin rest error 500  rest-error-writer


-> generer le client GO
-> faire une premiere PoC de la requete avec récupération du JWT
-> par exemple lister les users

 