package rest

import (
	"github.com/pydio/cells-sdk-go/v5/apiv1"
	v1client "github.com/pydio/cells-sdk-go/v5/apiv1/client"
	"github.com/pydio/cells-sdk-go/v5/apiv1/transport"
	http2 "github.com/pydio/cells-sdk-go/v5/apiv1/transport/http"
	"net/url"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func GetApiClient(sdkConfig *apiv1.SdkConfig, anonymous bool) (*v1client.PydioCellsRestAPI, error) {
	apiClient, err := GetApiRuntime(sdkConfig, anonymous)
	if err != nil {
		return nil, err
	}
	return v1client.New(apiClient, strfmt.Default), nil
}

func GetApiRuntime(sdkConfig *apiv1.SdkConfig, anonymous bool) (*client.Runtime, error) {
	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, apiv1.CellsApiResourcePath, []string{u.Scheme})
	options := []any{
		http2.WithSkipVerify(sdkConfig.SkipVerify),
		http2.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		options = append(options, http2.WithBearer(tp))
	}
	tp.Transport = transport.New(options...)
	return tp, nil
}
