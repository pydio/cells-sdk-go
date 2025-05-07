package transport

import (
	"context"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cellsSdk "github.com/pydio/cells-sdk-go/v5"
	sdkClient "github.com/pydio/cells-sdk-go/v5/client"
)

func GetRuntimeTransport(context context.Context, currConfig *cellsSdk.SdkConfig) (runtime.ClientTransport, error) {
	u, e := url.Parse(currConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, CellsApiPrefix, []string{u.Scheme})
	transportOptions := []interface{}{
		WithSkipVerify(currConfig.SkipVerify),
		WithCustomHeaders(currConfig.CustomHeaders),
	}
	tp.Context = context
	tp.Transport = New(transportOptions...)
	return tp, nil
}

func GetRestClient(currConfig *cellsSdk.SdkConfig, anon bool) (*sdkClient.PydioCellsRestAPI, error) {
	t, e := GetClientTransport(currConfig, anon)
	if e != nil {
		return nil, e
	}
	return sdkClient.New(t, strfmt.Default), nil
}

func GetClientTransport(currConfig *cellsSdk.SdkConfig, anonymous bool) (runtime.ClientTransport, error) {

	u, e := url.Parse(currConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, CellsApiPrefix, []string{u.Scheme})
	transportOptions := []interface{}{
		WithSkipVerify(currConfig.SkipVerify),
		WithCustomHeaders(currConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := TokenProviderFromConfig(currConfig)
		if e != nil {
			return nil, e
		}
		transportOptions = append(transportOptions, WithBearer(tp))
	}
	tp.Context = context.Background()
	tp.Transport = New(transportOptions...)

	return tp, nil
}
