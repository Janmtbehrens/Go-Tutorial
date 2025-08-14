package katabankocr

import (
	"os"
	"reflect"
	"strconv"
	"strings"
)

func readLine(path string, row int) []string {
	file, err := os.ReadFile(path)
	content := string(file)

	if err != nil {
		println(err)
	}

	var writtenBankOCR []string
	for i, line := range strings.Split(content, "\n") {
		if i-row < 3 {
			writtenBankOCR = append(writtenBankOCR, line+"     ")
		}
	}
	return writtenBankOCR
}

func readLetter(lines []string, offset int) []string {
	var letter []string
	for i := 0; i < 3; i++ {
		letter = append(letter, lines[i][offset:offset+3])
	}
	return letter
}

func interpretOCRLine(line []string) string {
	rowOffset := 0
	output := ""
	for _, l := range line {
		println(l)
	}
	for {
		letter := readLetter(line, rowOffset)

		output += interpretLetter(letter)

		rowOffset += 3

		if output == "void" {
			break
		}
	}
	return "123456789"
}

func interpretLetter(letter []string) string {
	for _, l := range letter {
		println(l)
	}

	// Void check
	letterCount := 0
	for _, l := range letter {
		l = strings.TrimSpace(l)
		letterCount += len(l)
	}

	for i := 0; i < 10; i++ {
		println("Reading line " + strconv.Itoa(i))
		knownLine := readLine("knownNumbers.txt", i+1)

		same := true
		for j := 0; j < 3; j++ {
			if !reflect.DeepEqual(strings.TrimSpace(letter[j]), strings.TrimSpace(knownLine[j])) {
				same = false
			}
		}
		println(same)
		if same {
			return strconv.Itoa(i)
		}
	}

	return "?"
}

func Read(path string, row int) (string, error) {
	writtenBankOCR := readLine(path, row)

	output := interpretOCRLine(writtenBankOCR)

	return output, nil
}
