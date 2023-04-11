package mesastate

import (
	"os"
	"strconv"

  "github.com/asimazbunzel/go-mesa-state/log"
)

// MESAproc contains information on the process running the MESA simulation such as the process id,
// the directory where the executable is located
type MESAEvolution struct {

  // variables related to actual process
  ProcessID    int    `json:"procid"`
  ExeDirectory string `json:"exedirectory"`
  CWD          string `json:"cwd"`
  ProcessName  string `json:"procname"`
  State        string `json:"state"`

  // environment variables used by the MESA code
  MESADir       string `json:"mesadir"`
  MESASdkRoot   string `json:"mesasdkroot"`
  MESAInlist    string `json:"mesainlist"`
  MESACachesDir string `json:"mesacachesdir"`
  OMPNumThreads int    `json:"ompnumthreads"`

  // flag to know if Evolution is isolated star or binary
  IsBinary bool `json:"isbinary"`

  // structs with MESA output
  Star1 MESAstar
  Star2 MESAstar
  Binary MESAbinary
}

// MESAstar holds information on a MESAstar evolution
type MESAstar struct {
  
  // MESA history logs
  HistoryName   string `json:"starhistoryname"`

  // name of columns in history logs
  HistoryColumnNames   []string

  // values of columns in history logs
  HistoryColumnValues   []float64
}

// MESAbinary holds information on a MESAbinary evolution
type MESAbinary struct {
  
  // MESA history logs
  HistoryName string `json:"binaryhistoryname"`

  // name of columns in history logs
  HistoryColumnNames []string

  // values of columns in history logs
  HistoryColumnValues []float64
}

// SetMESAEvolutionDefaults defines defaults values of MESAEvolution struct
func SetMESAEvolutionDefaults () MESAEvolution {

  log.LogInfo("mesa_state/mesa_state.go (func SetMESAEvolutionDefaults)", "setting MESAEvolution defaults")

  M := MESAEvolution{}

  M.OMPNumThreads = 1

  M.Star1.HistoryName = "history.data"
  M.Star2.HistoryName = "history.data"
  M.Binary.HistoryName = "binary_history.data"

  return M
}

// SetMESAEnvironVariables retrieves environment variables using OS module
func (M *MESAEvolution) SetMESAEnvironVariables () {
  
  log.LogInfo("mesa_state/mesa_state.go (func SetMESAEnvironVariables)", "setting MESA environment variables")

  // use os to get all these vars. If not found, they are empty strings
  M.MESADir = os.Getenv("MESA_DIR")
  M.MESASdkRoot = os.Getenv("MESASDK_ROOT")
  M.MESACachesDir = os.Getenv("MESA_CACHES_DIR")
  M.MESAInlist = os.Getenv("MESA_INLIST")

  numThreads, err := strconv.Atoi(os.Getenv("OMP_NUM_THREADS"))
  if err != nil {
    return
  }
  M.OMPNumThreads = numThreads
}
