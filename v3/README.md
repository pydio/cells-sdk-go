# Cells SDK Go client (v3)

[Go](https://golang.org/) SDK for communicating with a Pydio Cells Server.

This library allows fast implementation of clients using the Go language. 

It provides:

- an automatically generated API client using [go-swagger](https://github.com/go-swagger/go-swagger)
- a transport layer that handles authentication and wraps AWS SDK for file transfers
- a bunch of basic examples that use this SDK to perform simple actions on a running Cells server instance

Although we strive to maintain retrocompatibility in our APIs, this SDK assumes that your server is running Pydio Cells 3.x or later.

For more information, please visit [https://pydio.com/en/docs/developer-guide/cells-sdk-go](https://pydio.com)

## Usage

The library is publicly available, simply:

```sh
go get github.com/pydio/cells-sdk-go/v3 
```

and import necessary sub packages in your code.

The `transport` package provides utilitary methods to ease the set up of a communication with your target Cells instance. You might find the commands that are in the `example` package useful to jump in.

You can also have a look at the [Cells client repository](https://github.com/pydio/cells-client) to see more working examples.

## Migrate from version 2

Between version 2 and 3, we have updated the [go-swagger](https://github.com/go-swagger/go-swagger) version that we use, to [v0.28.0](https://github.com/go-swagger/go-swagger/releases/tag/v0.28.0), this has led to a few breaking changes: you might have to adapt your client code.

We have also reworked our error model in the JSON file: well-known HTTP status are now correctly handled and messages clearer.

Here is a short list of the modification you might have to do in your old code:

- `client.PydioCellsRest` is now `client.PydioCellsRestAPI`
- `enum` objects are managed more cleanly:
  - you have to dereference the pointer to make comparison: `*node.Type == models.TreeNodeTypeCOLLECTION`
  - you cannot use a string to create a parameter. This won't compile: `Type: "COLLECTION"` and you must rather write: `Type: models.NewTreeNodeType(models.TreeNodeTypeCOLLECTION)`
- `Acls` param in struc `models.RestCell` has been renamed `ACLs`

## Versioning policy

As from version 2.2, we release a minor version of the SDK for each minor version of the Pydio Cells Server.

## License

This library is licensed under Apache V2.0 license.