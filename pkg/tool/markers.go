package tool

import "strconv"

type Marker interface {
	Value(number int) string
}

type AlphabeticMarker struct {
	Offset int
}

var alphabetics = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (m AlphabeticMarker) Value(number int) string {
	number += m.Offset
	str := ""
	for ; number >= len(alphabetics); {
		str = string(alphabetics[number % len(alphabetics)]) + str
		number = number / len(alphabetics) - 1
	}
	str = string(alphabetics[number % len(alphabetics)]) + str
	return str
}

type NumericMarker struct {
	Offset int
}

func (m NumericMarker) Value(number int) string {
	return strconv.Itoa(number + m.Offset)
}


