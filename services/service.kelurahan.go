package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceKelurahan struct {
	kelurahan entities.EntityKelurahan
}

func NewServiceKelurahan(kelurahan entities.EntityKelurahan) *serviceKelurahan {
	return &serviceKelurahan{kelurahan: kelurahan}
}

/**
* =====================================
* Service Create New Kelurahan Teritory
*======================================
 */

func (s *serviceKelurahan) EntityCreate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan schemes.SchemeKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan
	kelurahan.ParentCodeKecamatan = input.ParentCodeKecamatan
	kelurahan.Name = input.Name

	res, err := s.kelurahan.EntityCreate(&kelurahan)
	return res, err
}

/**
* ======================================
* Service Results All Kelurahan Teritory
*=======================================
 */

func (s *serviceKelurahan) EntityResults(input *schemes.SchemeKelurahan) (*[]schemes.SchemeGetDataKelurahan, int64, schemes.SchemeDatabaseError) {
	var kelurahan schemes.SchemeKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan
	kelurahan.Sort = input.Sort
	kelurahan.Search = input.Search
	kelurahan.CodeNegara = input.CodeNegara
	kelurahan.NameNegara = input.NameNegara
	kelurahan.CodeProvinsi = input.CodeProvinsi
	kelurahan.NameProvinsi = input.NameProvinsi
	kelurahan.CodeKabupaten = input.CodeKabupaten
	kelurahan.NameKabupaten = input.NameKabupaten
	kelurahan.NameKecamatan = input.NameKecamatan
	kelurahan.Page = input.Page
	kelurahan.PerPage = input.PerPage
	kelurahan.ParentCodeKecamatan = input.ParentCodeKecamatan
	kelurahan.Name = input.Name

	res, totalData, err := s.kelurahan.EntityResults(&kelurahan)
	return res, totalData, err
}

/**
* =======================================
* Service Result Kelurahan By ID Teritory
*========================================
 */

func (s *serviceKelurahan) EntityResult(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan schemes.SchemeKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan

	res, err := s.kelurahan.EntityResult(&kelurahan)
	return res, err
}

/**
* =======================================
* Service Delete Kelurahan By ID Teritory
*========================================
 */

func (s *serviceKelurahan) EntityDelete(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan schemes.SchemeKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan

	res, err := s.kelurahan.EntityDelete(&kelurahan)
	return res, err
}

/**
* =======================================
* Service Update Kelurahan By ID Teritory
*========================================
 */

func (s *serviceKelurahan) EntityUpdate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan schemes.SchemeKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan
	kelurahan.ParentCodeKecamatan = input.ParentCodeKecamatan
	kelurahan.Name = input.Name

	res, err := s.kelurahan.EntityUpdate(&kelurahan)
	return res, err
}
