package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Some configs for program. TODO: rewrite.
type Config struct {
	// Resolution of a screen data.
	ScreenWidth  int `yaml:"screen_width"`
	ScreenHeight int `yaml:"screen_height"`

	// Logic things...
	ToKeepAlive   []int `yaml:"to_keep_alive"`
	ToBecomeAlive []int `yaml:"to_become_alive"`

	// Size of map, diameter
	MapSize int `yaml:"map_size"`

	// Tick per second
	TPS int `yaml:"TPS"`
}

// Loads config
func LoadConfig(path string) (config Config, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(b, &config)
	return
}
