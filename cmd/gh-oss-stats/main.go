package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.3.1"

func main() {
	if len(os.Args) < 2 {
		flag.Parse()
		runMainCmd()
		return
	}

	switch os.Args[1] {
	case "json":
		fromJSONCmd.Parse(os.Args[2:])
		runJSONCmd()
	case "mock":
		mockCmd.Parse(os.Args[2:])
		runMockCmd()
	case "version":
		fmt.Printf("gh-oss-stats v%s\n", version)
		os.Exit(0)
	default:
		flag.Parse()
		runMainCmd()
	}
}
