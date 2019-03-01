package _unittests

import (
	. "7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The ATM Machine", func() {
	Context("When getting the initial balance", func() {
		It("Returns a balance of `0.00.", func() {
			a := Account{}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			Expect(account.GetBalance()).To(Equal(0.00))
		})
	})

	Context("When depositing `5.00` dollars", func() {
		It("Returns a balance of `5.00.", func() {
			a := Account{}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			account.Deposit(5.00)
			Expect(account.GetBalance()).To(Equal(5.00))
		})
	})

	Context("When depositing `10.00` dollars", func() {
		It("Returns a balance of `10.00.", func() {
			a := Account{}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			account.Deposit(10.00)
			Expect(account.GetBalance()).To(Equal(10.00))
		})
	})

	Context("When withdrawing `1.00` dollar from a balance of `10.00` dollars", func() {
		It("Returns a balance of `9.00`.", func() {
			a := Account{10}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			account.Withdraw(1.00)
			Expect(account.GetBalance()).To(Equal(9.00))
		})
	})

	Context("When withdrawing `2.00` dollar from a balance of `8.00` dollars", func() {
		It("Returns a balance of `9.00`.", func() {
			a := Account{8}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			account.Withdraw(2.00)
			Expect(account.GetBalance()).To(Equal(6.00))
		})
	})

	Context("When withdrawing more money than in is in the account", func() {
		It("Withdraws no money and returns the appropriate error.", func() {
			a := Account{5}
			account, err := a.GetAccount()
			Expect(err).To(BeNil())
			err = account.Withdraw(6.00)
			Expect(err).ToNot(BeNil())
			Expect(account.GetBalance()).To(Equal(5.00))
		})
	})
})
