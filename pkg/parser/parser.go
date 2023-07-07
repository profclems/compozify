package parser

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"strconv"
	"strings"
)

type Parser struct {
	document *yaml.Node

	refs    map[string]*yaml.Node
	command []string
}

func NewParser(s string) (*Parser, error) {
	s = strings.TrimPrefix(strings.TrimSpace(s), "docker run")
	if s == "" {
		return nil, errors.New("empty docker command")
	}

	p := &Parser{
		command: strings.Fields(s),
		refs:    make(map[string]*yaml.Node),
	}

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
						Value: composeVersion,
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

	p.refs["services"] = servicesNode
	p.refs["$service"] = containerNode
	p.refs["$serviceTitleNode"] = containerTitleNode

	return p, nil
}

func (p *Parser) Parse() error {
	var parseErr error
	for {
		flag, value, err := p.parseOneFlag()
		if err != nil || flag == "" {
			parseErr = err
			break
		}

		dockerFlag, ok := dockerRunFlags[flag]
		if ok {
			if ref := dockerFlag.Reference; ref != "" {
				dockerFlag = dockerRunFlags[ref]
			}

			composePath := strings.Split(dockerFlag.ComposeName, ".")

			parent := p.document
			for len(composePath) > 0 {
				key := composePath[0]
				// find
				cNode := p.refs[key]
				if cNode == nil {
					p.addNode(parent, key, value, len(composePath))
				}
				parent = cNode
				composePath = composePath[1:]
			}
		}

		// TODO: what do we do with an unknown flag?
	}

	if errors.Is(parseErr, errNoMoreFlags) {
		parseErr = p.parseImage()
	}

	return parseErr
}

func (p *Parser) addNode(parent *yaml.Node, key, value string, depth int) {
	KeyNode := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: key,
	}

	valueNode := &yaml.Node{}
	switch depth {
	case 1: // scalar node
		valueNode.Kind = yaml.ScalarNode
		valueNode.Value = value
	case 2: // sequence node
		valueNode.Kind = yaml.SequenceNode
	default: // mapping node
		valueNode.Kind = yaml.MappingNode
	}
	p.refs[key] = valueNode
	parent.Content = append(parent.Content, KeyNode, valueNode)
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
	p.command = p.command[1:] // the rest are commands
	// TODO: what if the image name is profclems/glab:latest or just profclems/glab??
	//  we need to make sure the service name will be the last word after slash but without the
	//  tag version, like just "glab" in the example above
	p.refs["$serviceTitleNode"].Value = image
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

var (
	errNoMoreFlags = errors.New("no more flags available")
	errInvalidFlag = errors.New("invalid docker run flag")
)

func (p *Parser) parseOneFlag() (string, string, error) {
	if len(p.command) == 0 {
		return "", "", nil
	}

	f := p.command[0]
	if len(f) > 0 && f[0] != '-' {
		// this could possibly be the image name and not a flag
		return "", "", errNoMoreFlags
	}

	p.command = p.command[1:]

	if len(f) < 2 {
		return "", "", errInvalidFlag
	}
	numOfMinuses := 1
	// TODO: if the flag is a shorthand and not a boolean the value can be directly
	//  attached to the flag name like -p80:80 and -uroot
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
	for i := 0; i < len(name); i++ {
		if name[i] == '=' {
			hasValue = true
			value, name = name[i+1:], name[:i]
		}
	}

	// check if flag is a boolean flag
	if ft := flagType(name); ft != nil && *ft == BoolType {
		if hasValue {
			pv, err := strconv.ParseBool(name)
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
			} else if pv, err := strconv.ParseBool(arg); err != nil {
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
			// TODO: is it really? what about `docker run --label -myLabelWhichHasAHyphenIntheBeginning image`
			if len(p.command) > 0 {
				value, p.command = p.command[0], p.command[1:]
				hasValue = true
			}
		}

		if !hasValue {
			return "", "", fmt.Errorf("docker run flag %q is missing an argument", name)
		}
	}

	return name, value, nil
}

func flagType(name string) *FlagType {
	if dockerFlag, ok := dockerRunFlags[name]; ok {
		ftype := dockerFlag.Type
		if ref := dockerFlag.Reference; ref == "" {
			ftype = dockerRunFlags[ref].Type
		}

		return &ftype
	}

	return nil
}
