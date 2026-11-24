package main

import (
	"fmt"
	"os"
	"runtime"
)

func fileEnum(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read:", err)
		return
	}
	fmt.Printf("%s:\n", path)
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
	paths := []string{
		"/etc/passwd",
		"/etc/shadow",
		"/etc/hosts",
		"/etc/resolv.conf",
		"wrong/path/to/file",
	}
	for _, path := range paths {
		fileEnum(path)
	}
}
