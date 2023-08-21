package models

type ModelKecamatan struct {
	CodeKecamatan       string         `json:"code_kecamatan" gorm:"primary_key; unique"`
	ParentCodeKabupaten string         `json:"parent_code_kabupaten" gorm:"type:varchar; not null"`
	Name                string         `json:"name" gorm:"type:varchar; not null"`
	Kabupaten           ModelKabupaten `gorm:"foreignkey:ParentCodeKabupaten;references:CodeKabupaten"`
}
