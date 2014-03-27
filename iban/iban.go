package iban

import (
  //"fmt"
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

func positionInAlphabet(x rune) int {
  x = unicode.ToLower(x)
  return int(x - 'a')
}

func (iv *IbanValidator) Checksum(s string) int {
  runes := []rune(s)

  // if it is NOT in the right form, then check might even crash.
  // so, return wrong checksum.
  if IsIBANFormValid(runes) == false {
    return -1
  }

  // if valid, call the private (?) method.
  return checksum(runes)
}

func checksum(runes []rune) int {
  modulo := 97
  modulo_sum := 0

  // NOTE: Modulo addition, is addition modulo. (also multiplication)
  // So, here we avoid creating that HUGE 30+ digit number, which will
  // be about 2^100, which will need more than 64 bits and hence could
  // crash everything. [In our scenario, the largest value of modulo_sum
  // will be 96, nothing more.]
  for _, x := range runes[4:] {
    modulo_sum = (modulo_sum*10 + int(x-'0')) % modulo
  }

  // the first two alphabets. Since the give a two digit number,
  // multiply the existing modulo_sum by 100 to create enough place on
  // the right side, then add the values, take modulo.
  modulo_sum = (modulo_sum*100 + (positionInAlphabet(runes[0]) + 10)) % modulo
  modulo_sum = (modulo_sum*100 + (positionInAlphabet(runes[1]) + 10)) % modulo

  // the 3rd and 4th characters, the digits. In a single go.
  modulo_sum = (modulo_sum*100 + (int(runes[2]-'0') * 10) + int(runes[3]-'0')) % modulo

  return modulo_sum
}

func (iv *IbanValidator) IsValid(s string) bool {
  runes := []rune(s)

  // does the given string satisfy the general form of an IBAN?
  form_test := IsIBANFormValid(runes)

  // does the check sum meet the requirement?
  checksum_test := false
  if form_test {
    checksum_test = (checksum(runes) == 1)
  }

  // further requirements.

  return form_test && checksum_test
}
