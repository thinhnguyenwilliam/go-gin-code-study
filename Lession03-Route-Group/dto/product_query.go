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

type ProductImage struct {
	URL string `json:"url" binding:"required,url"`
}

type CreateProductRequest struct {
	ProductImage
	Display     *bool   `json:"display" binding:"omitempty"` //By changing Display to a pointer (*bool), you can detect if it was omitted:
	Name        string  `json:"name" binding:"required,min=3,max=100"`
	Description string  `json:"description" binding:"omitempty,max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,gte=0"`
	Email       string  `json:"email" binding:"omitempty,email"`
	CreatedAt   string  `json:"created_at,omitempty"` // return to client, but not accepted from client
}
