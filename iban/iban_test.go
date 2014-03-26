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
	Describe("Validation 0: valid for valid IBAN :)", func() {
		It("does return true for a valid IBAN", func() {
			Expect(iban.IsValid(validIbanFixture)).To(Equal(true))
		})
	})

	Describe("Validation 1: maximum length", func() {
		tooLongIban := "12345678901234567890123456789012345"

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
	})

	Describe("Validation 3: positions 3 and 4 are numbers", func() {
		posThreeWrong := "XXX4567890123456789012345678901234"
		posFourWrong := "XX3X567890123456789012345678901234"
		posThreeAndFourWrong := "XXXX567890123456789012345678901234"
		It("does return false if position 3 is not a number", func() {
			Expect(iban.IsValid(posThreeWrong)).To(Equal(false))
		})
		It("does return false if position 4 is not a number", func() {
			Expect(iban.IsValid(posFourWrong)).To(Equal(false))
		})
		It("does return false if position 4 is not a number", func() {
			Expect(iban.IsValid(posThreeAndFourWrong)).To(Equal(false))
		})
	})
})
