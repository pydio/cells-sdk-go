# Cells API client

Legacy Rest API Client in go for Pydio Cells. Please see v3 subfolder.

## WARNING LEGACY BRANCH

This version has been released with version 0.9 before cleaning and refactoring the API in order to provide backward compatibility for potential consumers of the SDK.

You should really rather use [the v3 branch](https://github.com/pydio/cells-sdk-go/blob/master/v3/README.md)

## How to use

Simply put the package under your project folder and add the following import:

```go
    "github.com/pydio/cells-sdk-go"
```

The `transport` sub-package provides some utilitary methods to ease the set up of a communication with your target Cells instance, and you might find the few sample commands that are in the `example` package useful to jump in.

You can also have a look at the [Cells client repository](https://github.com/pydio/cells-client) (also on Github) to see more working examples.
