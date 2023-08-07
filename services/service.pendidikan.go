package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type servicePendidikan struct {
	pendidikan entities.EntityPendidikan
}

func NewServicePendidikan(pendidikan entities.EntityPendidikan) *servicePendidikan {
	return &servicePendidikan{pendidikan: pendidikan}
}

/**
* ==========================================
* Service Create New Pendidikan Teritory
*===========================================
 */

func (s *servicePendidikan) EntityCreate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan schemes.SchemePendidikan
	pendidikan.Name = input.Name

	res, err := s.pendidikan.EntityCreate(&pendidikan)
	return res, err
}

/**
* ==========================================
* Service Results All Pendidikan Teritory
*===========================================
 */

func (s *servicePendidikan) EntityResults() (*[]models.ModelPendidikan, schemes.SchemeDatabaseError) {
	res, err := s.pendidikan.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Pendidikan By ID Teritory
*===========================================
 */

func (s *servicePendidikan) EntityResult(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan schemes.SchemePendidikan
	pendidikan.ID = input.ID

	res, err := s.pendidikan.EntityResult(&pendidikan)
	return res, err
}

/**
* ==========================================
* Service Delete Pendidikan By ID Teritory
*===========================================
 */

func (s *servicePendidikan) EntityDelete(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan schemes.SchemePendidikan
	pendidikan.ID = input.ID

	res, err := s.pendidikan.EntityDelete(&pendidikan)
	return res, err
}

/**
* ==========================================
* Service Update Pendidikan By ID Teritory
*===========================================
 */

func (s *servicePendidikan) EntityUpdate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan schemes.SchemePendidikan
	pendidikan.ID = input.ID
	pendidikan.Name = input.Name
	pendidikan.Active = input.Active

	res, err := s.pendidikan.EntityUpdate(&pendidikan)
	return res, err
}
