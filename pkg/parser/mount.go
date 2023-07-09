package parser

import (
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
)

type bindOpts struct {
	Propagation    string
	CreateHostPath bool
	Selinux        bool
}

type volumeOpts struct {
	NoCopy bool
}

type tmpfsOpts struct {
	Size string
	Mode string
}

// Mount represents a docker run mount flag
type Mount struct {
	Source      string
	Target      string
	Type        string
	ReadOnly    bool
	Consistency string
	NodeType    FlagType

	Bind   bindOpts
	Volume volumeOpts
	Tmpfs  tmpfsOpts
}

// ParseMount converts docker run mount format to docker-compose mount format
// into the Mount struct
// mount value format: --mount type=bind,source=/tmp,target=/tmp,readonly
func ParseMount(s string) (*Mount, error) {
	mount := &Mount{}
	mount.NodeType = ArrayType

	if s == "" {
		return nil, errInvalidFlag
	}

	mountSplit := strings.Split(s, ",")

	for _, s := range mountSplit {
		split := strings.SplitN(s, "=", 2)
		if len(split) < 1 {
			return nil, errInvalidFlag
		}

		if len(split) == 1 {
			if split[0] == "readonly" {
				mount.ReadOnly = true
				continue
			}
			return nil, errInvalidFlag
		}

		switch split[0] {
		case "type":
			mount.Type = split[1]
		case "source":
			mount.Source = split[1]
		case "target":
			mount.Target = split[1]
		case "readonly":
			mount.ReadOnly = true
		case "consistency":
			mount.Consistency = split[1]
		case "bind-propagation":
			mount.Bind.Propagation = split[1]
		case "volume-opt":
			split2 := strings.SplitN(split[1], "=", 2)
			if len(split2) < 1 {
				return nil, errInvalidFlag
			}
			switch split2[0] {
			case "nocopy":
				val := true
				if len(split) == 2 {
					val, _ = strconv.ParseBool(split2[1])
				}
				mount.Volume.NoCopy = val
			}
		case "tmpfs-size":
			mount.Tmpfs.Size = split[1]
		case "tmpfs-mode":
			mount.Tmpfs.Mode = split[1]
		}
	}

	return mount, nil
}

// YAML converts the Mount struct to a yaml.Node
func (m *Mount) YAML() (key, value *yaml.Node) {
	value = &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: []*yaml.Node{},
	}

	if m.Type != "" {
		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "type",
			}, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: m.Type,
			},
		)
	}

	if m.Source != "" {
		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "source",
			}, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: m.Source,
			},
		)
	}

	if m.Target != "" {
		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "target",
			}, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: m.Target,
			},
		)
	}

	if m.ReadOnly {
		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "readonly",
			}, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: strconv.FormatBool(m.ReadOnly),
			},
		)
	}

	if m.Consistency != "" {
		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "consistency",
			}, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: m.Consistency,
			},
		)
	}

	if m.Bind != (bindOpts{}) {
		bindNode := &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: m.Bind.Propagation,
		}

		value.Content = append(value.Content,
			&yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: "bind",
			}, bindNode,
		)

		if m.Bind.CreateHostPath {

		}
	}

	return key, value
}
