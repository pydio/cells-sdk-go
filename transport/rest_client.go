package transport

import (
	"context"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	client2 "github.com/pydio/cells-sdk-go/v5/client"
)

func GetRestClient(sdkConfig *cells_sdk.SdkConfig, anon bool) (*client2.PydioCellsRestAPI, error) {
	t, e := GetClientTransport(sdkConfig, anon)
	if e != nil {
		return nil, e
	}
	return client2.New(t, strfmt.Default), nil
}

func GetClientTransport(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (runtime.ClientTransport, error) {

	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, CellsApiResourcePath, []string{u.Scheme})
	transportOptions := []interface{}{
		WithSkipVerify(sdkConfig.SkipVerify),
		WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		transportOptions = append(transportOptions, WithBearer(tp))
	}
	tp.Context = context.Background()
	tp.Transport = New(transportOptions...)

	return tp, nil
}
