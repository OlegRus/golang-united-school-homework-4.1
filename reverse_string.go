package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var b strings.Builder
	for _, r := range []rune(input) {
		b.WriteRune(r)
	}
	return b.String()
}
