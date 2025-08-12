package kataflag

import (
	"strconv"
	"strings"
)

type FlagValue struct {
	Bool    bool
	Integer int
	String  string
	Array   []string
}

// InputType = bool, string, stringList, int
type Flag struct {
	Identifier string
	InputType  int
	Default    FlagValue
}

type Schema struct {
	flags []Flag
}

func GetSchema() Schema {
	flags := []Flag{
		{"d", 0, FlagValue{}},                   // Debug Flag
		{"v", 0, FlagValue{}},                   // Verbose Flag
		{"p", 1, FlagValue{String: "."}},        // File path
		{"args", 2, FlagValue{}},                // Any amount of string args
		{"timestamp", 3, FlagValue{Integer: 1}}, // Some timestamp
	}

	schema := Schema{flags}

	return schema
}

func FindFlag(schema Schema, flagIdentifier string) Flag {
	for _, flag := range schema.flags {
		if flag.Identifier == flagIdentifier {
			return flag
		}
	}
	return Flag{"", -1, FlagValue{}}
}

func (f Flag) Evaluate(s string) FlagValue {
	index := strings.Index(s, "-"+f.Identifier)

	// No flag provided, use default
	if index == -1 {
		return f.Default
	}

	ret := FlagValue{Bool: true}

	content := s[strings.Index(s, "-"+f.Identifier)+len(f.Identifier)+1:]
	content = strings.TrimSpace(content)

	if f.InputType == 3 {
		// Integer case
		minusIndex := strings.Index(content, "-")
		offset := 0
		if minusIndex == 0 {
			// We are an integer and there is a - right next to our flag
			offset = 1
		}

		if strings.Contains(content[offset:], "-") {
			// Cut off all other Flags
			ret.String = content[0:strings.Index(content[offset:], "-")]
		}
		ret.String = strings.TrimSpace(ret.String)

		num, err := strconv.Atoi(ret.String)

		if err != nil {
			println(err.Error())
		}

		ret.Integer = num

	} else {
		if strings.Contains(content, "-") {
			ret.String = content[0:strings.Index(content, "-")]
		}
		ret.String = strings.TrimSpace(ret.String)

		if f.InputType == 2 {
			// Inteprete content as a list
			ret.Array = strings.Split(ret.String, ",")
		}
	}

	return ret
}

/*
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

	if flag.InputType == 2 {
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

func IsVerbose(schema Schema, input string) bool {
	exists := false

	flag := findFlag(schema, "v")

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
*/
