package iban

import (
	"unicode"
)

type IbanValidator struct {
}

func (iv *IbanValidator) IsValid(s string) (ret bool) {
	runes := []rune(s)
	form_test_1 := (len(s) <= 34)
	form_test_2 := unicode.IsLetter(runes[0]) && unicode.IsLetter(runes[1])
	form_test_3 := unicode.IsDigit(runes[2]) && unicode.IsDigit(runes[3])

	form_test_full := form_test_1 && form_test_2 && form_test_3
	return form_test_full
}
