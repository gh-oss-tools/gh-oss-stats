package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mabd-dev/gh-oss-stats/pkg/ossstats/badge"
)

const version = "0.3.1"

func main() {
	badgeConfig = BadgeConfig{
		generate: false,
		style:    string(badge.DefaultBadgeStyle),
		variant:  string(badge.DefaultBadgeVariant),
		theme:    string(badge.DefaultBadgeTheme),
		output:   "",
		sort:     string(badge.DefaultSortBy),
		limit:    badge.DefaultPRsLimit,
	}

	if len(os.Args) < 2 {
		mainCmd()
		return
	}

	switch os.Args[1] {
	case "json":
		badgeConfig.registerBadgeFlags(fromJSONCmd)
		fromJSONCmd.Parse(os.Args[2:])
		runJSONCmd()
	case "mock":
		badgeConfig.registerBadgeFlags(mockCmd)
		mockCmd.Parse(os.Args[2:])
		runMockCmd()
	case "version":
		fmt.Printf("gh-oss-stats v%s\n", version)
		os.Exit(0)
	default:
		mainCmd()
		os.Exit(0)
	}
}

func mainCmd() {
	badgeConfig.registerBadgeFlags(flag.CommandLine)
	flag.Parse()
	//flag.CommandLine.Parse(os.Args[1:])
	runMainCmd()
}
