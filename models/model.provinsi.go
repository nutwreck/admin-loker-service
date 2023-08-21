package models

type ModelProvinsi struct {
	CodeProvinsi     string      `json:"code_provinsi" gorm:"primary_key; unique"`
	ParentCodeNegara string      `json:"parent_code_negara" gorm:"type:varchar; not null"`
	Name             string      `json:"name" gorm:"type:varchar; not null"`
	Negara           ModelNegara `gorm:"foreignkey:ParentCodeNegara;references:CodeNegara"`
}
