package hello

import "strings"

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"Go"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
