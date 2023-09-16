package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceKategoriPekerjaan struct {
	kategoriPekerjaan entities.EntityKategoriPekerjaan
}

func NewServiceKategoriPekerjaan(kategoriPekerjaan entities.EntityKategoriPekerjaan) *serviceKategoriPekerjaan {
	return &serviceKategoriPekerjaan{kategoriPekerjaan: kategoriPekerjaan}
}

/**
* ==============================================
* Service Create New Kategori Pekerjaan Teritory
*===============================================
 */

func (s *serviceKategoriPekerjaan) EntityCreate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan schemes.SchemeKategoriPekerjaan
	kategoriPekerjaan.Name = input.Name

	res, err := s.kategoriPekerjaan.EntityCreate(&kategoriPekerjaan)
	return res, err
}

/**
* ===============================================
* Service Results All Kategori Pekerjaan Teritory
*================================================
 */

func (s *serviceKategoriPekerjaan) EntityResults(input *schemes.SchemeKategoriPekerjaan) (*[]models.ModelKategoriPekerjaan, int64, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan schemes.SchemeKategoriPekerjaan
	kategoriPekerjaan.Sort = input.Sort
	kategoriPekerjaan.Page = input.Page
	kategoriPekerjaan.PerPage = input.PerPage
	kategoriPekerjaan.Name = input.Name
	kategoriPekerjaan.ID = input.ID

	res, totalData, err := s.kategoriPekerjaan.EntityResults(&kategoriPekerjaan)
	return res, totalData, err
}

/**
* ================================================
* Service Result Kategori Pekerjaan By ID Teritory
*=================================================
 */

func (s *serviceKategoriPekerjaan) EntityResult(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan schemes.SchemeKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID

	res, err := s.kategoriPekerjaan.EntityResult(&kategoriPekerjaan)
	return res, err
}

/**
* ================================================
* Service Delete Kategori Pekerjaan By ID Teritory
*=================================================
 */

func (s *serviceKategoriPekerjaan) EntityDelete(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan schemes.SchemeKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID

	res, err := s.kategoriPekerjaan.EntityDelete(&kategoriPekerjaan)
	return res, err
}

/**
* ================================================
* Service Update Kategori Pekerjaan By ID Teritory
*=================================================
 */

func (s *serviceKategoriPekerjaan) EntityUpdate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan schemes.SchemeKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID
	kategoriPekerjaan.Name = input.Name
	kategoriPekerjaan.Active = input.Active

	res, err := s.kategoriPekerjaan.EntityUpdate(&kategoriPekerjaan)
	return res, err
}
