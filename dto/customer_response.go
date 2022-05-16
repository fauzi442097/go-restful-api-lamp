package dto

import "time"

type CustomerResponse struct {
	ID            uint       `json:"id"`
	FullName      string     `json:"full_name"`
	KTPNo         string     `json:"ktp_no"`
	Address       *string    `json:"address"`
	IdWilayah     int        `json:"id_wilayah"`
	HpNo          *string    `json:"hp_no"`
	OfficeName    *string    `json:"office_name"`
	JobTittle     *string    `json:"job_tittle"`
	VisitNo       *uint      `json:"visit_no"`
	IsActive      bool       `json:"is_active"`
	CompanyId     uint       `json:"company_id"`
	BranchId      uint       `json:"branch_id"`
	UserCrtId     uint       `json:"user_crt_id"`
	CreatedAt     time.Time  `json:"created_at"`
	UserUpdId     *uint      `json:"user_upd_id"`
	UpdatedAt     *time.Time `json:"updated_at"`
	IsMapCustomer *bool      `json:"is_map_customer"`
	VillageId     *uint      `json:"village_id"`
	OfficeId      *uint      `json:"office_id"`
	AccNo         *string    `json:"acc_no"`
	MerchantId    *uint      `json:"merchant_id"`
	CIF           *string    `json:"cif"`
	Sallary       *float32   `json:"sallary"`
	NIP           *string    `json:"nip"`
	IsPotensi     *bool      `json:"is_potensi"`
	OfficeCode    *string    `json:"office_code"`
	UploadTypeId  *uint      `json:"upload_type_id"`
}
