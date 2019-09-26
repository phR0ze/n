package sys

import (
	"regexp"
)

var gRXSplitCmd = regexp.MustCompile(`'.+'|".+"|\S+`)

// SplitCmd splits this cmd into substrings around spaces taking into account bash like
// double and single quotes. Unmatched quotes throw and error and empty quotes are removed.
func SplitCmd(cmd string) (slice []string) {
	return gRXSplitCmd.FindAllString(cmd, -1)
}

func slice(x []string, i, j int) (result []string) {
	l := len(x)

	// Convert to postive notation
	if i < 0 {
		i = l + i
	}
	if j < 0 {
		j = l + j
	}

	// Move start/end within bounds
	if i < 0 {
		i = 0
	}
	if j >= l {
		j = l - 1
	}

	// Specifically offsetting j to get an inclusive behavior out of Go
	j++

	// Only operate when indexes are within bounds
	// allow j to be len of s as that is how we include last item
	if i >= 0 && i < l && j >= 0 && j <= l {
		result = x[i:j]
	} else {
		result = []string{}
	}
	return
}
