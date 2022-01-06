package file

import (
	"bufio"
	"errors"
	"os"
)

func NewFile(path string, value string) {
	file, err := os.Create(path)

	// TODO: use err handler here
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(value)

	file.Sync()
}

func AddToFile(path string, value string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// TODO: use err handler here
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("\n")
	writer.WriteString(value)

	writer.Flush()
}

func AddOrCreate (path string, value string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		NewFile(path, value)
	} else {
		AddToFile(path, value)
	}
}