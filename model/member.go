package model

import constants "github.com/infinity-api/constant"

type Member struct {
	Id             int              `json:"id"`
	FamilyNo       int              `json:"familyNo" validate:"required"`
	SubscriptionNo int              `json:"subscriptionNo" validate:"required"`
	Denotation     string           `json:"denotion" validate:"required"`
	Name           string           `json:"name" validate:"required"`
	NameEng        string           `json:"nameEng" validate:"required"`
	DateOfBirth    string           `json:"dateOfBirth"`
	Age            int              `json:"age" validate:"required"`
	IsBaptised     bool             `json:"isBaptised"`
	IsFamilyLeader bool             `json:"isFamilyLeader"`
	Gender         constants.Gender `json:"gender" validate:"required"`
	Address        Address          `json:"address"`
	FamilyMembers  []Member         `json:"familyMembers"`
}

type Address struct {
	DoorNo   int    `json:"doorNo"`
	Street   string `json:"street"`
	Town     string `json:"location" validate:"required"`
	District string `json:"district"`
	PinCode  int    `json:"pinCode"`
}

func (member *Member) GetFullName() string {
	return member.Denotation + "." + member.Name
}
