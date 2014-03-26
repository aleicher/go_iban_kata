package iban_test

import (
	. "github.com/aleicher/go_iban_kata/iban"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IbanValidator", func() {
	var iban *IbanValidator
	validIbanFixture := "DE89370400440532013000"
	BeforeEach(func() {
		iban = new(IbanValidator)
	})
	Describe("Validation 1: maximum length", func() {
		tooLongIban := "12345678901234567890123456789012345"

		It("does return true for a valid IBAN", func() {
			Expect(iban.IsValid(validIbanFixture)).To(Equal(true))
		})
		It("does return false for an IBAN with more than 34 characters", func() {
			Expect(iban.IsValid(tooLongIban)).To(Equal(false))
		})
	})
	Describe("Validation 2: first two positions are characters", func() {
		noCharIban := "1234567890123456789012345678901234"
		oneCharIban := "X234567890123456789012345678901234"
		It("does return false for an IBAN which starts with numbers", func() {
			Expect(iban.IsValid(noCharIban)).To(Equal(false))
		})
		It("does return false for an IBAN which starts with one character", func() {
			Expect(iban.IsValid(oneCharIban)).To(Equal(false))
		})
		It("does return true for the valid IBAN", func() {
			Expect(iban.IsValid(validIbanFixture)).To(Equal(true))
		})

	})
})
