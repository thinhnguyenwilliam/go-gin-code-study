package dto

//Only combine into UserQuery if you're absolutely sure all fields exist in the same route.

type UserUUIDQuery struct {
	UUID string `uri:"uuid" binding:"uuid4"`
}

type UserQuery struct {
	ID int `uri:"id" binding:"gt=0"`
}

type UserSlugQuery struct {
	Slug string `uri:"slug" binding:"required,min=5,max=100,slug"`
}
