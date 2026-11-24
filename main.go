package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "marauder: failed to get hostname: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("marauder: running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("hostname: %s\n", name)

	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read:", err)
	} else {
		fmt.Println("os-release:")
		fmt.Print(string(data))
	}
}
