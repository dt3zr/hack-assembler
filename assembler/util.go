package assembler

import (
	"os"
)

var nullWriter *os.File

func init() {
	nullWriter, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 777)
}
