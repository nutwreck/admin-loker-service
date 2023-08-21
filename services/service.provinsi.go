package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceProvinsi struct {
	provinsi entities.EntityProvinsi
}

func NewServiceProvinsi(provinsi entities.EntityProvinsi) *serviceProvinsi {
	return &serviceProvinsi{provinsi: provinsi}
}

/**
* ====================================
* Service Create New Provinsi Teritory
*=====================================
 */

func (s *serviceProvinsi) EntityCreate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi schemes.SchemeProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi
	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	res, err := s.provinsi.EntityCreate(&provinsi)
	return res, err
}

/**
* =====================================
* Service Results All Provinsi Teritory
*======================================
 */

func (s *serviceProvinsi) EntityResults(input *schemes.SchemeProvinsi) (*[]models.ModelProvinsi, int64, schemes.SchemeDatabaseError) {
	var provinsi schemes.SchemeProvinsi
	provinsi.Page = input.Page
	provinsi.PerPage = input.PerPage
	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	res, totalData, err := s.provinsi.EntityResults(&provinsi)
	return res, totalData, err
}

/**
* ======================================
* Service Result Provinsi By ID Teritory
*=======================================
 */

func (s *serviceProvinsi) EntityResult(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi schemes.SchemeProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	res, err := s.provinsi.EntityResult(&provinsi)
	return res, err
}

/**
* ======================================
* Service Delete Provinsi By ID Teritory
*=======================================
 */

func (s *serviceProvinsi) EntityDelete(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi schemes.SchemeProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	res, err := s.provinsi.EntityDelete(&provinsi)
	return res, err
}

/**
* ======================================
* Service Update Provinsi By ID Teritory
*=======================================
 */

func (s *serviceProvinsi) EntityUpdate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi schemes.SchemeProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi
	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	res, err := s.provinsi.EntityUpdate(&provinsi)
	return res, err
}
