package mesastate

import (
  "bytes"
  "bufio"
  "os"
  "path/filepath"
  "strconv"
  "strings"
  "errors"
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
func (M *MESAprocess) WalkProc () {

  // use filepath Walk inside /proc to search for executable_name in each process directory
  filepath.Walk("/proc", M.findMESAProcessID)
}

// findMESAproc: process with MESA simulation
// Idea from this post:
// https://stackoverflow.com/questions/41060457/golang-kill-process-by-name
func (M *MESAprocess) findMESAProcessID (path string, info os.FileInfo, err error) error {

  // get only process of the type `/proc/<pid>/status` where <pid> is just a number
  if strings.Count(path, "/") == 3 {

    // we want the `status` file to search for the executable name
    if strings.Contains(path, "/status") {
    
      // store PID in case it is the right case
      pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
      if err != nil {
        M.ProcessID = -1
        return errors.New("process error: could not extract PID")
      }

      // read status file. first line contains executable name
      f, err := os.ReadFile(path)
      if err != nil {
        M.ProcessID = -1
        return errors.New("process error: could not read `status` file")
      }
      processName := string(f[6:bytes.IndexByte(f, '\n')])

      // matching process name ? return data
      if processName == M.ProcessName {
        // store PID of founded matching process
        M.ProcessID = pid
        return filepath.SkipAll
      }
    }
  }

  return nil
}

// GetMESAProcessState retrieves the state of the process from the `status` file. It must be called
// after getting the process ID (PID) as it needs the path `/proc/PID/status`
func (M *MESAprocess) GetMESAProcessState () error {

  // full path to status file
  statusFilename := "/proc/" + strconv.Itoa(M.ProcessID) + "/status"
  
  // open file, this time to get status
  f, err := os.Open(statusFilename)
  if err != nil {
    return err
  }

  defer f.Close()

  // scan whole file
  scanner := bufio.NewScanner(f)

  // grab state of process from `status` file
  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "State") {
      stateLine := scanner.Text()
      M.State = stateLine[7:strings.LastIndex(stateLine, "")]
      return nil
    }
  }

  return errors.New("process error: state not found")
}

func (M *MESAprocess) GetMESAProcessRootDirectory () error {
  
  // full path to status file
  statusFilename := "/proc/" + strconv.Itoa(M.ProcessID) + "/exe"
  exe, err := os.Readlink(statusFilename)
  if err != nil {
    return err
  }

  // trim path to get directory only
  M.RootDirectory = exe[0:strings.LastIndex(exe, "/")]

  // return no error
  return nil
}


