package mesastate

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// SetStarHistoryName defines the name of the history file of MESAstar
func (M *MESAEvolution) SetStarHistoryName (id int, filename string) {
  
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
  M.Binary.HistoryName = filename
}

// SetColumnNames stores the names of the columns of the MESA output
func (M *MESAEvolution) SetColumnNames () {

  var (
    NR_COLUMN_NAMES = 6
  )

  if M.IsBinary {

    // open binary file
    f, err := os.Open(M.Binary.HistoryName)
    if err != nil {
      return
    }
    defer f.Close()

    // scan file
    scanner := bufio.NewScanner(f)

    columnNamesFound := false
    lineCount := 0
    for scanner.Scan(){

      lineCount++
      
      if lineCount == NR_COLUMN_NAMES {
        columnNames := strings.Fields(scanner.Text())
        columnNamesFound = true

        fmt.Println(columnNames)

      }

      if columnNamesFound {
        break
      }
    }
  }

}
