package main

import (
	"flag"

	"github.com/mabd-dev/gh-oss-stats/pkg/ossstats/badge"
)

var badgeConfig BadgeConfig

type BadgeConfig struct {
	generate bool
	style    string
	variant  string
	theme    string
	output   string
	sort     string
	limit    int
}

func (bf *BadgeConfig) registerBadgeFlags(fs *flag.FlagSet) {
	fs.BoolVar(&bf.generate, "badge", false, "Generate SVG badge")
	fs.StringVar(&bf.style, "badge-style", string(badge.DefaultBadgeStyle), "Badge style: summary, compact, detailed")
	fs.StringVar(&bf.variant, "badge-variant", string(badge.DefaultBadgeVariant), "Badge variants: default, text-based")
	fs.StringVar(&bf.theme, "badge-theme", string(badge.DefaultBadgeTheme), "Badge theme: dark, light, nord, dracula, ...")
	fs.StringVar(&bf.output, "badge-output", "", "Badge output file (default: badge.svg)")
	fs.StringVar(&bf.sort, "badge-sort", string(badge.DefaultSortBy), "Sort contributions by: prs, stars, commits")
	fs.IntVar(&bf.limit, "badge-limit", badge.DefaultPRsLimit, "Number of contributions to show")
}
