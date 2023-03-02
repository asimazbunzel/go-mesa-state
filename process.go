package mesastate

import (
  "bytes"
  // "bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// MESAproc contains information on the process running the MESA simulation such as the process id,
// the directory where the executable is located
type MESAprocess struct {
  ProcessID     int    `json:"procid"`
  RootDirectory string `json:"rootdirectory"`
  ProcessName   string `json:"procname"`
  State         string `json:"state"`
}

// WalkProc function searches for a given `executable_name` in the /proc directory
func (M *MESAprocess) WalkProc () error {

  fmt.Println("searching for process ID of executable:", M.ProcessName)
  
  // use filepath Walk inside /proc to search for executable_name in each process directory
  err := filepath.Walk("/proc", M.findMESAProcessID)
  if err != nil {
    fmt.Println("error found:", err)
    return err
  }

  fmt.Println("here we are:", M)

  return nil
}

// findMESAproc: process with MESA simulation
// Idea from this post:
// https://stackoverflow.com/questions/41060457/golang-kill-process-by-name
func (M *MESAprocess) findMESAProcessID (path string, info os.FileInfo, err error) error {

  // get only process of the type `/proc/<pid>/status` where <pid> is just a number
  if strings.Count(path, "/") == 3 {

    fmt.Println("looking for process in", path)

    // we want the `status` file to search for the executable name
    if strings.Contains(path, "/status") {

      // store PID in case it is the right case
      pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
      if err != nil {
        return err
      }

      // read status file. first line contains executable name
      f, err := os.ReadFile(path)
      processName := string(f[6:bytes.IndexByte(f, '\n')])
      if processName == M.ProcessName {
        
        fmt.Println("found process with name:", processName)

        // store PID of founded matching process
        M.ProcessID = pid
      }
    }
  }

  return nil
}
      

// fast way to check if file corresponds to executable
  //     f, err := os.ReadFile(path)
  //     if err != nil {
  //       return err
  //     }
  //     procName := string(f[6:bytes.IndexByte(f, '\n')])
  //     if procName == M.ProcName {
  //
  //       // set some values
  //       M.ProcID = pid
  //       M.ProcName = procName
  //     
  //       // open file, this time to get status
  //       f, err := os.Open(path)
  //       if err != nil {
  //         return err
  //       }
  //
  //       // scan whole file
  //       scanner := bufio.NewScanner(f)
  // 
  //       // grab state of process from `status` file
  //       for scanner.Scan() {
  //         if strings.Contains(scanner.Text(), "State") {
  //           state_line := scanner.Text()
  //           M.State = state_line[7:strings.LastIndex(state_line, "")]
  //         }
  //       }
  //
  //       // close file
  //       f.Close()
