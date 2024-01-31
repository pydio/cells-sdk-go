package rest

import (
	"context"
	"github.com/pydio/cells-sdk-go/v5/transport/http"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cells_sdk "github.com/pydio/cells-sdk-go/v5"
	client2 "github.com/pydio/cells-sdk-go/v5/client"
	"github.com/pydio/cells-sdk-go/v5/transport"
)

func GetApiClient(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (*client2.PydioCellsRestAPI, error) {
	apiClient, err := GetApiTransport(sdkConfig, anonymous)
	if err != nil {
		return nil, err
	}
	return client2.New(apiClient, strfmt.Default), nil
}

func GetApiTransport(sdkConfig *cells_sdk.SdkConfig, anonymous bool) (*client.Runtime, error) {
	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, transport.CellsApiResourcePath, []string{u.Scheme})
	transportOptions := []interface{}{
		http.WithSkipVerify(sdkConfig.SkipVerify),
		http.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		transportOptions = append(transportOptions, http.WithBearer(tp))
	}
	tp.Transport = transport.New(transportOptions...)
	return tp, nil
}

func GetClient(sdkConfig *cells_sdk.SdkConfig, anon bool) (context.Context, *client2.PydioCellsRestAPI, error) {

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
		http.WithSkipVerify(sdkConfig.SkipVerify),
		http.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, nil, e
		}
		transportOptions = append(transportOptions, http.WithBearer(tp))
	}
	tp.Context = context.Background()
	tp.Transport = transport.New(transportOptions...)

	return tp.Context, tp, nil
}
