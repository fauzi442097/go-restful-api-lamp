package dto

type CustomerRequest struct {
	KTPNo      string  `validate:"required,len=16,number" json:"ktp_no"`
	FullName   string  `validate:"required,max=200" json:"full_name"`
	HpNo       *string `validate:"min=8,max=16,number" json:"hp_no"`
	Address    *string `validate:"required,max=255" json:"address"`
	OfficeName *string `validate:"required,max=255" json:"office_name"`
	JobTittle  *string `validate:"required,max=100" json:"job_title"`
	IdWilayah  int     `json:"id_wilayah"`
}
