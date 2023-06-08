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

}

func (b *testBankingSuite) TestBanking_Withdraw() {

}

func (b *testBankingSuite) TestBanking_TestTable() {
	for _, t := range []struct {
		testName      string
		depositAmount bank.THB
		want          bank.THB
		actual        bank.Banking
	}{
		{
			testName: "init",
			want:     bank.THB(0),
			actual:   bank.Banking{},
		},
		{
			testName:      "deposit - 1000",
			depositAmount: bank.THB(1000),
			want:          bank.THB(1000),
			actual:        bank.Banking{},
		},
	} {
		b.Run(t.testName, func() {
			t.actual.Deposit(t.depositAmount)
			b.Equal(t.want, t.actual.Balance())
		})
	}
}
