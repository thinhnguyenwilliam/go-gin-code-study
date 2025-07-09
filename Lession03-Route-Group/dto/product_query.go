package dto

type ProductQuery struct {
	Search string `form:"search" validate:"required,min=3,max=50,alphanumspace"`
	Limit  int    `form:"limit" validate:"omitempty,gt=0"`
}
