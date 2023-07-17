package version

import (
	"fmt"
	"runtime/debug"
)

var (
	// Version is the version of the binary set at build time.
	// Do not use this directly, use GetVersion() instead.
	Version = "DEV"
	// BuildDate is the date of the build set at build time.
	// Do not use this directly, use GetVersion() instead.
	BuildDate = "Unknown"
)

// Info is the version of the binary.
type Info struct {
	version string
	build   string
}

func GetVersion() Info {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}

	return Info{Version, BuildDate}
}

func (v Info) String() string {
	return fmt.Sprintf("%s (%s)", v.version, v.build)
}

func (v Info) IsDev() bool {
	return v.version == "DEV"
}
