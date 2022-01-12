package errorHandler

import (
	// "fmt"
	"fmt"
	"os"
	"runtime"
	"whalego/file"
)

func LogFile(err error) {
	if err != nil {
		_, errfile, line, _ := runtime.Caller(1)

		errorMessage := err.Error() + " (file: " + errfile + ":%d)"

		errorMessage = fmt.Sprintf(errorMessage, line)

		file.AddOrCreate("./errors.txt", errorMessage)

		os.Exit(1)
	}
}