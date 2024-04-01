package appconf

import (
	"os"
	"testing"
)

func TestAppConf_UpdateFromEnv(t *testing.T) {
	type args struct {
		Key     string
		Env     string
		Val     string
		Default Value
	}
	tests := []struct {
		name string
		args args
		want Value
	}{
		{
			name: "integer environment test with int",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "456", Default: IntValue(123)},
			want: IntValue(456),
		},
		{
			name: "float environment test with float",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "456.789", Default: FloatValue(123.456)},
			want: FloatValue(456.789),
		},
		{
			name: "float environment test with int",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "456", Default: FloatValue(123.456)},
			want: FloatValue(456.0),
		},
		{
			name: "bool environment test with bool",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "true", Default: BoolValue(false)},
			want: BoolValue(true),
		},
		{
			name: "bool environment test with int",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "1", Default: BoolValue(false)},
			want: BoolValue(true),
		},
		{
			name: "string environment test",
			args: args{Key: "foo", Env: "TEST_APPCONF_FOO", Val: "baz", Default: StringValue("bar")},
			want: StringValue("baz"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := NewConf("Gizmo")
			err := conf.NewOption(tt.args.Key, WithEnv(tt.args.Env), WithDefaultValue(tt.args.Default))
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
			if conf.Options[tt.args.Key].Value != tt.want {
				t.Errorf("After UpdateFromEnv() values don't match, got %v, wanted %v", conf.Options[tt.args.Key].Value, tt.want)
			}
		})
	}
}
