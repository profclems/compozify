package version

// Info is the version of the binary.
type Info string

func GetVersion() Info {
	return Info(version)
}

func (v Info) String() string {
	return string(v)
}

func (v Info) IsDev() bool {
	return v.String() == "DEV"
}

// Version is the version of the binary injected in compilation time.
var version = "DEV"
