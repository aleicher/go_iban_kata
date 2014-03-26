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
  ret = (len(s) <= 34)
  ret = ret && IsAlpha(s[0]) && IsAlpha(s[1])
  return ret
}
