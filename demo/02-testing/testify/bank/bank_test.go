package bank_test

import (
	"example/testify/bank"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testBankingSuite struct {
	suite.Suite
	underTest bank.Banking
}

func TestBanking(t *testing.T) {
	suite.Run(t, &testBankingSuite{})
}

func (b *testBankingSuite) SetupTest() {
	b.underTest = bank.Banking{}
}

func (b *testBankingSuite) TearDownTest() {}

func (b *testBankingSuite) TestBanking_Deposit() {
	want := bank.THB(1000)
	b.underTest.Deposit(want)
	b.Equal(want, b.underTest.Balance())
}

func (b *testBankingSuite) TestBanking_Withdraw() {
	b.Run("withdraw success", func() {
		withDrawAmount := bank.THB(500)
		b.underTest.Deposit(1500)

		b.underTest.Withdraw(withDrawAmount)

		b.Equal(bank.THB(1000), b.underTest.Balance())
	})
	b.Run("withdraw fail", func() {
		b.underTest = bank.Banking{}
		withDrawAmount := bank.THB(500)
		b.underTest.Deposit(300)

		err := b.underTest.Withdraw(withDrawAmount)

		b.Equal("Cannot withdraw, money not enough", err.Error())
	})
}

func (b *testBankingSuite) TestBanking_TestTable() {
	for _, t := range []struct {
		testName      string
		depositAmount bank.THB
		want          bank.THB
		actual        bank.THB
	}{
		{
			testName: "init",
			want:     bank.THB(0),
			actual:   bank.THB(0),
		},
		{
			testName:      "deposit - 300",
			depositAmount: bank.THB(300),
			want:          bank.THB(300),
			actual:        bank.THB(300),
		},
		{
			testName:      "deposit - 1000",
			depositAmount: bank.THB(1000),
			want:          bank.THB(1000),
			actual:        bank.THB(1000),
		},
	} {
		b.Run(t.testName, func() {
			b.underTest = bank.Banking{}
			b.underTest.Deposit(t.depositAmount)
			b.Equal(t.want, t.actual)
		})
	}
}
