package dtos

type ReqUser struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	PhoneNumber     string `json:"phone_number" binding:"required"`
	InputedPassword string `json:"password" binding:"required"`
}

type ResUser struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	CreationDate string `json:"creation_date"`
}
