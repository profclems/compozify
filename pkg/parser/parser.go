package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errNoMoreFlags = errors.New("no more flags available")
	errInvalidFlag = errors.New("invalid docker run flag")
	errSkipFlag    = errors.New("skip flag")
)

// Parser parses a docker run command into a docker compose file format.
type Parser struct {
	document *yaml.Node
	version  string

	refs    map[string]*yaml.Node
	vars    *variables
	command []string

	yamlBytes []byte
}

// SetVersion sets the docker compose version.
func (p *Parser) SetVersion(v string) {
	p.version = v
}

// NewParser creates a new Parser.
func NewParser(s string) (*Parser, error) {
	s = strings.TrimPrefix(strings.TrimSpace(s), "docker run")
	if s == "" {
		return nil, errors.New("empty docker command")
	}

	p := &Parser{
		version: composeVersion,
		refs:    make(map[string]*yaml.Node),
		vars:    newVariables(),
	}

	command, err := parseArgs(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse docker run command: %w", err)
	}
	p.command = command

	containerTitleNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: defaultServiceName,
	}

	containerNode := &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: []*yaml.Node{},
	}

	servicesNode := &yaml.Node{
		Kind: yaml.MappingNode,
		Content: []*yaml.Node{
			containerTitleNode,
			containerNode,
		},
	}

	p.document = &yaml.Node{
		Kind: yaml.DocumentNode,
		Content: []*yaml.Node{
			{
				Kind: yaml.MappingNode,
				Content: []*yaml.Node{
					{
						Kind:  yaml.ScalarNode,
						Value: "version",
					},
					{
						Kind:  yaml.ScalarNode,
						Value: p.version,
						Style: yaml.DoubleQuotedStyle,
					},
					{
						Kind:  yaml.ScalarNode,
						Value: "services",
					},
					servicesNode,
				},
			},
		},
	}

	p.refs["^services"] = servicesNode
	p.refs["$service"] = containerNode
	p.refs["$serviceTitleNode"] = containerTitleNode

	return p, nil
}

// Parse parses the docker run command into a docker compose file format.
func (p *Parser) Parse() error {
	var parseErr error
	for {
		flag, value, err := p.parseOneFlag()
		if err != nil {
			if errors.Is(err, errSkipFlag) {
				continue
			}
			parseErr = err
			break
		}

		if flag == "" {
			break
		}

		if dockerFlag := p.vars.Get(flag); dockerFlag != nil {
			if dockerFlag.ComposeName == "" {
				continue
			}

			composePath := strings.Split(dockerFlag.ComposeName, ".")

			parent := p.document
			for len(composePath) > 0 {
				key := composePath[0]
				composePath = composePath[1:]
				val := value

				kind := dockerFlag.Type

				cNode := p.refs[key]
				if ftype := p.vars.GetType(key); !ftype.IsZero() {
					kind = ftype
				}
				if cNode == nil {
					cNode, err = p.addNode(parent, flag, key, val, kind)
					if err != nil {
						return err
					}
				}
				parent = cNode
			}
		}

		// TODO: what do we do with an unknown flag?
	}

	if errors.Is(parseErr, errNoMoreFlags) {
		parseErr = p.parseImage()
	}

	if parseErr != nil {
		return parseErr
	}

	p.yamlBytes, parseErr = yaml.Marshal(p.document)
	return parseErr
}

func trimQuotes(s string) string {
	return strings.Trim(s, `"'`)
}

func (p *Parser) addNode(parent *yaml.Node, flag, key, value string, ftype FlagType) (*yaml.Node, error) {
	kind := ftype.YamlKind()
	valueNode := &yaml.Node{}
	mainKey := key

	value = trimQuotes(value)

	if key == "$var" {
		key = ""
		switch ftype {
		case MapType:
			kind = yaml.ScalarNode
			vals := strings.SplitN(value, "=", 2)
			switch len(vals) {
			case 1:
				key = vals[0]
				value = ""
			case 2:
				key, value = vals[0], trimQuotes(vals[1])
			default:
				return nil, fmt.Errorf("invalid value %s for docker run flag %q", value, flag)
			}
		case ArrayType:
			kind = yaml.ScalarNode
		case UlimitType:
			ulimit, err := ParseUlimit(value)
			if err != nil {
				return nil, err
			}
			key, valueNode = ulimit.YAML()
		case MountType:
			mount, err := ParseMount(value)
			if err != nil {
				return nil, err
			}
			key, valueNode = mount.YAML()
		}
	}

	valueNode.Value = value
	valueNode.Kind = kind

	if key != "" {
		parent.Content = append(parent.Content, &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: key,
		})

		if mainKey != "$var" {
			p.refs[key] = valueNode
		}
	}

	parent.Content = append(parent.Content, valueNode)

	return valueNode, nil
}

func (p *Parser) parseImage() error {
	image := p.command[0]
	imageNode := []*yaml.Node{
		{
			Kind:  yaml.ScalarNode,
			Value: "image",
		},
		{
			Kind:  yaml.ScalarNode,
			Value: image,
		},
	}
	// if the image name is for example, profclems/glab:latest or just profclems/glab
	//  we need to make sure the service name will be the last word after slash but without the
	//  tag version, like just "glab" in the example above
	p.command = p.command[1:] // the rest are commands
	ns := strings.Split(image, "/")
	serviceName := strings.SplitN(ns[len(ns)-1], ":", 2)[0]

	p.refs["$serviceTitleNode"].Value = serviceName
	p.refs["$service"].Content = append(p.refs["$service"].Content, imageNode...)

	if len(p.command) > 0 {
		commandsNode := &yaml.Node{
			Kind: yaml.SequenceNode,
		}

		for len(p.command) > 0 {
			c := p.command[0]
			commandsNode.Content = append(commandsNode.Content, &yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: c,
			})

			p.command = p.command[1:]
		}

		p.refs["$service"].Content = append(p.refs["$service"].Content, &yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: "command",
		}, commandsNode)
	}

	return nil
}

func (p *Parser) parseOneFlag() (string, string, error) {
	if len(p.command) == 0 {
		return "", "", nil
	}

	f := p.command[0]
	if len(f) > 0 && f[0] != '-' {
		if f[0] == '\\' {
			// shows a line break, skip
			p.command = p.command[1:]
			return "", "", errSkipFlag
		}
		// this could possibly be the image name and not a flag
		return "", "", errNoMoreFlags
	}

	p.command = p.command[1:]

	if len(f) < 2 {
		return "", "", errInvalidFlag
	}
	numOfMinuses := 1
	if f[1] == '-' {
		numOfMinuses++

		if len(f) == 2 {
			// "--" indicates flag termination
			return "", "", errNoMoreFlags
		}
	}

	name := f[numOfMinuses:]
	if len(name) == 0 || name == "-" || name == "=" {
		return "", "", errNoMoreFlags
	}

	// check if flag has value attached
	value := ""
	hasValue := false
	varType := p.vars.GetVarType(name)
	if numOfMinuses == 1 && len(name) > 1 && name[1] != '=' {
		// if the flag is a shorthand and not a boolean the value can be directly
		//  attached to the flag name like -p80:80 and -uroot
		fname := name[:1]
		varType = p.vars.GetVarType(fname)
		if varType == BoolType {
			p.command = append([]string{"-" + name[1:]}, p.command...)
			name = fname
			value = "true"
		} else {
			value = name[1:]
			name = fname
		}
		hasValue = true
	}
	for i := 0; i < len(name); i++ {
		if name[i] == '=' {
			hasValue = true
			value, name = name[i+1:], name[:i]
			varType = p.vars.GetVarType(name)
			break
		}
	}

	// check if flag is a boolean flag
	if varType == BoolType {
		if hasValue {
			pv, err := strconv.ParseBool(value)
			if err != nil {
				return "", "", fmt.Errorf("invalid value %q for docker run flag %q: %s", value, name, err)
			}
			value = strconv.FormatBool(pv)
			return name, value, nil
		}
		// What if the next argument is a value for the flag.
		// We are checking with >1 because if for example `docker run -t false`
		// docker assumes "false" to be the image name and -t will be true but if you run
		// `docker run -t false false`, t will be set to false and the image
		// name will be false.
		if len(p.command) > 1 {
			arg := p.command[0]
			if arg[0] == '-' {
				// next arg is also a flag, so we can assume the value is true
				value = "true"
			} else if pv, err := strconv.ParseBool(arg); err == nil {
				value = strconv.FormatBool(pv)
				p.command = p.command[1:]
			} else {
				value = "true"
			}
		} else {
			value = "true"
		}
	} else {
		if !hasValue {
			// next argument is definitely the value
			if len(p.command) > 0 {
				arg := p.command[0]
				// TODO: is it really? what about `docker run --label -myLabelWhichHasAHyphenIntheBeginning image`
				if arg[0] != '-' {
					value, p.command = arg, p.command[1:]
					hasValue = true
				}
			}
		}

		if !hasValue {
			return "", "", fmt.Errorf("docker run flag %q is missing an argument", name)
		}
	}

	return name, value, nil
}

func (p *Parser) Bytes() []byte {
	return p.yamlBytes
}

func (p *Parser) String() string {
	return string(p.Bytes())
}
