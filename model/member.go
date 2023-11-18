package model

type Member struct {
	Id      int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}
