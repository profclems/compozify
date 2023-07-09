package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		want     string
		wantErr  string
		parseErr string
	}{
		{
			name:    "empty command",
			command: "",
			want:    "",
			wantErr: "empty docker command",
		},
		{
			name:    "basic command",
			command: "docker run -i -t --rm alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        image: alpine
`,
		},
		{
			name:    "command with multiple port mapping",
			command: "docker run -i -t --rm -p 8080:80 -p 8081:80 alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        ports:
            - 8080:80
            - 8081:80
        image: alpine
`,
		},
		{
			name:    "command with multiple volume mapping",
			command: "docker run -i -t --rm -v /tmp:/tmp -v /var/log:/var/log -v /usr/bin:/usr/bin alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        volumes:
            - /tmp:/tmp
            - /var/log:/var/log
            - /usr/bin:/usr/bin
        image: alpine
`,
		},
		{
			name:    "command with multiple environment variables",
			command: "docker run -i -t --rm -e ENV1=VALUE1 -e ENV2=VALUE2 -e ENV3=VALUE3 alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        environment:
            ENV1: VALUE1
            ENV2: VALUE2
            ENV3: VALUE3
        image: alpine
`,
		},
		{
			name:    "command with multiple environment variables and multiple port mapping",
			command: "docker run -i -t --rm -p 8080:80 -p 8081:80 -e ENV1=VALUE1 -e ENV2=VALUE2 -e ENV3=VALUE3 alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        ports:
            - 8080:80
            - 8081:80
        environment:
            ENV1: VALUE1
            ENV2: VALUE2
            ENV3: VALUE3
        image: alpine
`,
		},
		{
			name:    "command with multiple environment variables and multiple volume mapping",
			command: "docker run -i -t --rm -v /tmp:/tmp -v /var/log:/var/log -v /usr/bin:/usr/bin -e ENV1=VALUE1 -e ENV2=VALUE2 -e ENV3=VALUE3 alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        volumes:
            - /tmp:/tmp
            - /var/log:/var/log
            - /usr/bin:/usr/bin
        environment:
            ENV1: VALUE1
            ENV2: VALUE2
            ENV3: VALUE3
        image: alpine
`,
		},
		{
			name:    "command with multiple environment variables and multiple volume mapping and multiple port mapping",
			command: "docker run -i -t --rm -p 8080:80 -p 8081:80 -v /tmp:/tmp -v /var/log:/var/log -v /usr/bin:/usr/bin -e ENV1=VALUE1 -e ENV2=VALUE2 -e ENV3=VALUE3 alpine",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        ports:
            - 8080:80
            - 8081:80
        volumes:
            - /tmp:/tmp
            - /var/log:/var/log
            - /usr/bin:/usr/bin
        environment:
            ENV1: VALUE1
            ENV2: VALUE2
            ENV3: VALUE3
        image: alpine
`,
		},
		{
			name:    "command with multiple environment variables and multiple volume mapping and multiple port mapping and multiple commands",
			command: "docker run -i -t --rm -p 8080:80 -p 8081:80 -v /tmp:/tmp -v /var/log:/var/log -v /usr/bin:/usr/bin -e ENV1=VALUE1 -e ENV2=VALUE2 -e ENV3=VALUE3 alpine sh -c ls -l",
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        ports:
            - 8080:80
            - 8081:80
        volumes:
            - /tmp:/tmp
            - /var/log:/var/log
            - /usr/bin:/usr/bin
        environment:
            ENV1: VALUE1
            ENV2: VALUE2
            ENV3: VALUE3
        image: alpine
        command:
            - sh
            - -c
            - ls
            - -l
`,
		},
		{
			name: "command with sysctls",
			command: `docker run -i -t --rm \
--sysctl net.core.somaxconn=1024 \
--sysctl net.ipv4.tcp_syncookies=0 \
--sysctl net.ipv4.tcp_max_syn_backlog=2048 \
--sysctl net.ipv4.tcp_synack_retries=2 \
alpine`,
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        sysctls:
            net.core.somaxconn: 1024
            net.ipv4.tcp_syncookies: 0
            net.ipv4.tcp_max_syn_backlog: 2048
            net.ipv4.tcp_synack_retries: 2
        image: alpine
`,
		},
		{
			name: "command with log driver and log options",
			command: `docker run -i -t --rm \
--log-driver syslog \
--log-opt syslog-address=udp://
--log-opt syslog-facility=daemon \
--log-opt syslog-format=rfc5424micro \
--log-opt tag="{{.Name}}/{{.ID}}" \
alpine`,
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        logging:
            driver: syslog
            options:
                syslog-address: udp://
                syslog-facility: daemon
                syslog-format: rfc5424micro
                tag: '{{.Name}}/{{.ID}}'
        image: alpine
`,
		},
		{
			name: "command with ulimits",
			command: `docker run -i -t --rm \
--ulimit core=100000000:100000000 \
--ulimit memlock=-1:-1 \
--ulimit nofile=1024:1024 \
--ulimit nproc=65535:65535 \
alpine`,
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        ulimits:
            core:
                soft: 100000000
                hard: 100000000
            memlock:
                soft: -1
                hard: -1
            nofile:
                soft: 1024
                hard: 1024
            nproc:
                soft: 65535
                hard: 65535
        image: alpine
`,
		},

		{
			name: "command with mount options",
			command: `docker run -i -t --rm \
--mount type=bind,source=/tmp,target=/tmp, \
alpine`,
			want: `version: "3.8"
services:
    alpine:
        stdin_open: true
        tty: true
        volumes:
            - type: bind
              source: /tmp
              target: /tmp
        image: alpine
`,
		},
		{
			name: "command with tmpfs options",
			command: `docker run -i -t --rm \
--mount type=tmpfs,destination=/tmp,tmpfs-size=100000000,tmpfs-mode=1777 \
alpine`,
			want: `version: "3.8"
services:
	alpine:
		stdin_open: true
		tty: true
		volumes:
			- type: tmpfs
			  destination: /tmp
			  tmpfs_opts:
				  size: 100000000
				  mode: 1777
		image: alpine
		
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := NewParser(tt.command)
			if tt.wantErr != "" {
				require.NotNil(t, err)
				require.Nil(t, parser)
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			err = parser.Parse()
			if tt.parseErr != "" {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			require.Equal(t, tt.want, parser.String())
		})
	}
}
