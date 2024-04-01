package appconf_test

import (
	"encoding/json"
	"github.com/daemotron/appconf"
	"os"
	"strings"
	"testing"
)

const (
	testLength = 3
	testPort   = 8080
	testData   = `{
  "server": {
    "host": "localhost",
    "port": 8080
  },
  "testbed": true
}`
)

func TestTraverseJsonFile(t *testing.T) {
	reader := strings.NewReader(testData)
	var data interface{}
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := appconf.TraverseJsonFile(data, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != testLength {
		t.Fatalf("number of members returned: %d (expected: %d)", len(result), testLength)
	}
}

func TestParseJsonFile(t *testing.T) {
	file, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())
	_, err = file.WriteString(testData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := appconf.ParseJsonFile(file.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != testLength {
		t.Fatalf("number of members returned: %d (expected: %d)", len(result), testLength)
	}
}

func TestUpdateFromJsonFile(t *testing.T) {
	file, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	defer func(name string) {
		_ = os.Remove(name)
	}(file.Name())
	_, err = file.WriteString(testData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	conf := appconf.NewConf("Gizmo")
	err = conf.NewOption("server.port", appconf.WithDefaultValue(appconf.IntValue(3000)), appconf.WithJson("server.port"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = conf.UpdateFromJsonFile(file.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	port, err := conf.Options["server.port"].Value.ToInt()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if port != testPort {
		t.Fatalf("incorrect datum: %d (expected: %d)", port, testPort)
	}
}
