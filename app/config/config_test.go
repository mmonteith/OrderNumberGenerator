package config

import (
	"errors"
	"os"
	"testing"

	"github.com/urbn/ordernumbergenerator/app"
)

func TestLoadConfig_WithOverride(t *testing.T) {
	expected := "US-PA"
	os.Setenv("DATACENTER_ID", expected)
	config, err := LoadConfig()

	if err != nil {
		t.Errorf("Did not expect an error but got one %s.", err.Error())
	}

	if expected != config.DatacenterId {
		t.Errorf("Expected to get %s in datacenter id but got: %s", expected, config.DatacenterId)
	}
}

func TestLoadConfig_Error(t *testing.T) {
	GetConfigProc = func(prefix string, Config *app.Specification) error {
		return errors.New("failed to load config, environment variable missing")
	}

	_, err := LoadConfig()
	if err == nil {
		t.Error("Expect an error but none.")
	}
}
