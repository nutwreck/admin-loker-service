package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceJenisPerusahaan struct {
	jenisPerusahaan entities.EntityJenisPerusahaan
}

func NewServiceJenisPerusahaan(jenisPerusahaan entities.EntityJenisPerusahaan) *serviceJenisPerusahaan {
	return &serviceJenisPerusahaan{jenisPerusahaan: jenisPerusahaan}
}

/**
* ============================================
* Service Create New Jenis Perusahaan Teritory
*=============================================
 */

func (s *serviceJenisPerusahaan) EntityCreate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan schemes.SchemeJenisPerusahaan
	jenisPerusahaan.Name = input.Name

	res, err := s.jenisPerusahaan.EntityCreate(&jenisPerusahaan)
	return res, err
}

/**
* =============================================
* Service Results All Jenis Perusahaan Teritory
*==============================================
 */

func (s *serviceJenisPerusahaan) EntityResults(input *schemes.SchemeJenisPerusahaan) (*[]models.ModelJenisPerusahaan, int64, schemes.SchemeDatabaseError) {
	var jenisPerusahaan schemes.SchemeJenisPerusahaan
	jenisPerusahaan.Sort = input.Sort
	jenisPerusahaan.Page = input.Page
	jenisPerusahaan.PerPage = input.PerPage
	jenisPerusahaan.Name = input.Name
	jenisPerusahaan.ID = input.ID

	res, totalData, err := s.jenisPerusahaan.EntityResults(&jenisPerusahaan)
	return res, totalData, err
}

/**
* ==============================================
* Service Result Jenis Perusahaan By ID Teritory
*===============================================
 */

func (s *serviceJenisPerusahaan) EntityResult(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan schemes.SchemeJenisPerusahaan
	jenisPerusahaan.ID = input.ID

	res, err := s.jenisPerusahaan.EntityResult(&jenisPerusahaan)
	return res, err
}

/**
* ==============================================
* Service Delete Jenis Perusahaan By ID Teritory
*===============================================
 */

func (s *serviceJenisPerusahaan) EntityDelete(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan schemes.SchemeJenisPerusahaan
	jenisPerusahaan.ID = input.ID

	res, err := s.jenisPerusahaan.EntityDelete(&jenisPerusahaan)
	return res, err
}

/**
* ==============================================
* Service Update Jenis Perusahaan By ID Teritory
*===============================================
 */

func (s *serviceJenisPerusahaan) EntityUpdate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan schemes.SchemeJenisPerusahaan
	jenisPerusahaan.ID = input.ID
	jenisPerusahaan.Name = input.Name
	jenisPerusahaan.Active = input.Active

	res, err := s.jenisPerusahaan.EntityUpdate(&jenisPerusahaan)
	return res, err
}
