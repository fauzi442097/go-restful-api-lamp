package models

import "time"

type Customer struct {
	ID            uint    `gorm:"column:id;primaryKey"`
	FullName      string  `gorm:"type:varchar(200)"`
	KTPNo         string  `gorm:"varchar(50)"`
	Address       *string `gorm:"type:varchar(255)"`
	IdWilayah     int
	HpNo          *string `gorm:"type:varchar(50)"`
	OfficeName    *string `gorm:"type:varchar(255)"`
	JobTittle     *string `gorm:"type:varchar(100)"`
	VisitNo       *uint
	IsActive      bool
	CompanyId     uint
	BranchId      uint
	UserCrtId     uint
	CreatedAt     time.Time
	UserUpdId     *uint
	UpdatedAt     *time.Time
	IsMapCustomer *bool
	VillageId     *uint
	OfficeId      *uint
	AccNo         *string `gorm:"type:varchar(30)"`
	MerchantId    *uint
	CIF           *string `gorm:"type:varchar(15)"`
	Sallary       *float32
	NIP           *string `gorm:"type:varchar(255)"`
	IsPotensi     *bool
	OfficeCode    *string `gorm:"type:varchar(20)"`
	UploadTypeId  *uint
}

// TableName overrides the table name used by Customer to `master_customer`
func (Customer) TableName() string {
	return "master_customer"
}
