package mesastate

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
  ErrorPIDNotFound    = errors.New("process error: could not extract PID")
  ErrorStatusNotFound = errors.New("process error: could not read `status` file")
  ErrorStateNotFound  = errors.New("process error: state not found")
)

// GetProcessInfo looks for information on the process corresponding to an executable
// This will only work if the ProcessName field has already been set
func (M *MESAEvolution) GetProcessInfo () {

  // first, get process id (PID)
  M.walkProc()

  // second, get process state
  M.getMESAProcessState()

  // last, get directory where executable is located
  M.getMESAProcessExeDirectory()

  // also, get current working directory
  M.getMESAProcessCWD()

}

// WalkProc function searches for a given `executable_name` in the /proc directory
func (M *MESAEvolution) walkProc () {

  // use filepath Walk inside /proc to search for executable_name in each process directory
  filepath.Walk("/proc", M.findMESAProcessID)
}

// findMESAproc search ID of process associated to `executable_name`
// Idea from this post:
// https://stackoverflow.com/questions/41060457/golang-kill-process-by-name
func (M *MESAEvolution) findMESAProcessID (path string, info os.FileInfo, err error) error {

  // get only process of the type `/proc/<pid>/status` where <pid> is just a number
  if strings.Count(path, "/") == 3 {

    // we want the `status` file to search for the executable name
    if strings.Contains(path, "/status") {
    
      // store PID in case it is the right case
      pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
      if err != nil {
        M.ProcessID = -1
        return ErrorPIDNotFound
      }

      // read status file. first line contains executable name
      f, err := os.ReadFile(path)
      if err != nil {
        M.ProcessID = -1
        return ErrorStatusNotFound
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
func (M *MESAEvolution) getMESAProcessState () error {

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

  return ErrorStateNotFound
}

// GetMESAProcessExeDirectory locates the exec directory where `executable_name` is running
func (M *MESAEvolution) getMESAProcessExeDirectory () error {
  
  // full path to status file
  statusFilename := "/proc/" + strconv.Itoa(M.ProcessID) + "/exe"
  exe, err := os.Readlink(statusFilename)
  if err != nil {
    return err
  }

  // trim path to get directory only
  M.ExeDirectory = exe[0:strings.LastIndex(exe, "/")]

  // return no error
  return nil
}

// GetMESAProcessCWD locates the current working directory
func (M *MESAEvolution) getMESAProcessCWD () error {
  
  // full path to status file
  statusFilename := "/proc/" + strconv.Itoa(M.ProcessID) + "/cwd"
  cwd, err := os.Readlink(statusFilename)
  if err != nil {
    return err
  }

  M.CWD = cwd

  // return no error
  return nil
}
