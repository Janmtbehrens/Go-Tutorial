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

//func (r *Roman) ToDigit() int {
//	return 0
//}

func ToRoman(number int) Roman {
	output := Roman("")
	numberAsString := strconv.Itoa(number)
	for i, ch := range numberAsString {
		roman := romanNumerals()[i*2]
		num := int(ch - '0')

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
		return Roman(r.Successor().Successor())
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
