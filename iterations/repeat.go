package iterations

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	var repeated strings.Builder
	for range repeatCount {
		repeated.WriteString(char)
	}

	return repeated.String()
}
