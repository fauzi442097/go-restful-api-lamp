package dto

import "time"

type UserResponse struct {
	ID             uint       `json:"id"`
	FullName       string     `json:"full_name"`
	UserCode       *string    `json:"user_code"`
	BranchId       uint       `json:"branch_id"`
	RoleId         uint       `json:"role_id"`
	Password       string     `json:"password"`
	IsActive       *bool      `json:"is_active"`
	CompanyId      *uint      `json:"company_id"`
	UserCrtId      uint       `json:"user_crt_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UserUpdId      *uint      `json:"user_upd_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Username       string     `json:"username"`
	Email          string     `json:"email"`
	PhoneNo        *string    `json:"phone_no"`
	ProductOwnerId *uint      `json:"product_owner_id"`
}
