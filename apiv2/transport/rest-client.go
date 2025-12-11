package transport

import (
	"context"
	"net/url"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/apiv2"
	v2Client "github.com/pydio/cells-sdk-go/v5/apiv2/client"
)

func GetRuntimeTransport(context context.Context, currConfig *apiv2.SdkConfig) (runtime.ClientTransport, error) {
	u, e := url.Parse(currConfig.Url)
	if e != nil {
		return nil, e
	}

	tp := client.New(u.Host, getApiPrefix(currConfig), []string{u.Scheme})
	transportOptions := []any{
		WithSkipVerify(currConfig.SkipVerify),
		WithCustomHeaders(currConfig.CustomHeaders),
	}
	tp.Context = context
	tp.Transport = New(transportOptions...)
	return tp, nil
}

func GetRestClient(currConfig *apiv2.SdkConfig, anon bool) (*v2Client.PydioCellsRestAPI, error) {
	t, e := GetClientTransport(currConfig, anon)
	if e != nil {
		return nil, e
	}
	return v2Client.New(t, strfmt.Default), nil
}

func GetClientTransport(currConfig *apiv2.SdkConfig, anonymous bool) (runtime.ClientTransport, error) {
	u, e := url.Parse(currConfig.Url)
	if e != nil {
		return nil, e
	}

	tp := client.New(u.Host, getApiPrefix(currConfig), []string{u.Scheme})
	options := []any{
		WithSkipVerify(currConfig.SkipVerify),
		WithCustomHeaders(currConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := TokenProviderFromConfig(currConfig)
		if e != nil {
			return nil, e
		}
		options = append(options, WithBearer(tp))
	}
	tp.Context = context.Background()
	tp.Transport = New(options...)

	return tp, nil
}

// getApiPrefix always use the default CellsApiPrefix for runtime transports.
func getApiPrefix(currConfig *apiv2.SdkConfig) string {
	if currConfig.ApiResourcePrefix != "" {
		return currConfig.ApiResourcePrefix
	}
	return CellsApiPrefix
}

