package romankata

import "strconv"

type Roman string

func romanNumerals() []Roman {
	return []Roman{"I", "V", "X", "L", "C", "D", "M"}
}

func associateRomanNumeral(r Roman) int {
	for i, num := range romanNumerals() {
		if num == r {
			return i
		}
	}
	return -1
}

func (r Roman) Predecessor() Roman {
	entry := associateRomanNumeral(r)
	if entry == 0 {
		return ""
	}
	return romanNumerals()[entry-1]
}

func (r Roman) Successor() Roman {
	entry := associateRomanNumeral(r)
	if entry == len(romanNumerals())-1 {
		return ""
	}
	return romanNumerals()[entry+1]
}

func reverseString(s string) string {
	// Convert the string to a slice of runes to handle Unicode characters correctly.
	runes := []rune(s)

	// Use two pointers, 'i' starting from the beginning and 'j' from the end.
	// Iterate until 'i' crosses 'j'.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Swap the runes at positions 'i' and 'j'.
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the modified slice of runes back to a string and return it.
	return string(runes)
}

//func (r *Roman) ToDigit() int {
//	return 0
//}

func ToRoman(number int) Roman {
	output := Roman("")
	numberAsString := strconv.Itoa(number)
	numberAsString = reverseString(numberAsString)
	entry := 0
	for _, ch := range numberAsString {
		roman := romanNumerals()[entry*2]
		num := int(ch - '0')
		entry++

		if num == 0 {
			continue
		}
		output = romanRow(roman, num) + output
	}
	return output
}

func romanRow(r Roman, rowNum int) Roman {
	switch rowNum {
	case 0:
		return ""
	case 1:
		return Roman(r)
	case 2:
		return Roman(r + r)
	case 3:
		return Roman(r + r + r)
	case 4:
		return Roman(r + r.Successor())
	case 5:
		return Roman(r.Successor())
	case 6:
		return Roman(r.Successor() + r)
	case 7:
		return Roman(r.Successor() + r + r)
	case 8:
		return Roman(r.Successor() + r + r + r)
	case 9:
		return Roman(r + r.Successor().Successor())
	}
	return "---"
}

func (r Roman) ToDigit() int {
	output := 0

	runeRoman := []rune(r)

	for range runeRoman {
		lastFourChars := runeRoman[len(runeRoman)-4:]

		num, cut := reverseRomanRow(Roman(lastFourChars))
		output += num
		runeRoman = runeRoman[:len(runeRoman)-cut]
	}

	return 0
}

func reverseRomanRow(input Roman) (int, int) {
	r := Roman(input[0])
	if r == "V" || r == "L" || r == "D" {
		r = r.Predecessor()
	}
	mult := 0
	switch r {
	case "I":
		mult = 1
	case "X":
		mult = 10
	case "C":
		mult = 100
	case "M":
		mult = 1000
	}

	switch input {
	case "":
		return mult * 0, 0
	case r:
		return mult * 1, 1
	case r + r:
		return mult * 2, 2
	case r + r + r:
		return mult * 3, 3
	case r + r.Successor():
		return mult * 4, 2
	case r.Successor():
		return mult * 5, 1
	case r.Successor() + r:
		return mult * 6, 2
	case r.Successor() + r + r:
		return mult * 7, 3
	case r.Successor() + r + r + r:
		return mult * 8, 4
	case r + r.Successor().Successor():
		return mult * 9, 2
	}
	return 0, 0
}
