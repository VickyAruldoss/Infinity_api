package dto

type MemberDTO struct {
	Id              *int    `db:"id" validate:"required"`
	FamilyNo        *int    `db:"family_no"`
	SubscriptionNo  *int    `db:"member_id"`
	FirstName       *string `db:"first_name"`
	LastName        *string `db:"last_name,omitempty"`
	BaptisedName    *string `db:"baptised_name,omitempty"`
	EngLastName     *string `db:"eng_last_Name,omitempty"`
	EngBaptisedName *string `db:"eng_baptised_name,omitempty"`
	Gender          *string `db:"gender" validate:"required"`
	MaritalStatus   *string `db:"marital_status" validate:"required"`
	Street          string  `db:"street"`
	Town            string  `db:"town"`
	City            string  `db:"city"`
	PostalCode      int     `db:"postal_code"`
}
