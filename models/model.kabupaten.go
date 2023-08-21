package models

type ModelKabupaten struct {
	CodeKabupaten      string        `json:"code_kabupaten" gorm:"primary_key; unique"`
	ParentCodeProvinsi string        `json:"parent_code_provinsi" gorm:"type:varchar; not null"`
	Name               string        `json:"name" gorm:"type:varchar; not null"`
	Provinsi           ModelProvinsi `gorm:"foreignkey:ParentCodeProvinsi;references:CodeProvinsi"`
}
