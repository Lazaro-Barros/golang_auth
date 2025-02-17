package dtos

type ReqLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResLogin struct {
	Token string `json:"token"`
}
