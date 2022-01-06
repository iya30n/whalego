package errorHandler

import "whalego/file"

func LogFile(err error) {
	if err != nil {
		file.AddOrCreate("./errors.txt", err.Error())
	}
}