// Package cli provides some utility functions for command line interfaces
package cli

import "github.com/phR0ze/n"

// ParseCliOpts parses command line options
func ParseCliOpts(opts []string) map[string]interface{} {
	return n.Q(opts).MapF(func(x n.O) n.O {
		return n.A(x.(string)).Split(",").Map(func(y string) n.O {
			return n.A(y).Split("=").YamlKeyVal()
		})
	}).M()
}
