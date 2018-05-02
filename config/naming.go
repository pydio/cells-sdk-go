package config

import (
	hashiversion "github.com/hashicorp/go-version"
)

var (
	// PackageType stores and exposes current package type.
	PackageType string
	// PackageLabel stores and exposes current package label.
	PackageLabel string

	// The 3 below vars are initialized by the go linker directly
	// in the resulting binary when doing 'make main'
	version = "0.1.0"
	// BuildStamp stores the current build stamp.
	BuildStamp string
	// BuildRevision stores the current build version.
	BuildRevision string
)

// Version retrieves the current build version of the app. Note that this is updated via the linker upon build.
func Version() *hashiversion.Version {
	v, _ := hashiversion.NewVersion(version)
	return v
}
