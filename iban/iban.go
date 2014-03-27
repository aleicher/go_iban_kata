package iban

import (
  "unicode"
)

type IbanValidator struct {
}

func IsIBANFormValid(runes []rune) bool {

  // a valid iban must be shorter than 34 characters
  form_test_1 := (len(runes) <= 34)

  // if the following tests are not to fail, it must be at least 4 char long.
  if len(runes) < 4 {
    return false
  }

  // the first two characters need to be alphabets
  form_test_2 := unicode.IsLetter(runes[0]) && unicode.IsLetter(runes[1])

  // the next two must be digits
  form_test_3 := unicode.IsDigit(runes[2]) && unicode.IsDigit(runes[3])

  // all the remaining characters must be digits too.
  form_test_4 := true
  for _, x := range runes[4:] {
    if unicode.IsDigit(x) == false {
      form_test_4 = false
      // break off at the first fail
      break
    }
  }

  return form_test_1 && form_test_2 && form_test_3 && form_test_4
}

func (iv *IbanValidator) IsValid(s string) bool {

  // does the given string satisfy the general form of an IBAN?
  form_test := IsIBANFormValid([]rune(s))

  return form_test
}
