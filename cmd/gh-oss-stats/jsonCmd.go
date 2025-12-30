package main

import (
	"flag"
	"fmt"
	"os"
)

// flag set
var fromJSONCmd = flag.NewFlagSet("json", flag.ExitOnError)

// flags
var data = fromJSONCmd.String("data", "", "Stats as json string")

func runJSONCmd() {
	fmt.Printf("data=%s\n", *data)
	os.Exit(0)
}
