package main

import (
	"fmt"
	"os"
)

func info(msg string) {
	fmt.Printf("INFO: %s\n", msg)
}

func warn(msg string) {
	fmt.Fprintf(os.Stderr, "WARNING: %s\n", msg)
}

func fatal(msg string, err error) {
	fmt.Fprintf(os.Stderr, "FATAL: %s \n ============ \n %s \n", msg, err)
	os.Exit(1)
}