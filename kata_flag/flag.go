package kataflag

import "strings"

// InputType = int, string
type Flag struct {
	Identifier             string
	ExpectedParameterCount int
	InputType              int
}

type Schema struct {
	flags []Flag
}

func GetSchema() Schema {
	flags := []Flag{
		{"d", 0, 0},     // Debug Flag
		{"args", -1, 1}, // Any amount of string args
		{"p", 1, 1},     // File path
	}

	schema := Schema{flags}

	return schema
}

func findFlag(schema Schema, flagIdentifier string) Flag {
	for _, flag := range schema.flags {
		if flag.Identifier == flagIdentifier {
			return flag
		}
	}
	return Flag{"", 0, 0}
}

func interpreteInput(flag Flag, input string) (bool, string, []string) {

	index := strings.Index(input, "-"+flag.Identifier)

	if index == -1 {
		return false, "", []string{""}
	}

	content := input[strings.Index(input, "-"+flag.Identifier)+len(flag.Identifier)+1:]

	if strings.Contains(content, "-") {
		content = content[0:strings.Index(content, "-")]
	}

	content = strings.TrimSpace(content)

	if flag.ExpectedParameterCount == -1 {
		// Inteprete content as a list
		split := strings.Split(content, ",")
		return true, content, split
	}

	return true, content, []string{""}
}

func InDebugMode(schema Schema, input string) bool {
	exists := false

	flag := findFlag(schema, "d")

	exists, _, _ = interpreteInput(flag, input)

	return exists
}

func GetPath(schema Schema, input string) string {
	output := ""

	flag := findFlag(schema, "p")

	_, output, _ = interpreteInput(flag, input)

	return output
}

func GetArgs(schema Schema, input string) []string {
	output := []string{""}

	flag := findFlag(schema, "args")

	_, _, output = interpreteInput(flag, input)

	return output
}
