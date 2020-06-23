package cmd

import (
	"fmt"
	"io"
)

var (
	// Major version number
	major = "0"
	// Minor version number
	minor = "1"
	// Patch version number
	patch = "0"
	// BuildVersion version string of build with repro and date
	buildVersion = ""
	// BuildTag version tag string of build with repro and a build number
	buildTag = ""
	// BuildDate version string of build date
	buildDate = ""
)

// PrintVersion prints the tool version string into writer
func PrintVersion(w io.Writer) {
	//fmt.Fprintf(w, "bitpacks - bit package structure generator - v0.0.1\n")
	fmt.Fprintf(w, "pktfmt - packet formatter - v%s.%s.%s\n", major, minor, patch)
	fmt.Fprintf(w, "  Build-%s (%s)\n", buildTag, buildDate)
	fmt.Fprintf(w, "  author : raibru <raibru@github.com>\n")
	fmt.Fprintf(w, "  license: MIT\n\n")
}
