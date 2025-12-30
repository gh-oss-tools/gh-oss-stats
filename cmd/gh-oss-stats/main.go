package main

import (
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
		runMainCmd(os.Args[1:])
		return
	}

	switch os.Args[1] {
	case "json":
		runJSONCmd(os.Args[2:])
	case "mock":
		runMockCmd(os.Args[2:])
	case "version":
		fmt.Printf("gh-oss-stats v%s\n", version)
		os.Exit(0)
	default:
		runMainCmd(os.Args[1:])
		os.Exit(0)
	}
}
