package appconf

import (
	"os"
	"testing"
)

func TestAppConf_isFile(t *testing.T) {
	file, err := os.CreateTemp("", "test-*.file")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	dir, err := os.MkdirTemp("", "appconf-test-*")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())
	defer func(name string) {
		_ = os.RemoveAll(name)
	}(dir)
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Check with file", args: args{path: file.Name()}, want: true},
		{name: "Check with dir", args: args{path: dir}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFile(tt.args.path); got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppConf_ConfigFiles(t *testing.T) {
	conf := NewConf("Gizmo")
	_, err := conf.ConfigFiles()
	if err != nil {
		t.Errorf("error while retrieving AppConf.ConfigFiles(): %v", err)
	}
}

func TestAppConf_UpdateFromFiles(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.UpdateFromFiles()
	if err != nil {
		t.Errorf("error while retrieving AppConf.ConfigFiles(): %v", err)
	}
}
