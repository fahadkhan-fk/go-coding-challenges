package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Unable to get current file info")
		return
	}

	dir := filepath.Dir(filename)
	base := filepath.Base(filename)

	fmt.Println("Project directory:", dir)
	fmt.Println("Filename:", base)
}
