package cmd

import (
	"flag"
	"os"

	"github.com/asimazbunzel/go-mesa-state/config"
  "github.com/asimazbunzel/go-mesa-state/log"
)


func Start(cfg *config.Config) {
  
  log.LogDebug("cmd/root.go", "set argument parser")

  var cfgFile string
  flag.StringVar(&cfgFile, "config-file", "config.yaml", "Name of configuration file")
  flag.Parse()

  err := cfg.ParseYAML(cfgFile)
  if err != nil {
    os.Exit(1)
  }

}
