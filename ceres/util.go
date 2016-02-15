package ceres


import (
	"fmt"
	"os"
	"strconv"
)

func checkerr(e error, label string, fatal bool) {
	if e == nil {
		return
	}

	fmt.Printf("%s: %s\n", label, e)

	if (fatal) {
		os.Exit(1)
	}
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)

	if e == nil {
		return i
	}

	return 0
}