package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mabd-dev/gh-oss-stats/pkg/ossstats"
)

// badgeCmd flag set
var badgeCmd = flag.NewFlagSet("badge", flag.ExitOnError)

// Badge command flags
var (
	badgeFromFile = badgeCmd.String("from-file", "", "Path to stats JSON file")
	badgeData     = badgeCmd.String("data", "", "Stats as JSON string")
)

func init() {
	badgeCmd.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gh-oss-stats badge [options]\n\n")
		fmt.Fprintf(os.Stderr, "Generate badge from existing stats JSON.\n\n")
		fmt.Fprintf(os.Stderr, "This command allows you to generate badges without re-fetching data from GitHub,\n")
		fmt.Fprintf(os.Stderr, "which is useful for creating multiple badge variants from the same stats.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		badgeCmd.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Generate badge from file\n")
		fmt.Fprintf(os.Stderr, "  gh-oss-stats badge --from-file stats.json --badge-style summary\n\n")
		fmt.Fprintf(os.Stderr, "  # Generate badge from JSON string\n")
		fmt.Fprintf(os.Stderr, "  gh-oss-stats badge --data '{\"username\":\"...\",...}' --badge-style compact\n\n")
	}
}

func runBadgeCmd(args []string) {
	badgeConfig := newBadgeConfig()
	badgeConfig.registerBadgeFlags(badgeCmd)
	badgeCmd.Parse(args)

	var stats ossstats.Stats
	err := json.Unmarshal([]byte(statsJson), &stats)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	badgeOption, err := createBadgeOptions(*badgeConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := writeBadge(
		badgeOption,
		badgeConfig.output,
		verbose,
		&stats,
	); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating badge: %v\n", err)
		os.Exit(1)
	}

}

// later i will read json file or json string from command flags
var statsJson = `
	{ "username": "mabd-dev", "generatedAt": "2025-12-31T05:33:15.031869Z", "summary": { "totalProjects": 7, "totalPRsMerged": 17, "totalCommits": 17, "totalAdditions": 0, "totalDeletions": 0 }, "contributions": [ { "repo": "ibad-al-rahman/android-public", "owner": "ibad-al-rahman", "repoName": "android-public", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 6, "commits": 6, "additions": 0, "deletions": 0, "firstContribution": "2025-11-21T14:48:30Z", "lastContribution": "2025-12-17T05:14:39Z" }, { "repo": "nsh07/Tomato", "owner": "nsh07", "repoName": "Tomato", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 2, "commits": 2, "additions": 0, "deletions": 0, "firstContribution": "2025-11-19T12:06:16Z", "lastContribution": "2025-11-21T05:45:34Z" }, { "repo": "qamarelsafadi/JetpackComposeTracker", "owner": "qamarelsafadi", "repoName": "JetpackComposeTracker", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 2, "commits": 2, "additions": 0, "deletions": 0, "firstContribution": "2025-06-14T20:55:24Z", "lastContribution": "2025-07-21T21:39:53Z" }, { "repo": "android/nav3-recipes", "owner": "android", "repoName": "nav3-recipes", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 2, "commits": 2, "additions": 0, "deletions": 0, "firstContribution": "2025-06-09T18:24:56Z", "lastContribution": "2025-06-09T18:31:50Z" }, { "repo": "android/cahier", "owner": "android", "repoName": "cahier", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 2, "commits": 2, "additions": 0, "deletions": 0, "firstContribution": "2025-06-03T14:08:20Z", "lastContribution": "2025-07-11T12:52:46Z" }, { "repo": "esatgozcu/Compose-Rolling-Number", "owner": "esatgozcu", "repoName": "Compose-Rolling-Number", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 2, "commits": 2, "additions": 0, "deletions": 0, "firstContribution": "2025-02-17T15:46:51Z", "lastContribution": "2025-03-26T21:33:08Z" }, { "repo": "zuzmuz/nvimawscli", "owner": "zuzmuz", "repoName": "nvimawscli", "description": "Android app for Ibad Al-Rahman", "repoURL": "https://github.com/ibad-al-rahman/android-public", "stars": 15, "prsMerged": 1, "commits": 1, "additions": 0, "deletions": 0, "firstContribution": "2024-05-06T20:37:13Z", "lastContribution": "2024-05-06T20:37:13Z" } ] }
`
