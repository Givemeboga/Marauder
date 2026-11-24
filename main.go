package main

import (
	"fmt"
	"os"
	"runtime"
)

func fileEnum(header string, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read:", err)
		return
	}
	fmt.Println(header)
	fmt.Print(string(data))
}

func main() {
	name, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "marauder: failed to get hostname: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("marauder: running on %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("hostname: %s\n", name)
	fileEnum("os-release:", "/etc/os-release")
	fileEnum("kernel:", "/proc/version")
}
