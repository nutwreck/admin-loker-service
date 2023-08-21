package models

type ModelKelurahan struct {
	CodeKelurahan       string         `json:"code_kelurahan" gorm:"primary_key; unique"`
	ParentCodeKecamatan string         `json:"parent_code_kecamatan" gorm:"type:varchar; not null"`
	Name                string         `json:"name" gorm:"type:varchar; not null"`
	Kecamatan           ModelKecamatan `gorm:"foreignkey:ParentCodeKecamatan;references:CodeKecamatan"`
}
