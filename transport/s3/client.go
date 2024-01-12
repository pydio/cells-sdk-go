/*
 * Copyright (c) 2019. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

package s3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	cells_sdk "github.com/pydio/cells-sdk-go/v4"
	"github.com/pydio/cells-sdk-go/v4/transport"
	http2 "github.com/pydio/cells-sdk-go/v4/transport/http"
)

// // GetClient creates and configure a new S3 client at each request.
// func GetClient(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.S3, error) {

// 	tp, e := transport.TokenProviderFromConfig(sdc)
// 	if e != nil {
// 		return nil, e
// 	}
// 	htCl := http2.GetClient(sdc)
// 	conf := aws.NewConfig()
// 	conf.WithCredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp)))
// 	conf.WithHTTPClient(htCl)
// 	conf.WithEndpoint(s3c.Endpoint)
// 	conf.WithRegion(s3c.Region)
// 	s3Session, err := session.NewSession(conf)
// 	if err != nil {
// 		log.Fatal("cannot initialise default S3 session: " + err.Error())
// 	}
// 	return s3.New(s3Session), nil

// }

func GetClient(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.Client, error) {
	htCl := http2.GetClient(sdc)
	tp, err := transport.TokenProviderFromConfig(sdc)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(s3c.Region),
		config.WithHTTPClient(htCl),
		config.WithCredentialsProvider(AsS3CredentialsProvider(tp)),
	)
	if err != nil {
		log.Fatal("cannot load default S3 session configuration:", err.Error())
	}

	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg,
		func(o *s3.Options) {
			o.BaseEndpoint = aws.String(s3c.Endpoint)
		}), nil
}

// // LoadS3Config retrieves current S3 configuration to start a new client
// func LoadS3Config(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (aws.Config, error) {
// 	// tp, e := transport.TokenProviderFromConfig(sdc)
// 	// if e != nil {
// 	// 	return nil, e
// 	// }
// 	// htCl := http2.GetClient(sdc)

// 	// conf := aws.NewContransportfig()
// 	// conf.WithCredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp)))
// 	// conf.WithHTTPClient(htCl)
// 	// conf.WithEndpoint(s3c.Endpoint)
// 	// conf.WithRegion(s3c.Region)

// 	// s3Session, err := session.NewSession(conf)
// 	// if err != nil {
// 	// 	log.Fatal("cannot initialise default S3 session: " + err.Error())
// 	// }

// 	htCl := http2.GetClient(sdc)
// 	cfg, err := config.LoadDefaultConfig(
// 		context.TODO(),
// 		config.WithRegion(s3c.Region),
// 		// config.WithCtransportredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp))),
// 		config.WithHTTPClient(htCl),
// 		// config.WithEndpoint(s3c.Endpoint),
// 	)
// 	if err != nil {
// 		log.Fatal("cannot load default S3 session configuration:", err.Error())
// 	}
// 	return cfg, nil

// }
