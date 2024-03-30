package appconf_test

import (
	"appconf"
	"testing"
)

func TestNew(t *testing.T) {
	_ = appconf.New()
}

func TestNewWithConfFile(t *testing.T) {
	conf := appconf.New(appconf.WithConfFile("Foo"))
	if len(conf.ConfFiles) != 1 {
		t.Fatalf("Configuration file list length: %v (expected: 1)", len(conf.ConfFiles))
	}
	if conf.ConfFiles[0] != "Foo" {
		t.Fatalf("Configuration file member: %v", conf.ConfFiles[0])
	}
}

func TestNewWithConfFiles(t *testing.T) {
	conf := appconf.New(appconf.WithConfFiles([]string{"Foo", "Bar", "Baz"}))
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
