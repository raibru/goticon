package cmd

import (
	"fmt"
	"io"
)

// var section
var (
	major     = "0"
	minor     = "1"
	patch     = "0"
	buildTag  = "-"
	buildDate = "-"
	appName   = "pktfmt - packet formatter"
	author    = "raibru <github.com/raibru>"
	license   = "MIT License (c) 2020 raibru"
)

// PrintVersion prints the tool versions string
func PrintVersion(w io.Writer) {
	fmt.Fprintf(w, "%s - v%s.%s.%s\n", appName, major, minor, patch)
	fmt.Fprintf(w, "  Build-%s (%s)\n", buildTag, buildDate)
	fmt.Fprintf(w, "  author : %s\n", author)
	fmt.Fprintf(w, "  license: %s\n\n", license)
}
