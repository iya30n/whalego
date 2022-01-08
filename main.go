package main

import (
	"os"
)

func newFile (val []byte) {
    f, err := os.Create("./result.json")
    if err != nil {
        panic(err)
    }

    defer f.Close()

    f.Write(val)

    f.Sync()
}

func main() {
    
}
