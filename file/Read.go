package file

import "os"

func Read(path string)[]byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}