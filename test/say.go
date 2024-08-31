package test

import "strings"

func Say(names []string) string {
	if len(names) > 1 {
		return "Hello, " + strings.Join(names, " ")
	}
	return "Hello, World"
}
