package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// fallthrough
	case "linux":
		fmt.Println("Linux.")
	default:
		// FreeBSD, openBSD,
		// plan9, windows, ...
		fmt.Printf("%s.", os)
	}
}
