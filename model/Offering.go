package model

type Offering struct {
	Id             int          `json:"id" validate:"required"`
	WorshipId      int          `json:"worshipId" validate:"required"`
	OfferingTypeId int          `json:"offeringTypeId" validate:"required"`
	Denomination   Denomination `json:"Denomination" validate:"required"`
}
