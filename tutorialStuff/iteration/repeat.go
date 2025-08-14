package iteration

import "strings"

func Repeating(input string, times int) string {
	var output strings.Builder
	for i := 0; i < times; i++ {
		output.WriteString(input)
	}
	return output.String()
}

func GotAnA(input string) bool {
	return strings.Contains(input, "a")
}