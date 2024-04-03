package appconf

import (
	"os"
	"testing"
)

func TestAppConf_UpdateFromFlags(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithFlag("foo"), WithDefaultString("bar"))
	if err != nil {
		t.Errorf("unexpected error while registering option: %v", err)
	}
	err = conf.NewOption("qux", WithFlag("qux"), WithDefaultInt(123))
	if err != nil {
		t.Errorf("unexpected error while registering option: %v", err)
	}
	err = conf.NewOption("corge", WithFlag("corge"), WithDefaultFloat(123.456))
	if err != nil {
		t.Errorf("unexpected error while registering option: %v", err)
	}
	err = conf.NewOption("waldo", WithFlag("waldo"), WithDefaultBool(false))
	if err != nil {
		t.Errorf("unexpected error while registering option: %v", err)
	}

	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-foo", "baz", "-qux", "456", "-corge", "234.567", "-waldo"}

	err = conf.UpdateFromFlags()
	if err != nil {
		t.Errorf("error while updating flags: %v", err)
	}

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"flag test with string", "foo", "baz"},
		{"flag test with int", "qux", "456"},
		{"flag test with float", "corge", "234.567"},
		{"flag test with bool", "waldo", "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if conf.Options[tt.key].Value.ToString() != tt.want {
				t.Errorf("conf.Options['%s'].Value = %s, expected: %s", tt.key, conf.Options["foo"].Value.ToString(), tt.want)
			}
		})
	}
}
