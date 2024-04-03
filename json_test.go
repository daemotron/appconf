package appconf

import (
	"encoding/json"
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

func Test_traverseJsonFile(t *testing.T) {
	reader := strings.NewReader(testData)
	var data interface{}
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	result, err := traverseJsonFile(data, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != testLength {
		t.Fatalf("number of members returned: %d (expected: %d)", len(result), testLength)
	}
}

func Test_parseJsonFile(t *testing.T) {
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
	result, err := parseJsonFile(file.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != testLength {
		t.Fatalf("number of members returned: %d (expected: %d)", len(result), testLength)
	}
}

func Test_updateFromJsonFile(t *testing.T) {
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
	conf := NewConf("Gizmo")
	err = conf.NewOption("server.port", WithDefaultValue(IntValue(3000)), WithJson("server.port"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = conf.updateFromJsonFile(file.Name())
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
