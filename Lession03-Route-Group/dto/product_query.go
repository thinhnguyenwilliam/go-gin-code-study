package dto

type UserUUIDQuery struct {
	UUID string `uri:"uuid" binding:"uuid4"`
}

type UserQuery struct {
	ID int `uri:"id" binding:"gt=0"`
}

type ProductQuery struct {
	Search string `form:"search" validate:"required,min=3,max=50,alphanumspace"`
	Limit  int    `form:"limit" validate:"omitempty,gt=0"`
}
