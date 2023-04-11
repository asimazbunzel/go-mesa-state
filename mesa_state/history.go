package mesastate

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"

  "github.com/asimazbunzel/go-mesa-state/log"
)
  
var (
  NR_COLUMN_NAMES = 6
)

// SetStarHistoryName defines the name of the history file of MESAstar
func (M *MESAEvolution) SetStarHistoryName (id int, filename string) {

  log.LogInfo("mesa_state/history.go (func SetStarHistoryName)", "setting star history name")
  
  // the name will change according the id of the star. Possible values
  // are `1` for star 1 or `2` for star 2. In case of an isolated star
  // evolution, set it to `1`
  if id == 1 {
    M.Star1.HistoryName = filename
  } else {
    M.Star2.HistoryName = filename
  }

}

// SetBinaryHistoryName defines the name of the history file of MESAbinary
func (M *MESAEvolution) SetBinaryHistoryName (filename string) {
  
  log.LogInfo("mesa_state/history.go (func SetBinaryHistoryName)", "setting binary history name")

  M.Binary.HistoryName = filename
}

// SetColumnNames stores the names of the columns of the MESA output
func (M *MESAEvolution) SetColumnNames () {
  
  log.LogInfo("mesa_state/history.go (func SetColumnNames)", "setting column names")

  if M.IsBinary {

    M.Binary.HistoryColumnNames = findColumnNames(M.Binary.HistoryName)
    M.Star1.HistoryColumnNames  = findColumnNames(M.Star1.HistoryName)
    M.Star2.HistoryColumnNames  = findColumnNames(M.Star2.HistoryName)

  } else {
    
    M.Star1.HistoryColumnNames = findColumnNames(M.Star1.HistoryName)
  }
}

// SetColumnValues stores the values of the columns retrieved by the SetColumnNames method
func (M *MESAEvolution) SetColumnValues () {
  
  log.LogInfo("mesa_state/history.go (func SetColumnValues)", "setting column values")
  
  if M.IsBinary {

    M.Binary.HistoryColumnValues = findColumnValues(M.Binary.HistoryName)
    M.Star1.HistoryColumnValues  = findColumnValues(M.Star1.HistoryName)
    M.Star2.HistoryColumnValues  = findColumnValues(M.Star2.HistoryName)

  } else {
    
    M.Star1.HistoryColumnValues  = findColumnValues(M.Star1.HistoryName)

  }
}

// findColumnNames retrieves the row with the column names of MESA output
func findColumnNames (filename string) []string {

  log.LogDebug("mesa_state/history.go (func findColumnNames)", "finding column names")

  // open file
  f, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer f.Close()

  // scan file
  scanner := bufio.NewScanner(f)
  
  // search each line until reaching row number with column names
  columnNamesFound := false
  var columnNames []string
  lineCount := 0
  for scanner.Scan(){

    lineCount++
    
    // get the correct row number for the column names
    if lineCount == NR_COLUMN_NAMES {
      columnNames = strings.Fields(scanner.Text())
      columnNamesFound = true
    }

    if columnNamesFound {
      break
    }
  }

  return columnNames
}

// findColumnValues get last line from a filename
// Idea from this post:
// https://stackoverflow.com/questions/17863821/how-to-read-last-lines-from-a-big-file-with-go-every-10-secs
func findColumnValues (filename string) []float64 {
  
  log.LogDebug("mesa_state/history.go (func findColumnValues)", "finding column values")

  // open file
  f, err := os.Open(filename)
  if err != nil {
    return nil
  }
  defer f.Close()

  line := ""
  var columnValuesStr []string
  var columnValuesFlt []float64
  var cursor int64 = 0
  stat, _ := f.Stat()
  filesize := stat.Size()
  for {
    cursor -= 1
    f.Seek(cursor, io.SeekEnd)

    char := make([]byte, 1)
    f.Read(char)

    // stop if we find a line
    if cursor != -1 && (char[0] == 10 || char[0] == 13) {
      break
    }

    line = fmt.Sprintf("%s%s", string(char), line)

    // stop if we are at the begining
    if cursor == -filesize {
      break
    }
  }
  
  // parse string to array of strings and then to array of floats
  columnValuesStr = strings.Fields(line)
  for _, value := range columnValuesStr {
    n, _ := strconv.ParseFloat(value, 64)
    columnValuesFlt = append(columnValuesFlt, n)
  }

  return columnValuesFlt
}
