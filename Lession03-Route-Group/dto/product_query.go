package dto

type ProductQuery struct {
	Search string `form:"search" validate:"required,min=3,max=50,alphanumspace"`
	Limit  int    `form:"limit" validate:"omitempty,gt=0"`
	Email  string `form:"email" binding:"omitempty,email"`
	Date   string `form:"date" binding:"omitempty,datetime=2006-01-02"`
}

type ProductLangUri struct {
	Lang string `uri:"lang" binding:"required,oneof=php golang python"`
}
