package rest

import (
	"net/url"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	cellssdk "github.com/pydio/cells-sdk-go/v5"
	sdkclient "github.com/pydio/cells-sdk-go/v5/client"
	"github.com/pydio/cells-sdk-go/v5/transport"
	"github.com/pydio/cells-sdk-go/v5/transport/http"
)

func GetApiClient(sdkConfig *cellssdk.SdkConfig, anonymous bool) (*sdkclient.PydioCellsRestAPI, error) {
	apiClient, err := GetApiTransport(sdkConfig, anonymous)
	if err != nil {
		return nil, err
	}
	return sdkclient.New(apiClient, strfmt.Default), nil
}

func GetApiTransport(sdkConfig *cellssdk.SdkConfig, anonymous bool) (*client.Runtime, error) {
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
