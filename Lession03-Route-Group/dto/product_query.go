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
	URL     string `json:"url" binding:"required,url,imgext"`
	AltText string `json:"alt_text" binding:"omitempty,max=100"`
}
type AvartarImage struct {
	URL string `json:"url" binding:"required,url,imgext"`
	Alt string `json:"alt" binding:"omitempty,max=100"`
}
type ProductInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}

type CreateProductRequest struct {
	ProductInfo map[string]ProductInfo `json:"product_info" binding:"required,dive"`
	Avartar     AvartarImage           `json:"avartar" binding:"required"`
	Tags        []string               `json:"tags" binding:"omitempty,dive,required,min=2,max=30"`
	Image       []ProductImage         `json:"image" binding:"required,dive"` // // ✅ Required & dive into each item
	Display     *bool                  `json:"display" binding:"omitempty"`   //By changing Display to a pointer (*bool), you can detect if it was omitted:
	Name        string                 `json:"name" binding:"required,min=3,max=100"`
	Description string                 `json:"description" binding:"omitempty,max=500"`
	Price       float64                `json:"price" binding:"required,gt=0,lte=100"`
	Stock       int                    `json:"stock" binding:"required,gte=0"`
	Email       string                 `json:"email" binding:"omitempty,email"`
	CreatedAt   string                 `json:"created_at,omitempty"` // return to client, but not accepted from client
}
