package parser

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParseUlimit(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		want       *Ulimit
		yamlString string
		wantErr    bool
	}{
		{
			name:    "empty ulimit",
			s:       "",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid ulimit",
			s:       "nofile",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid ulimit value",
			s:       "nofile=1024:2048:4096",
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid ulimit",
			s:    "nofile=1024:2048",
			want: &Ulimit{
				Name:     "nofile",
				Soft:     1024,
				Hard:     2048,
				NodeType: MapType,
			},
			yamlString: `nofile:
    soft: 1024
    hard: 2048
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUlimit(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUlimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUlimit() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErr {
				s, err := got.YAMLString()
				require.NoError(t, err)
				require.Equal(t, tt.yamlString, s)
			}
		})
	}
}
