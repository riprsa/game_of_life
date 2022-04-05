package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ScreenWidth  int `yaml:"screen_width"`
	ScreenHeight int `yaml:"screen_height"`

	ToKeepAlive   []int `yaml:"to_keep_alive"`
	ToBecomeAlive []int `yaml:"to_become_alive"`

	TPS int `yaml:"TPS"`
}

func LoadConfig(path string) (config Config, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(b, &config)
	return
}
