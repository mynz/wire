package wire

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

const configPath = "_config.json"

type Config struct {
	RootDir string
}

func NewConfig() *Config {
	conf := &Config{"."}
	conf.Setup()
	return conf
}

func (conf *Config) Setup() {
	if ap, err := filepath.Abs(conf.RootDir); err == nil {
		conf.RootDir = ap
	}
}

func (conf *Config) SaveFile() error {
	bs, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(configPath, bs, 0666); err != nil {
		return err
	}
	return nil
}
