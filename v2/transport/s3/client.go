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
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	cells_sdk "github.com/pydio/cells-sdk-go/v2"
	"github.com/pydio/cells-sdk-go/v2/transport"
	http2 "github.com/pydio/cells-sdk-go/v2/transport/http"
)

// GetClient creates and configure a new S3 client at each request.
func GetClient(sdc *cells_sdk.SdkConfig, s3c *cells_sdk.S3Config) (*s3.S3, error) {

	tp, e := transport.TokenProviderFromConfig(sdc)
	if e != nil {
		return nil, e
	}
	htCl := http2.GetClient(sdc)
	conf := aws.NewConfig()
	conf.WithCredentials(credentials.NewCredentials(AsS3CredentialsProvider(tp)))
	conf.WithHTTPClient(htCl)
	conf.WithEndpoint(s3c.Endpoint)
	conf.WithRegion(s3c.Region)
	s3Session, err := session.NewSession(conf)
	if err != nil {
		log.Fatal("cannot initialise default S3 session: " + err.Error())
	}
	return s3.New(s3Session), nil

}
