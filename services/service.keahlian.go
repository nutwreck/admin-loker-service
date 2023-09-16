package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceKeahlian struct {
	keahlian entities.EntityKeahlian
}

func NewServiceKeahlian(keahlian entities.EntityKeahlian) *serviceKeahlian {
	return &serviceKeahlian{keahlian: keahlian}
}

/**
* ====================================
* Service Create New Keahlian Teritory
*=====================================
 */

func (s *serviceKeahlian) EntityCreate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian schemes.SchemeKeahlian
	keahlian.Name = input.Name

	res, err := s.keahlian.EntityCreate(&keahlian)
	return res, err
}

/**
* =====================================
* Service Results All Keahlian Teritory
*======================================
 */

func (s *serviceKeahlian) EntityResults(input *schemes.SchemeKeahlian) (*[]models.ModelKeahlian, int64, schemes.SchemeDatabaseError) {
	var keahlian schemes.SchemeKeahlian
	keahlian.Sort = input.Sort
	keahlian.Page = input.Page
	keahlian.PerPage = input.PerPage
	keahlian.Name = input.Name
	keahlian.ID = input.ID

	res, totalData, err := s.keahlian.EntityResults(&keahlian)
	return res, totalData, err
}

/**
* ======================================
* Service Result Keahlian By ID Teritory
*=======================================
 */

func (s *serviceKeahlian) EntityResult(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian schemes.SchemeKeahlian
	keahlian.ID = input.ID

	res, err := s.keahlian.EntityResult(&keahlian)
	return res, err
}

/**
* ======================================
* Service Delete Keahlian By ID Teritory
*=======================================
 */

func (s *serviceKeahlian) EntityDelete(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian schemes.SchemeKeahlian
	keahlian.ID = input.ID

	res, err := s.keahlian.EntityDelete(&keahlian)
	return res, err
}

/**
* ======================================
* Service Update Keahlian By ID Teritory
*=======================================
 */

func (s *serviceKeahlian) EntityUpdate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian schemes.SchemeKeahlian
	keahlian.ID = input.ID
	keahlian.Name = input.Name
	keahlian.Active = input.Active

	res, err := s.keahlian.EntityUpdate(&keahlian)
	return res, err
}
