package iteration

import (
	"strings"
)

func Repeat(char string, count int) string {
	if count == 0 || count < 0 {
		return ""
	}

	var repeated strings.Builder
	for range count {
		repeated.WriteString(char)
	}
	return repeated.String()
}
