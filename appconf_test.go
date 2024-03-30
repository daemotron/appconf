package appconf_test

import (
	"appconf"
	"testing"
)

func TestNewConf(t *testing.T) {
	conf := appconf.NewConf("Gizmo")
	if conf.Name != "Gizmo" {
		t.Fatalf("Configuration app name: %v (expected 'Gizmo')", conf.Name)
	}
}

func TestNewConfWithAuthor(t *testing.T) {
	conf := appconf.NewConf("Gizmo", appconf.WithAuthor("Ken"))
	if conf.Author != "Ken" {
		t.Fatalf("Configuration app author: %v (expected: 'Ken')", conf.Author)
	}
}

func TestNewConfWithConfFile(t *testing.T) {
	conf := appconf.NewConf("Gizmo", appconf.WithConfFile("Foo"))
	if len(conf.ConfFiles) != 1 {
		t.Fatalf("Configuration file list length: %v (expected: 1)", len(conf.ConfFiles))
	}
	if conf.ConfFiles[0] != "Foo" {
		t.Fatalf("Configuration file member: %v", conf.ConfFiles[0])
	}
}

func TestNewConfWithConfFiles(t *testing.T) {
	conf := appconf.NewConf("Gizmo", appconf.WithConfFiles([]string{"Foo", "Bar", "Baz"}))
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

func TestNewConfWithRoaming(t *testing.T) {
	conf := appconf.NewConf("Gizmo", appconf.WithRoaming())
	if !conf.Roaming {
		t.Fatalf("Configuration roaming flag: %v (true expected)", conf.Roaming)
	}
}

func TestNewConfWithVersion(t *testing.T) {
	conf := appconf.NewConf("Gizmo", appconf.WithVersion("1.0"))
	if conf.Version != "1.0" {
		t.Fatalf("Configuration app version: %v (expected: '1.0')", conf.Version)
	}
}
