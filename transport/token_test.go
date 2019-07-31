// +build ignore

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

package transport

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pydio/cells-sdk-go/transport/oidc"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	rootPath string
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestJWT(t *testing.T) {

	if RunEnvAwareTests {
		Convey("Given a clean environment than can access a running Cells instance", t, func() {

			Convey("Set up default env aware configuration", func() {
				crp, err := filepath.Abs(".")
				So(err, ShouldBeNil)

				rootPath = filepath.Dir(crp)
				SetUpEnvironment(GetDefaultConfigFiles(rootPath))

				Convey("Parameters have been set", func() {
					// fmt.Printf("## DEFAULT CONF \n%s - %s - %s - %s - %s - %s - %v\n", DefaultConfig.Protocol, DefaultConfig.Url, DefaultConfig.ClientKey, DefaultConfig.ClientSecret, DefaultConfig.User, DefaultConfig.Password, DefaultConfig.SkipVerify)
					So(DefaultConfig.Url, ShouldNotBeEmpty)
					So(DefaultConfig.ClientKey, ShouldNotBeEmpty)
					So(DefaultConfig.ClientSecret, ShouldNotBeEmpty)
					So(DefaultConfig.User, ShouldNotBeEmpty)
					So(DefaultConfig.Password, ShouldNotBeEmpty)
					So(DefaultConfig.SkipVerify, ShouldNotBeEmpty)
					sdkc, s3c := GetDefaultConfigFiles(rootPath)
					So(sdkc, ShouldNotBeNil)
					So(s3c, ShouldNotBeNil)
				})

				Convey("JWT token must be retrievable and longer than 256 chars.", func() {
					jwt, err := oidc.RetrieveToken(DefaultConfig)
					So(err, ShouldBeNil)
					So(len(jwt), ShouldBeGreaterThan, 256)
				})
			})
		})
	}
}
