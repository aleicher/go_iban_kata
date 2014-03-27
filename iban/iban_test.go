package iban_test

import (
	. "github.com/aleicher/go_iban_kata/iban"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IbanValidator", func() {
	var iban *IbanValidator
	validIbanFixture := "DE89370400440532013000"
	validIbanFixtureChecksum := 1
	BeforeEach(func() {
		iban = new(IbanValidator)
	})

	Context("Validates a valid IBAN", func() {
		Describe("Validation 0: valid for valid IBAN :)", func() {
			It("does return true for a valid IBAN", func() {
				Expect(iban.IsValid(validIbanFixture)).Should(BeTrue())
			})
		})
	})

	Context("Validates the generic form of an IBAN", func() {
		Describe("Validation 1.1: maximum length", func() {
			tooLongIban := "12345678901234567890123456789012345"

			It("does return false for an IBAN with more than 34 characters", func() {
				Expect(iban.IsValid(tooLongIban)).Should(BeFalse())
			})
		})

		Describe("Validation 1.2: first two positions are characters", func() {
			noCharIban := "1234567890123456789012345678901234"
			oneCharIban := "X234567890123456789012345678901234"
			It("does return false for an IBAN which starts with numbers", func() {
				Expect(iban.IsValid(noCharIban)).Should(BeFalse())
			})
			It("does return false for an IBAN which starts with one character", func() {
				Expect(iban.IsValid(oneCharIban)).Should(BeFalse())
			})
		})

		Describe("Validation 1.3: positions 3 and 4 are numbers", func() {
			posThreeWrong := "XXX4567890123456789012345678901234"
			posFourWrong := "XX3X567890123456789012345678901234"
			posThreeAndFourWrong := "XXXX567890123456789012345678901234"
			It("does return false if position 3 is not a number", func() {
				Expect(iban.IsValid(posThreeWrong)).Should(BeFalse())
			})
			It("does return false if position 4 is not a number", func() {
				Expect(iban.IsValid(posFourWrong)).Should(BeFalse())
			})
			It("does return false if position 4 is not a number", func() {
				Expect(iban.IsValid(posThreeAndFourWrong)).Should(BeFalse())
			})
		})

		Describe("Validation 1.4: everything after position 4 is a number", func() {
			hasOneCharacterAfterPositionFour := "DE8937X400440532013000"
			hasTwoCharactersAfterPositionFour := "DE8937X40044053201X000"
			It("does return false if there is a character after position four", func() {
				Expect(iban.IsValid(hasOneCharacterAfterPositionFour)).Should(BeFalse())
			})
			It("does return false if there is more than one character after position four", func() {
				Expect(iban.IsValid(hasTwoCharactersAfterPositionFour)).Should(BeFalse())
			})

		})
	})
	Context("Validates the checksum", func() {
		Describe("Validation 2.1: validation of the checksum", func() {
			It("does return false for an IBAN with wrong checksum", func() {
				hasInvalidChecksum := "DE89370400440532013001"
				Expect(iban.IsValid(hasInvalidChecksum)).To(Equal(false))
			})
			It("does provide a function to calculate the checksum", func() {
				Expect(iban.Checksum(validIbanFixture)).To(Equal(validIbanFixtureChecksum))
			})
			PIt("does calculate the right checksum for ", func() {
				//Expect(iban.Checksum())
			})

		})
	})
})
