# Cells API client

Rest API Client for Pydio Cells.

Temporary note: SDK has been updated on **Jun. 17th 2020** for 2.1.0-rc0 to follow the new development cycle of Cells 2.1+ that is not yet released.

If you are using this in production and until we release Cells 2.1.0, we strongly advise that you stick to previous **2.0.x** revision, typically, with go modules:

```go
github.com/pydio/cells-sdk-go v0.0.0-20200313145255-e8831efc2d1b
```

Where SDK has been updated on **Mar. 11th 2020** with git commit **[07dc255](https://github.com/pydio/cells/commit/07dc2557e3a7099e8bb3b6c7b0dcb4ac80d31601)** that works with **Cells v2.0.9** and previous releases.

## Overview

This API client was generated by the [go-swagger](https://github.com/go-swagger/go-swagger) project.

For more information, please visit [https://pydio.com](https://pydio.com)

## How to use

Simply put the package under your project folder and add the following import:

```go
    "github.com/pydio/cells-sdk-go"
```

The `transport` sub-package provides some utilitary methods to ease the set up of a communication with your target Cells instance, and you might find the few sample commands that are in the `example` package useful to jump in.

You can also have a look at the [Cells client repository](https://github.com/pydio/cells-client) (also on Github) to see more working examples.
