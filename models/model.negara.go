package models

type ModelNegara struct {
	CodeNegara string `json:"code_negara" gorm:"primary_key; unique"`
	ParentCode string `json:"parent_code" gorm:"type:varchar; not null"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
}
