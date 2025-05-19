package rest

import (
	"net/url"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cellssdk "github.com/pydio/cells-sdk-go/v4"
	sdkclient "github.com/pydio/cells-sdk-go/v4/client"
	"github.com/pydio/cells-sdk-go/v4/transport"
	"github.com/pydio/cells-sdk-go/v4/transport/http"
)

func GetApiClient(sdkConfig *cellssdk.SdkConfig, anonymous bool) (*sdkclient.PydioCellsRestAPI, error) {
	apiClient, err := GetApiRuntime(sdkConfig, anonymous)
	if err != nil {
		return nil, err
	}
	return sdkclient.New(apiClient, strfmt.Default), nil
}

func GetApiRuntime(sdkConfig *cellssdk.SdkConfig, anonymous bool) (*client.Runtime, error) {
	u, e := url.Parse(sdkConfig.Url)
	if e != nil {
		return nil, e
	}
	tp := client.New(u.Host, cellssdk.CellsApiResourcePath, []string{u.Scheme})

	options := []any{
		http.WithSkipVerify(sdkConfig.SkipVerify),
		http.WithCustomHeaders(sdkConfig.CustomHeaders),
	}

	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		options = append(options, http.WithBearer(tp))
	}
	tp.Transport = transport.New(options...)
	return tp, nil
}
