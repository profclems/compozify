package parser

import (
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

// Mount represents a docker run mount flag
type Mount struct {
	Type        string `name:"type" compose:"type" compose-type:"string"`
	Source      string `name:"source" compose:"source" compose-type:"string"`
	Target      string `name:"target" compose:"target" compose-type:"string"`
	Readonly    string `name:"readonly" compose:"read_only" compose-type:"bool"`
	Consistency string `name:"consistency" compose:"consistency" compose-type:"string"`

	// bind-* options
	BindPropagation    string `name:"bind-propagation" compose:"bind.propagation" compose-type:"string"`
	BindCreatehostpath string `name:"-" compose:"bind.create_host_path" compose-type:"bool"`
	BindSelinux        string `name:"-" compose:"bind.selinux" compose-type:"string"`

	// TODO: properly support volume mounts
	// volume-* options
	VolumeNocopy string `name:"-" compose:"volume.nocopy" compose-type:"bool"`

	// tmpfs-* options
	TmpfsSize string `name:"tmpfs-size" compose:"tmpfs.size" compose-type:"string"`
	TmpfsMode string `name:"tmpfs-mode" compose:"tmpfs.mode" compose-type:"string"`
}

// ParseMount converts docker run mount format to docker-compose mount format
// into the Mount struct
// mount value format: --mount type=bind,source=/tmp,target=/tmp,readonly
func ParseMount(s string) (*Mount, error) {
	altNames := map[string]string{
		"src":         "Source",
		"dst":         "Target",
		"destination": "Target",
	}
	mount := &Mount{}

	if s == "" {
		return nil, errInvalidFlag
	}

	mountSplit := strings.Split(s, ",")

	for _, s := range mountSplit {
		split := strings.SplitN(s, "=", 2)

		if len(split) < 1 {
			return nil, errInvalidFlag
		}

		key := split[0]
		if altName, ok := altNames[key]; ok {
			key = altName
		}

		fieldName := strings.ToLower(key)
		options := strings.Split(fieldName, "-")
		if len(options) > 0 {
			fieldName = ""
			for len(options) > 0 {
				str := strings.ToUpper(options[0][:1])
				fieldName += str + options[0][1:]
				options = options[1:]
			}
		}
		mountElem := reflect.ValueOf(mount).Elem()

		field := mountElem.FieldByName(fieldName)
		if len(split) == 1 {
			f, ok := mountElem.Type().FieldByName(fieldName)
			if ok && f.Tag.Get("compose-type") == "bool" { // TODO: read type tag
				field.Set(reflect.ValueOf("true"))
				continue
			}
			return nil, errInvalidFlag
		}

		value := split[1]
		if field.CanSet() {
			field.Set(reflect.ValueOf(value))
		}
	}

	return mount, nil
}

// YAML converts the Mount struct to a yaml.Node
func (m *Mount) YAML() (key string, value *yaml.Node) {
	value = &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: []*yaml.Node{},
	}

	mountElem := reflect.ValueOf(m).Elem()
	mapNode := make(map[string]*yaml.Node)

	for i := 0; i < mountElem.NumField(); i++ {
		field := mountElem.Field(i)
		if !field.CanSet() {
			continue
		}

		fieldName := mountElem.Type().Field(i).Name
		if fieldName == "NodeType" {
			continue
		}

		fieldTag := mountElem.Type().Field(i).Tag
		fieldName = fieldTag.Get("compose")
		if fieldName == "" {
			fieldName = fieldTag.Get("name")
		}

		_ = fieldTag.Get("compose-type")

		fieldValue := field.Interface()
		if fieldValue == "" {
			continue
		}

		fieldNameArr := strings.Split(fieldName, ".")
		parent := value
		for i := 0; i < len(fieldNameArr); i++ {
			fieldName := fieldNameArr[i]
			if i < len(fieldNameArr)-1 {
				if p, ok := mapNode[fieldName]; ok {
					parent = p
					continue
				}
				parent = &yaml.Node{
					Kind:    yaml.MappingNode,
					Content: []*yaml.Node{},
				}
				value.Content = append(value.Content, &yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: fieldName,
				}, parent)
				mapNode[fieldName] = parent
				continue
			}
			fieldValueNode := &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: fieldValue.(string),
			}
			parent.Content = append(parent.Content, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: fieldName,
			}, fieldValueNode)
		}
	}

	return "", value
}
