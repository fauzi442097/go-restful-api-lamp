package models

import "time"

type User struct {
	ID             uint    `gorm:"column:id;primaryKey"`
	FullName       string  `gorm:"type:varchar(150)"`
	UserCode       *string `gorm:"type:varchar(5)"`
	BranchId       uint
	RoleId         uint
	Password       string `gorm:"type:varchar(255)"`
	IsActive       *bool
	CompanyId      *uint
	UserCrtId      uint
	CreatedAt      time.Time
	UserUpdId      *uint
	UpdatedAt      *time.Time
	Username       string  `gorm:"type:varchar(50),unique"`
	Email          string  `gorm:"type:varchar(100),unique"`
	PhoneNo        *string `gorm:"type:varchar(20)"`
	ProductOwnerId *uint
}

func (User) TableName() string {
	return "master_user"
}
