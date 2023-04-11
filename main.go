package main

import (
	"github.com/asimazbunzel/go-mesa-state/cmd"
  "github.com/asimazbunzel/go-mesa-state/config"
	"github.com/asimazbunzel/go-mesa-state/log"
  "github.com/asimazbunzel/go-mesa-state/mesa_state"
)

func main() {

  log.LogInfo("main.go (func main)", "start main function")

  var cfg config.Config
  cmd.Start(&cfg)

  mEvol := mesastate.SetMESAEvolutionDefaults()
  mEvol.SetMESAEnvironVariables()
  mEvol.ProcessName = "binary"
  mEvol.GetProcessInfo()
}
