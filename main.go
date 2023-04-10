package main

import (
	"github.com/asimazbunzel/go-mesa-state/cmd"
  "github.com/asimazbunzel/go-mesa-state/config"
	"github.com/asimazbunzel/go-mesa-state/log"
)

func main() {

  log.LogDebug("main.go", "start main function")

  var cfg config.Config
  cmd.Start(&cfg)
	
}
