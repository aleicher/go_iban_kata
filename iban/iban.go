package iban

type IbanValidator struct {
}

// The following three functions are kind of unnecessary.
// because I don't know how to do it for a single byte in Go
func IsLower(c uint8) bool {
  return (('a' <= c) && (c <= 'z'))
}

func IsUpper(c uint8) bool {
  return (('A' <= c) && (c <= 'Z'))
}

func IsAlpha(c uint8) bool {
  return IsUpper(c) || IsLower(c)
}

func (iv *IbanValidator) IsValid(s string) (ret bool) {
  form_test_1 := (len(s) <= 34)
  form_test_2 := IsAlpha(s[0]) && IsAlpha(s[1])

  form_test_full := form_test_1 && form_test_2
  return form_test_full
}
