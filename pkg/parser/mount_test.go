package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMount(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Mount
		wantErr bool
	}{
		{
			name:    "empty mount",
			args:    "",
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid mount",
			args: "type=bind,source=/tmp,destination=/tmp,bind-propagation=rprivate",
			want: &Mount{
				Type:            "bind",
				Source:          "/tmp",
				Target:          "/tmp",
				BindPropagation: "rprivate",
			},
		},
		{
			name: "valid mount with readonly",
			args: "type=bind,source=/tmp,destination=/tmp,readonly",
			want: &Mount{
				Type:     "bind",
				Source:   "/tmp",
				Target:   "/tmp",
				Readonly: "true",
			},
		},
		{
			name: "valid mount with tmpfs",
			args: "type=tmpfs,destination=/tmp,tmpfs-size=1000000,tmpfs-mode=1777",
			want: &Mount{
				Type:      "tmpfs",
				Target:    "/tmp",
				TmpfsSize: "1000000",
				TmpfsMode: "1777",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mount, err := ParseMount(tt.args)
			if tt.wantErr {
				require.NotNil(t, err)
				require.Nil(t, mount)
				return
			}
			require.Nil(t, err)
			require.EqualValues(t, tt.want, mount)
		})
	}
}
