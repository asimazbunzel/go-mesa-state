package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	
  "github.com/asimazbunzel/go-mesa-state/log"
)

type Config struct {
  level string `yaml:"level"`
}

func (c *Config) ParseYAML(filename string) error {

  log.LogDebug("config/config.go", "parsing options from YAML file")
  
  // read YAML data file into bytes
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return err
  }

  return yaml.Unmarshal(data, c)
}
