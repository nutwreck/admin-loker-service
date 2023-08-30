package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceLevelPekerjaan struct {
	levelPekerjaan entities.EntityLevelPekerjaan
}

func NewServiceLevelPekerjaan(levelPekerjaan entities.EntityLevelPekerjaan) *serviceLevelPekerjaan {
	return &serviceLevelPekerjaan{levelPekerjaan: levelPekerjaan}
}

/**
* ===========================================
* Service Create New Level Pekerjaan Teritory
*============================================
 */

func (s *serviceLevelPekerjaan) EntityCreate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan schemes.SchemeLevelPekerjaan
	levelPekerjaan.Name = input.Name

	res, err := s.levelPekerjaan.EntityCreate(&levelPekerjaan)
	return res, err
}

/**
* ============================================
* Service Results All Level Pekerjaan Teritory
*=============================================
 */

func (s *serviceLevelPekerjaan) EntityResults(input *schemes.SchemeLevelPekerjaan) (*[]models.ModelLevelPekerjaan, int64, schemes.SchemeDatabaseError) {
	var levelPekerjaan schemes.SchemeLevelPekerjaan
	levelPekerjaan.Page = input.Page
	levelPekerjaan.PerPage = input.PerPage
	levelPekerjaan.Name = input.Name
	levelPekerjaan.ID = input.ID

	res, totalData, err := s.levelPekerjaan.EntityResults(&levelPekerjaan)
	return res, totalData, err
}

/**
* =============================================
* Service Result Level Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceLevelPekerjaan) EntityResult(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan schemes.SchemeLevelPekerjaan
	levelPekerjaan.ID = input.ID

	res, err := s.levelPekerjaan.EntityResult(&levelPekerjaan)
	return res, err
}

/**
* =============================================
* Service Delete Level Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceLevelPekerjaan) EntityDelete(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan schemes.SchemeLevelPekerjaan
	levelPekerjaan.ID = input.ID

	res, err := s.levelPekerjaan.EntityDelete(&levelPekerjaan)
	return res, err
}

/**
* =============================================
* Service Update Level Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceLevelPekerjaan) EntityUpdate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan schemes.SchemeLevelPekerjaan
	levelPekerjaan.ID = input.ID
	levelPekerjaan.Name = input.Name
	levelPekerjaan.Active = input.Active

	res, err := s.levelPekerjaan.EntityUpdate(&levelPekerjaan)
	return res, err
}
