package tool

import (
	"bytes"
	"strings"
	"time"
)

// GetFileNameWithTimestamp s
func GetFileNameWithTimestamp(fileName string) (formatFileName string) {
	var buffer bytes.Buffer
	dot := strings.LastIndex(fileName, ".")
	timeOfNow := time.Now()
	if dot >= 0 {
		buffer.WriteString(fileName[:dot])
		buffer.WriteString(timeOfNow.Format("_20060102150405"))
		buffer.WriteString(fileName[dot:])
	} else {
		buffer.WriteString(fileName)
		buffer.WriteString(timeOfNow.Format("_20060102150405"))
		buffer.WriteString(".log")
	}
	formatFileName = buffer.String()
	return
}
