package main

import (
	"fmt"
	"os"
)

var version = "version"
var commit = "commit"
var date = "date"

func main() {
	os.Exit(runMain())
}

func runMain() int {
	fmt.Println("Version:", version)
	fmt.Println("Commit Hash:", commit)
	fmt.Println("Build Date:", date)
	return 0
}
