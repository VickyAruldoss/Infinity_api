package model

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type MemberTestSuite struct {
	suite.Suite
	Member
}

func TestMemberTestSuite(t *testing.T) {
	suite.Run(t, new(MemberTestSuite))
}

func (suite *MemberTestSuite) SetupTest() {
}

func (s *MemberTestSuite) TestMember_ShouldReturnFullNameWithDenotion() {
	member := &Member{
		Name:       "Vicky Aruldoss",
		Denotation: "Thiru",
	}
	expectedFullName := "Thiru.Vicky Aruldoss"
	actualFullName := member.GetFullName()
	s.Suite.Equal(expectedFullName, actualFullName)
}
