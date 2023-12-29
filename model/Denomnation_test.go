package model

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DenominationTestSuite struct {
	suite.Suite
}

func TestDenominationTestSuiteSetup(t *testing.T) {
	suite.Run(t, new(DenominationTestSuite))
}

func (suite *DenominationTestSuite) TestDenomination_ShouldReturnTotal() {
	d := &Denomination{
		FiveHundred: 2,
		Hundred:     5,
		Fifty:       10,
		Twenty:      20,
		Ten:         70,
		Five:        10,
		Coins:       40,
	}
	expected := 3190
	actual := d.GetTotal()
	suite.Equal(expected, actual)
}
