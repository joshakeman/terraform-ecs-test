package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var ErrInvalidConfiguration = errors.New("invalid configuration")

type Settings struct {
	Environment string ` mapstructure:"ENVIRONMENT"`
	list        map[string]string
}

var settings = &Settings{
	list: map[string]string{
		"Environment": "",
	},
}

func Load() error {
	settingsFile, err := ioutil.ReadFile("settings.json")

	if err == nil {
		if err := json.Unmarshal(settingsFile, settings); err != nil {
			return errors.Wrap(err, "error reading configuration file.")
		}
	}

	if os.Getenv("ENVIRONMENT") != "" {
		settings.Environment = os.Getenv("ENVIRONMENT")
	}

	settings.list["ENVIRONMENT"] = settings.Environment

	if settings.Environment == "" {
		return ErrInvalidConfiguration
	}

	return nil
}

func Get(key string) string {
	return settings.list[key]
}
