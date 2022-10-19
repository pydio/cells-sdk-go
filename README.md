# Cells API client

Go SDK client for Pydio Cells.

**WARNING** the 1.x branch of the curent repository (that is shown here) is **_legacy_**.
You should really use version 4+ that is up to date in `./v4` folder. 

Please see [the corresponding ReadMe](https://github.com/pydio/cells-sdk-go/blob/master/v4/README.md).

## Legacy ReadMe

This version has been released with version 0.9 before cleaning and refactoring the API in order to provide backward compatibility for potential consumers of the SDK.

## How to use

Simply put the package under your project folder and add the following import:

```go
    "github.com/pydio/cells-sdk-go"
```

The `transport` sub-package provides some utilitary methods to ease the set up of a communication with your target Cells instance, and you might find the few sample commands that are in the `example` package useful to jump in.

You can also have a look at the [Cells client repository](https://github.com/pydio/cells-client) (also on Github) to see more working examples.
