package markdown

import "strings"

func Prepare(str string) string {
	newStr := strings.ReplaceAll(str, "\n", "\n\n")

	return newStr
}
