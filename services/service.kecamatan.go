package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceKecamatan struct {
	kecamatan entities.EntityKecamatan
}

func NewServiceKecamatan(kecamatan entities.EntityKecamatan) *serviceKecamatan {
	return &serviceKecamatan{kecamatan: kecamatan}
}

/**
* =====================================
* Service Create New Kecamatan Teritory
*======================================
 */

func (s *serviceKecamatan) EntityCreate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan schemes.SchemeKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan
	kecamatan.ParentCodeKabupaten = input.ParentCodeKabupaten
	kecamatan.Name = input.Name

	res, err := s.kecamatan.EntityCreate(&kecamatan)
	return res, err
}

/**
* ======================================
* Service Results All Kecamatan Teritory
*=======================================
 */

func (s *serviceKecamatan) EntityResults(input *schemes.SchemeKecamatan) (*[]schemes.SchemeGetDataKecamatan, int64, schemes.SchemeDatabaseError) {
	var kecamatan schemes.SchemeKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan
	kecamatan.Sort = input.Sort
	kecamatan.Search = input.Search
	kecamatan.CodeNegara = input.CodeNegara
	kecamatan.NameNegara = input.NameNegara
	kecamatan.CodeProvinsi = input.CodeProvinsi
	kecamatan.NameProvinsi = input.NameProvinsi
	kecamatan.NameKabupaten = input.NameKabupaten
	kecamatan.Page = input.Page
	kecamatan.PerPage = input.PerPage
	kecamatan.ParentCodeKabupaten = input.ParentCodeKabupaten
	kecamatan.Name = input.Name

	res, totalData, err := s.kecamatan.EntityResults(&kecamatan)
	return res, totalData, err
}

/**
* =======================================
* Service Result Kecamatan By ID Teritory
*========================================
 */

func (s *serviceKecamatan) EntityResult(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan schemes.SchemeKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan

	res, err := s.kecamatan.EntityResult(&kecamatan)
	return res, err
}

/**
* =======================================
* Service Delete Kecamatan By ID Teritory
*========================================
 */

func (s *serviceKecamatan) EntityDelete(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan schemes.SchemeKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan

	res, err := s.kecamatan.EntityDelete(&kecamatan)
	return res, err
}

/**
* =======================================
* Service Update Kecamatan By ID Teritory
*========================================
 */

func (s *serviceKecamatan) EntityUpdate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan schemes.SchemeKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan
	kecamatan.ParentCodeKabupaten = input.ParentCodeKabupaten
	kecamatan.Name = input.Name

	res, err := s.kecamatan.EntityUpdate(&kecamatan)
	return res, err
}
