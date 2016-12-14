package wire

import "path/filepath"

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
