package cliutil

import "strings"

// ParseOptions 解决cli bug
func ParseOptions(args []string, options []string) ([]string, map[string]string) {
	var result = make(map[string]string)
	var rest []string
	// option --x=xxx
	for _, arg := range args {
		var remove bool
		// start with --=
		if strings.HasPrefix(arg, "--") {
			for _, option := range options {
				// --x=xxx
				if strings.HasPrefix(arg, "--"+option+"=") {
					result[option] = strings.TrimPrefix(arg, "--"+option+"=")
					remove = true
				}
			}
		}
		if !remove {
			rest = append(rest, arg)
		}
	}
	return rest, result
}
