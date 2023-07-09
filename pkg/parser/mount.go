package parser

import (
	"gopkg.in/yaml.v3"
)

type Mount struct {
	Name        string // name in docker-compose
	Source      string
	Target      string
	Type        string
	ReadOnly    bool
	Consistency string
	NodeType    FlagType

	Bind struct {
		Propagation    string
		CreateHostPath bool
		Selinux        bool
		NodeType       FlagType
	}

	Volume struct {
		NoCopy   bool
		NodeType FlagType
	}

	Tmpfs struct {
		Size     string
		Mode     string
		NodeType FlagType
	}
}

// ParseMount converts docker run mount format to docker-compose mount format
// into the Mount struct
func ParseMount(s string) (*Mount, error) {
	mount := &Mount{}
	//mount.NodeType = MapType
	//
	//if s == "" {
	//	return nil, errInvalidFlag
	//}
	//
	//mountSplit := strings.Split(s, ",")
	//
	//for _ := range mountSplit {
	//
	//}

	return mount, nil
}

// YAML converts the Mount struct to a yaml.Node
func (m *Mount) YAML() (key, value *yaml.Node) {
	return key, value
}
