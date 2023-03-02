package mesastate

import (
  "bytes"
  "bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MESAproc contains information on the process running the MESA simulation such as the process id,
// the directory where the executable is located
type MESAproc struct {
  ProcID        int    `json:"procid"`
  RootDirectory string `json:"rootdirectory"`
  ProcName      string `json:"procname"`
  State         string `json:"state"`
}

// WalkProc function searches for a given `executable_name` in the /proc directory
func (M *MESAproc) WalkProc () {

  fmt.Println("searching for executable:", M.ProcName)
  
  // use filepath Walk inside /proc to search for executable_name
  err := filepath.Walk("/proc", M.findMESAproc)

  if err != nil {
    fmt.Println("error found:", err)
  }

}

// findMESAproc: process with MESA simulation
// Idea from this post:
// https://stackoverflow.com/questions/41060457/golang-kill-process-by-name
func (M *MESAproc) findMESAproc (path string, info os.FileInfo, err error) error {

  // get only process of the type `/proc/<pid>/status` where <pid> is just a number
  if strings.Count(path, "/") == 3 {

    // we want the `status` file to search for the executable name
    if strings.Contains(path, "/status") {
      pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
      if err != nil {
        return err
      }

      // fast way to check if file corresponds to executable
      f, err := os.ReadFile(path)
      if err != nil {
        return err
      }
      procName := string(f[6:bytes.IndexByte(f, '\n')])
      if procName == M.ProcName {

        // set some values
        M.ProcID = pid
        M.ProcName = procName
      
        // open file, this time to get status
        f, err := os.Open(path)
        if err != nil {
          return err
        }

        // scan whole file
        scanner := bufio.NewScanner(f)
  
        // grab state of process from `status` file
        for scanner.Scan() {
          if strings.Contains(scanner.Text(), "State") {
            state_line := scanner.Text()
            M.State = state_line[7:strings.LastIndex(state_line, "")]
          }
        }

        // close file
        f.Close()

      }
    }
  }
  return nil
}
