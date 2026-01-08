# Cells SDK Go (v5)

> **WARNING**: the v5 branch of the SDK is **not yet stable**: the API _**will**_ change before the GA release. We yet start to publish some snapshot versions to prepare the adaptation of the various clients.
> **WARNING**: due to a breaking change in the AWS SDK we use under the hood to manage transfers, we had to adapt the server. As a consequence, the latest version of the SDKs (4.4+) requires at least a server that has version 4.4.**12**, or large uploads will fail.  
> If you are stuck with a target server that has an older version, please rather use the latest SDK of the 4.1 branch.  

[Go](https://golang.org/) SDK for communicating with a Pydio Cells Server.

This library allows fast implementation of clients using the Go language.

It provides:

- an automatically generated API client using [go-swagger](https://github.com/go-swagger/go-swagger)
- a transport layer that handles authentication and wraps AWS SDK for file transfers
- a bunch of basic examples that use this SDK to perform simple actions on a running Cells server instance

To use this SDK, your server should be running Pydio Cells 4.4.12 or later.

For further details about migrating from older version of the SDK please refer to the detailed instructions below.

For more information, please visit our [developer guide](https://pydio.com/en/docs/developer-guide).

## Usage

The library is publicly available, simply:

```sh
go get github.com/pydio/cells-sdk-go/v5
```

and import necessary sub-packages in your code.

The `transport` package provides utility methods to ease the set-up of a communication with your target Cells instance. You might find the commands that are in the `example` package useful to jump in.

You can also have a look at the [Cells client repository](https://github.com/pydio/cells-client) to see more working examples.

## Migrate from older versions

### From SDK v4.4

For Cells v5, we started implementing a v2 version of the Cells API. We decided to yet include both version API v1 and API v2 in the SDK.

Thus, we introduced a new level in the packages, you will then find at:

- `pydio/cells-sdk-go/v5/apiv1`: The v1 of the API, mostly stable since the latest releases of the v4, you should only replace your imports from `pydio/cells-sdk-go/v4` in your client code
- `pydio/cells-sdk-go/v5/apiv2`: A first limited version of the new REST API that is exposed by the server. **WARNING** this API is still work in progress and both the API and the SDK **will** change in the near future. Please use with care.

### From SDK v4.1

Nothing changes for the clients, only the dependencies (mainly AWS SDK v2) have been updated: you should only ensure that you are targeting a Cells Server that runs at least v4.4.12 or large file upload will fail.

### From SDK v4.0

For the v4.1, we switched to the v2 of the AWS SDK for Go that we use under the hood to transfer files.

If you are relying upon some of the class defined in the `pydio/cells-sdk-go/v4/transport` package, you have to update your code.  
You can have a look at the `wget` and `list-bucket` commands of the `pydio/cells-sdk-go/v4/example` package to see sample code and refer to the official [AWS documentation](https://aws.github.io/aws-sdk-go-v2/docs/migrating/) to get more help.

### From SDK v3

Between version 3 and 4, we have updated the [go-swagger](https://github.com/go-swagger/go-swagger) version that we use, to [v0.30.3](https://github.com/go-swagger/go-swagger/releases/tag/v0.30.3) and then to [v0.30.5](https://github.com/go-swagger/go-swagger/releases/tag/v0.30.5), this should have no impact on your client code.

We have also changed the name of the swagger spec file in Cells that is now: `cellsapi-rest.swagger.json`.

### From SDK v2

Between version 2 and 3, we have updated the [go-swagger](https://github.com/go-swagger/go-swagger) version that we use, to [v0.28.0](https://github.com/go-swagger/go-swagger/releases/tag/v0.28.0), this has led to a few breaking changes: you might have to adapt your client code.

We have also reworked our error model in the JSON file: well-known HTTP status are now correctly handled and messages clearer.

Here is a short list of the modification you might have to do in your old code:

- `client.PydioCellsRest` is now `client.PydioCellsRestAPI`
- `enum` objects are managed more cleanly:
  - you have to dereference the pointer to make comparison: `*node.Type == models.TreeNodeTypeCOLLECTION`
  - you cannot use a string to create a parameter. This won't compile: `Type: "COLLECTION"` and you must rather write: `Type: models.NewTreeNodeType(models.TreeNodeTypeCOLLECTION)`
- `Acls` param in struct `models.RestCell` has been renamed `ACLs`

## Versioning policy

Between version 2.2 and 3.99 of the Cells server, as the API was still moving and even if we tried to stay backward compatible, we released a minor version of the SDK for each minor version of the Pydio Cells Server: if you are stuck with some servers in v2 or v3, you might still want to use a SDK that has same major.minor version to communicate with your server.

## License

This library is licensed under Apache v2.0 license.
