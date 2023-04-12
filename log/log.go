package log

import (
	"fmt"
  "time"

	"github.com/TwiN/go-color"
)

// LogInfo: logging function, INFO level
func LogInfo(reference, data string) {

  dt := time.Now().Format("01-02-2006 15:04:05.000")

	fmt.Println(color.Ize(color.Green, dt + " --INFO-- " + data + " [" + reference + "]"))
}

// LogInfo: logging function, ERROR level
func LogError(reference, data string) {

  dt := time.Now().Format("01-02-2006 15:04:05.000")

	fmt.Println(color.Ize(color.Red, dt + " --ERROR-- " + data + " [" + reference + "]"))
}

// LogInfo: logging function, DEBUG level
func LogDebug(reference, data string) {

  dt := time.Now().Format("01-02-2006 15:04:05.000")

	fmt.Println(color.Ize(color.Cyan, dt + " --DEBUG-- " + data + " [" + reference + "]"))
}
