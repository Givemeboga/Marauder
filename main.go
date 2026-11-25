package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
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

func lineParser() {
	data, err := os.ReadFile("/etc/passwd")
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not read:", err)
		return
	}
	text := string(data)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) == 7 {
			fmt.Printf("User: %s, UID: %s, GID: %s, Home: %s, Shell: %s\n", parts[0], parts[2], parts[3], parts[5], parts[6])
		}
	}
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
		"/etc/os-release",
		"/proc/version",
	}
	for _, path := range paths {
		fileEnum(path)
	}
	lineParser()
}
