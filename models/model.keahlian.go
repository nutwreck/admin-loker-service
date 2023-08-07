package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-loker-service/constants"
	"gorm.io/gorm"
)

type ModelKeahlian struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar;  not null"`
	Active    *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelKeahlian) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelKeahlian) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
