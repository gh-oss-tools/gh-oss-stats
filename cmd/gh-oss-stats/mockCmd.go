package main

import (
	"flag"
	"fmt"
)

var mockCmd = flag.NewFlagSet("mock", flag.ExitOnError)

func runMockCmd(args []string) {
	badgeConfig.registerBadgeFlags(mockCmd)
	mockCmd.Parse(args)
	fmt.Println("Mockkkkkkkkkkkkkkk ")
}
