package dto

type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required,max=200"`
	Email    string `json:"email" validate:"required,max=100,email"`
	NoHp     string `json:"no_hp" validate:"max=16,number"`
	Username string `json:"username" validate:"required,min=5"`
	RoleId   uint   `json:"role_id" validate:"required"`
}
