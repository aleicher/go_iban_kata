package iban

type IbanValidator struct {
}

func (iv *IbanValidator) IsValid(s string) (ret bool) {
  ret = (len(s) <= 34)

  return ret
}
