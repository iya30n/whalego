package errorHandler

import (
	"os"
	"whalego/file"
)

func LogFile(err error) {
	if err != nil {
		file.AddOrCreate("./errors.txt", err.Error())

		os.Exit(1)
	}
}