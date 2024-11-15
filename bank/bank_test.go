package bank_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"hop-unit-testing/bank"
)

var _ = Describe("Bank", func() {
	var account bank.Account

	BeforeEach(func() {
		account.Balance = 100.0
	})
	AfterEach(func() {
		account.Balance = 0
	})

	Describe("Deposit", func() {
		It("should return an error if amount is less than or equal to zero", func() {
			err := account.Deposit(0)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("the amount cannot be less than equal zero"))
		})

		It("should return the appropriate amount of account balance after deposit", func() {
			err := account.Deposit(100.50)
			Expect(err).Should(BeNil())
			Expect(account.Balance).To(Equal(200.50))
		})
	})

	Describe("WithDraw", func() {
		It("should return an error if amount is less than or equal to zero", func() {
			err := account.Withdraw(0)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("the amount cannot be less than equal zero"))
		})

		It("should return an error if amount is greater than balance", func() {
			err := account.Withdraw(500)
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).To(Equal("insufficient funds"))
		})

		It("should return the appropriate amount of account balance after withdraw", func() {
			err := account.Withdraw(20.0)
			Expect(err).Should(BeNil())
			Expect(account.Balance).To(Equal(80.0))
		})
	})
})
