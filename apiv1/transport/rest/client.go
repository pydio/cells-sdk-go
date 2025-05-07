package rest

import (
	cellssdk "github.com/pydio/cells-sdk-go/v5/apiv1"
	sdkclient "github.com/pydio/cells-sdk-go/v5/apiv1/client"
	"github.com/pydio/cells-sdk-go/v5/apiv1/transport"
	http2 "github.com/pydio/cells-sdk-go/v5/apiv1/transport/http"
	"net/url"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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
	tp := client.New(u.Host, cellssdk.CellsApiResourcePath, []string{u.Scheme})
	transportOptions := []interface{}{
		http2.WithSkipVerify(sdkConfig.SkipVerify),
		http2.WithCustomHeaders(sdkConfig.CustomHeaders),
	}
	if !anonymous {
		tp, e := transport.TokenProviderFromConfig(sdkConfig)
		if e != nil {
			return nil, e
		}
		transportOptions = append(transportOptions, http2.WithBearer(tp))
	}
	tp.Transport = transport.New(transportOptions...)
	return tp, nil
}
