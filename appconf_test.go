package appconf

import (
	"testing"
)

func TestAppConf_NewConf(t *testing.T) {
	conf := NewConf("Gizmo")
	if conf.Name != "Gizmo" {
		t.Fatalf("Configuration app name: %v (expected 'Gizmo')", conf.Name)
	}
}

func TestAppConf_NewConf_WithAuthor(t *testing.T) {
	conf := NewConf("Gizmo", WithAuthor("Ken"))
	if conf.Author != "Ken" {
		t.Fatalf("Configuration app author: %v (expected: 'Ken')", conf.Author)
	}
}

func TestAppConf_NewConf_WithConfFile(t *testing.T) {
	conf := NewConf("Gizmo", WithConfFile("Foo"))
	if len(conf.ConfFiles) != 1 {
		t.Fatalf("Configuration file list length: %v (expected: 1)", len(conf.ConfFiles))
	}
	if conf.ConfFiles[0] != "Foo" {
		t.Fatalf("Configuration file member: %v", conf.ConfFiles[0])
	}
}

func TestAppConf_NewConf_WithConfFiles(t *testing.T) {
	conf := NewConf("Gizmo", WithConfFiles([]string{"Foo", "Bar", "Baz"}))
	if len(conf.ConfFiles) != 3 {
		t.Fatalf("Configuration file list length: %v (expected: 3)", len(conf.ConfFiles))
	}
	if conf.ConfFiles[0] != "Foo" {
		t.Fatalf("Configuration files member 1: %v (expected: 'Foo')", conf.ConfFiles[0])
	}
	if conf.ConfFiles[1] != "Bar" {
		t.Fatalf("Configuration files member 2: %v (expected: 'Bar')", conf.ConfFiles[1])
	}
	if conf.ConfFiles[2] != "Baz" {
		t.Fatalf("Configuration files member 3: %v (expected: 'Baz')", conf.ConfFiles[2])
	}
}

func TestAppConf_NewConf_WithRoaming(t *testing.T) {
	conf := NewConf("Gizmo", WithRoaming())
	if !conf.Roaming {
		t.Fatalf("Configuration roaming flag: %v (true expected)", conf.Roaming)
	}
}

func TestAppConf_NewConf_WithVersion(t *testing.T) {
	conf := NewConf("Gizmo", WithVersion("1.0"))
	if conf.Version != "1.0" {
		t.Fatalf("Configuration app version: %v (expected: '1.0')", conf.Version)
	}
}

func TestAppConf_NewOption(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	_, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
}

func TestAppConf_NewOption_WithDefaultValue(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithDefaultString("bar"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	val, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
	if val.Default.ToString() != "bar" {
		t.Fatalf("Default value not correct: %s (expected: 'bar')", val.Default.ToString())
	}
	if val.Value.ToString() != "bar" {
		t.Fatalf("Current value not correct: %s (expected: 'bar')", val.Value.ToString())
	}
}

func TestAppConf_NewOption_WithFlag(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithFlag("f"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	val, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
	if val.Flag != "f" {
		t.Fatalf("Flag not correct: %s (expected: 'f')", val.Flag)
	}
}

func TestAppConf_NewOption_WithJson(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithJson("app.foo"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	val, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
	if val.Json != "app.foo" {
		t.Fatalf("Flag not correct: %s (expected: 'app.foo')", val.Flag)
	}
}

func TestAppConf_NewOption_WithEnv(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithEnv("FOO"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	val, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
	if val.Env != "FOO" {
		t.Fatalf("Flag not correct: %s (expected: 'FOO')", val.Flag)
	}
}

func TestAppConf_NewOption_WithHelp(t *testing.T) {
	conf := NewConf("Gizmo")
	err := conf.NewOption("foo", WithHelp("foo help text"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	val, ok := conf.Options["foo"]
	if !ok {
		t.Fatalf("Key of registered option not found")
	}
	if val.Help != "foo help text" {
		t.Fatalf("Flag not correct: %s (expected: 'foo help text')", val.Flag)
	}
}
