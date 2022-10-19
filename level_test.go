package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name string
		want string
		l    Level
	}{
		{
			name: "debug",
			l:    DebugLevel,
			want: "DEBUG",
		},
		{
			name: "info",
			l:    InfoLevel,
			want: "INFO",
		},
		{
			name: "warn",
			l:    WarnLevel,
			want: "WARN",
		},
		{
			name: "error",
			l:    ErrorLevel,
			want: "ERROR",
		},
		{
			name: "off",
			l:    OffLevel,
			want: "OFF",
		},
		{
			name: "higher",
			l:    Level(99),
			want: "OFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.String())
		})
	}
}

func TestLevel_Enabled(t *testing.T) {
	type args struct {
		proba Level
	}
	tests := []struct {
		name string
		l    Level
		args args
		want bool
	}{
		{
			name: "warn_enabled_warn",
			l:    WarnLevel,
			args: args{proba: WarnLevel},
			want: true,
		},
		{
			name: "warn_enabled_error",
			l:    WarnLevel,
			args: args{proba: ErrorLevel},
			want: true,
		},
		{
			name: "warn_not_enable_off",
			l:    WarnLevel,
			args: args{proba: OffLevel},
			want: false,
		},
		{
			name: "warn_not_enable_info",
			l:    WarnLevel,
			args: args{proba: InfoLevel},
			want: false,
		},
		{
			name: "warn_not_enable_debug",
			l:    WarnLevel,
			args: args{proba: DebugLevel},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.l.Enabled(tt.args.proba), "Enabled(%v)", tt.args.proba)
		})
	}
}

func TestLevel_MarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		l       Level
		wantErr bool
	}{
		{
			name: "debug",
			l:    DebugLevel,
			want: "DEBUG",
		},
		{
			name: "info",
			l:    InfoLevel,
			want: "INFO",
		},
		{
			name: "warn",
			l:    WarnLevel,
			want: "WARN",
		},
		{
			name: "error",
			l:    ErrorLevel,
			want: "ERROR",
		},
		{
			name: "off",
			l:    OffLevel,
			want: "OFF",
		},
		{
			name: "higher",
			l:    Level(99),
			want: "OFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := tt.l.MarshalYAML()
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want, v)
			}
		})
	}
}

func TestLevel_UnmarshalYAML(t *testing.T) {
	initLevel := InfoLevel
	tests := []struct {
		in      string
		want    Level
		wantErr bool
	}{
		{
			in:      "",
			wantErr: false,
		},
		{
			in:      "0",
			wantErr: true,
		},
		{
			in:   "INFO",
			want: InfoLevel,
		},
		{
			in:   "debug",
			want: DebugLevel,
		},
		{
			in:   "Warn",
			want: WarnLevel,
		},
		{
			in:   "'ERROR'",
			want: ErrorLevel,
		},
		{
			in:   `"OFF"`,
			want: OffLevel,
		},
		{
			in:      "vv",
			wantErr: true,
		},
		{
			in:      "abc",
			wantErr: true,
		},
		{
			in:      `''`,
			wantErr: true,
		},
		{
			in:      `""`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			level := initLevel
			err := yaml.Unmarshal([]byte(tt.in), &level)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, tt.want, level)
		})
	}
}
