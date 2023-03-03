package mesastate

import (
	"os"
	"strconv"
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

  // MESA history logs
  StarHistoryName   string `json:"starhistoryname"`
  BinaryHistoryName string `json:"binaryhistoryname"`
  
  // name of columns in history logs
  StarHistoryColumnNames   []string
  BinaryHistoryColumnNames []string

  // values of columns in history logs
  StarHistoryColumnValues   []float64
  BinaryHistoryColumnValues []float64
}

// SetMESAEvolutionDefaults defines defaults values of MESAEvolution struct
func SetMESAEvolutionDefaults () MESAEvolution {

  M := MESAEvolution{}

  M.OMPNumThreads = 1

  M.StarHistoryName = "LOGS/history.data"
  M.BinaryHistoryName = "binary_history.data"

  return M
}

// SetMESAEnvironVariables retrieves environment variables using OS module
func (M *MESAEvolution) SetMESAEnvironVariables () {

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
