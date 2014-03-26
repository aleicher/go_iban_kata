package iban_test

import (
	. "github.com/aleicher/go_iban_kata/iban"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IbanValidator", func() {
	var iban *IbanValidator
	shortIban := "123456"
	longIban := "1234567890123456789012345678901234"
	tooLongIban := "12345678901234567890123456789012345"
	BeforeEach(func() {
		iban = new(IbanValidator)
	})
	Describe("Validation 1: maximum length", func() {
		It("does provide a function to check if an IBAN is valid", func() {
			Expect(iban.IsValid(shortIban)).To(Equal(true))
		})
		It("does return true for a IBAN with 34 characters", func() {
			Expect(iban.IsValid(longIban)).To(Equal(true))
		})
		It("does return false for an IBAN with more than 34 characters", func() {
			Expect(iban.IsValid(tooLongIban)).To(Equal(false))
		})
	})
})
