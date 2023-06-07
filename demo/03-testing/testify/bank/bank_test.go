package bank_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testBankingSuite struct {
	suite.Suite
}

func TestBanking(t *testing.T) {
	suite.Run(t, &testBankingSuite{})
}

func (b *testBankingSuite) SetupTest() {}

func (b *testBankingSuite) TearDownTest() {}
