package appconf

import (
	"os"
	"testing"
)

func TestAppConf_UpdateFromEnv_String(t *testing.T) {
	type args struct {
		Key     string
		Env     string
		Val     string
		Default string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "string environment test with string",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "baz", Default: "bar"},
			want: "baz",
		},
		{
			name: "string environment test with empty string",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "", Default: "bar"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := NewConf("Gizmo")
			err := conf.NewOption(tt.args.Key, WithEnv(tt.args.Env), WithDefaultString(tt.args.Default))
			if err != nil {
				t.Errorf("unexpected error while registering option: %v", err)
			}
			err = os.Setenv(tt.args.Env, tt.args.Val)
			if err != nil {
				t.Errorf("unexpected error while setting environment variable: %v", err)
			}
			err = conf.UpdateFromEnv()
			if err != nil {
				t.Errorf("error while updating conf from environment variable: %v", err)
			}
			if conf.Options[tt.args.Key].Value.ToString() != tt.want {
				t.Errorf("After UpdateFromEnv() values don't match, got %v, wanted %v", conf.Options[tt.args.Key].Value.ToString(), tt.want)
			}
		})
	}
}

func TestAppConf_UpdateFromEnv_Int(t *testing.T) {
	type args struct {
		Key     string
		Env     string
		Val     string
		Default int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "int environment test with positive int",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "456", Default: 123},
			want: 456,
		},
		{
			name: "int environment test with negative int",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "-456", Default: 123},
			want: -456,
		},
		{
			name: "int environment test with zero",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "0", Default: 123},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := NewConf("Gizmo")
			err := conf.NewOption(tt.args.Key, WithEnv(tt.args.Env), WithDefaultInt(tt.args.Default))
			if err != nil {
				t.Errorf("unexpected error while registering option: %v", err)
			}
			err = os.Setenv(tt.args.Env, tt.args.Val)
			if err != nil {
				t.Errorf("unexpected error while setting environment variable: %v", err)
			}
			err = conf.UpdateFromEnv()
			if err != nil {
				t.Errorf("error while updating conf from environment variable: %v", err)
			}
			got, err := conf.Options[tt.args.Key].Value.ToInt()
			if err != nil {
				t.Errorf("error while retrieving int value from configuration: %v", err)
			}
			if got != tt.want {
				t.Errorf("After UpdateFromEnv() values don't match, got %v, wanted %v", got, tt.want)
			}
		})
	}
}
