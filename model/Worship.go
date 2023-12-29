package model

type Worship struct {
	Id            int    `json:"id"`
	Date          string `json:"date"`
	WorshipTypeId int    `json:"worshipTypeId"`
}
