package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var b strings.Builder
	runes := []rune(input)
	for i := range runes {
		b.WriteRune(runes[len(runes)-1-i])
	}
	return b.String()
}
