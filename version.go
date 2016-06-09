package main

import (
	"fmt"
	"runtime"
)

const FULLSTACK_VERSION = "0.0.1"

func GetVersion() string {
	return fmt.Sprintf("Fullstack version:\t%s\nGo version:\t\t%s\nSystem architecture:\t%s/%s\n", FULLSTACK_VERSION, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
