package sys

import "regexp"

var gRXSplitCmd = regexp.MustCompile(`'.+'|".+"|\S+`)

// SplitCmd splits this cmd into substrings around spaces taking into account bash like
// double and single quotes. Unmatched quotes throw and error and empty quotes are removed.
func SplitCmd(cmd string) (slice []string) {
	return gRXSplitCmd.FindAllString(cmd, -1)
}
