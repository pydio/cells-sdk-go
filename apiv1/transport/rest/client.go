package rest

import (
	"net/url"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/apiv1"
	v1client "github.com/pydio/cells-sdk-go/v5/apiv1/client"
	v1transport "github.com/pydio/cells-sdk-go/v5/apiv1/transport"
	v1http "github.com/pydio/cells-sdk-go/v5/apiv1/transport/http"
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
		v1http.WithSkipVerify(sdkConfig.SkipVerify),
		v1http.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := v1transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		options = append(options, v1http.WithBearer(tp))
	}
	tp.Transport = v1transport.New(options...)
	return tp, nil
}
