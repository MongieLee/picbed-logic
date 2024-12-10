package models

type RegisterDto struct {
	LoginDto
	RePassword string `json:"rePassword"  binding:"required,eqfield=Password"`
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
