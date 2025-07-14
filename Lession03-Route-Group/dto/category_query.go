// dto/category.go
package dto

type CreateCategoryRequest struct {
	Name        string `form:"name" binding:"required,min=3,max=100"`
	Description string `form:"description" binding:"omitempty,max=255"`
	ImageURL    string `form:"image_url" binding:"required,url,imgext"`
}

type UploadCategoryForm struct {
	Name        string `form:"name" binding:"required,min=3,max=100"`
	Description string `form:"description" binding:"omitempty,max=500"`
}

type UploadCategoryQuery struct {
	UserID string `form:"user_id" binding:"omitempty,uuid4"`
	Source string `form:"source" binding:"omitempty,oneof=admin merchant api"`
}
