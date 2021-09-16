package rest

import (
	"context"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cells_sdk "github.com/pydio/cells-sdk-go/v3"
	client2 "github.com/pydio/cells-sdk-go/v3/client"
	"github.com/pydio/cells-sdk-go/v3/transport"
)

func GetClient(sdkConfig *cells_sdk.SdkConfig, anon bool) (context.Context, *client2.PydioCellsRest, error) {

	c, t, e := GetClientTransport(sdkConfig, anon)
	if e != nil {
		return nil, nil, e
	}
	return c, client2.New(t, strfmt.Default), nil

}

func GetClientTransport(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (context.Context, runtime.ClientTransport, error) {

	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, nil, e
	}
	tp := client.New(u.Host, transport.CellsApiResourcePath, []string{u.Scheme})
	transportOptions := []interface{}{
		transport.WithSkipVerify(sdkConfig.SkipVerify),
		transport.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, nil, e
		}
		transportOptions = append(transportOptions, transport.WithBearer(tp))
	}
	tp.Context = context.Background()
	tp.Transport = transport.New(transportOptions...)

	return tp.Context, tp, nil
}
