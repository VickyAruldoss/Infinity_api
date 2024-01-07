package model

import constants "github.com/infinity-api/constant"

type Member struct {
	Id              int                     `json:"id"`
	FamilyNo        int                     `json:"familyNo" validate:"required"`
	SubscriptionNo  int                     `json:"subscriptionNo" validate:"required"`
	Title           string                  `json:"title" validate:"required"`
	FirstName       *string                 `json:"firstName" validate:"required"`
	LastName        *string                 `json:"lastName" validate:"required"`
	FullName        string                  `json:"fullName" validate:"required"`
	NameEng         *string                 `json:"nameEng" validate:"required"`
	BaptisedNameEng *string                 `json:"baptisedNameEng,omitempty"`
	DateOfBirth     string                  `json:"dateOfBirth,omitempty"`
	Age             int                     `json:"age,omitempty"`
	IsBaptised      bool                    `json:"isBaptised"`
	IsFamilyLeader  bool                    `json:"isFamilyLeader"`
	Gender          constants.Gender        `json:"gender" validate:"required"`
	MaritalStatus   constants.MaritalStatus `json:"maritalStatus" validate:"required"`
	Address         Address                 `json:"address"`
	FamilyMembers   []Member                `json:"familyMembers,omitempty"`
}

type Address struct {
	DoorNo   int    `json:"doorNo,omitempty"`
	Street   string `json:"street"`
	Town     string `json:"location" validate:"required"`
	District string `json:"district"`
	PinCode  int    `json:"pinCode"`
}

func (member *Member) GetFullName() string {
	fullName := member.Title + ". " + validString(member.FirstName) + validString(member.LastName)
	return fullName
}

func validString(str *string) string {
	var s string
	if str != nil {
		s = *str
	} else {
		s = ""
	}
	return s
}

func (member *Member) GetTitle() string {
	var title string
	if member.Gender == constants.Male && member.MaritalStatus == constants.Married {
		title = constants.MR_T
	}
	if member.Gender == constants.Male && member.MaritalStatus == constants.Single {
		title = constants.MASTER_T
	}
	if member.Gender == constants.Female && (member.MaritalStatus == constants.Married || member.MaritalStatus == constants.Divorced || member.MaritalStatus == constants.Widow) {
		title = constants.MRS_T
	}
	if member.Gender == constants.Female && member.MaritalStatus == constants.Single {
		title = constants.SELVI
	}
	return title
}
