package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	file, _ := os.CreateTemp(t.TempDir(), "*.yaml")
	b := "test_cfg:\n  test_field: \"testvalue\""
	file.WriteString(b)

	testStruct := &struct {
		TestCfg struct {
			TestField string `yaml:"test_field"`
		} `yaml:"test_cfg"`
	}{}

	err := ConfigLoader(testStruct, file.Name())
	if err != nil {
		t.Errorf("%s : when load config", err)
	}

	if testStruct.TestCfg.TestField != "testvalue" {
		t.Errorf("Loaded field not equal with expected")
	}
}
