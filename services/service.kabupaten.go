package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceKabupaten struct {
	kabupaten entities.EntityKabupaten
}

func NewServiceKabupaten(kabupaten entities.EntityKabupaten) *serviceKabupaten {
	return &serviceKabupaten{kabupaten: kabupaten}
}

/**
* =====================================
* Service Create New Kabupaten Teritory
*======================================
 */

func (s *serviceKabupaten) EntityCreate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten schemes.SchemeKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten
	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	res, err := s.kabupaten.EntityCreate(&kabupaten)
	return res, err
}

/**
* ======================================
* Service Results All Kabupaten Teritory
*=======================================
 */

func (s *serviceKabupaten) EntityResults(input *schemes.SchemeKabupaten) (*[]schemes.SchemeGetDataKabupaten, int64, schemes.SchemeDatabaseError) {
	var kabupaten schemes.SchemeKabupaten
	kabupaten.Sort = input.Sort
	kabupaten.Search = input.Search
	kabupaten.CodeNegara = input.CodeNegara
	kabupaten.NameNegara = input.NameNegara
	kabupaten.NameProvinsi = input.NameProvinsi
	kabupaten.Page = input.Page
	kabupaten.PerPage = input.PerPage
	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name
	kabupaten.CodeKabupaten = input.CodeKabupaten

	res, totalData, err := s.kabupaten.EntityResults(&kabupaten)
	return res, totalData, err
}

/**
* =======================================
* Service Result Kabupaten By ID Teritory
*========================================
 */

func (s *serviceKabupaten) EntityResult(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten schemes.SchemeKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	res, err := s.kabupaten.EntityResult(&kabupaten)
	return res, err
}

/**
* =======================================
* Service Delete Kabupaten By ID Teritory
*========================================
 */

func (s *serviceKabupaten) EntityDelete(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten schemes.SchemeKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	res, err := s.kabupaten.EntityDelete(&kabupaten)
	return res, err
}

/**
* =======================================
* Service Update Kabupaten By ID Teritory
*========================================
 */

func (s *serviceKabupaten) EntityUpdate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten schemes.SchemeKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten
	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	res, err := s.kabupaten.EntityUpdate(&kabupaten)
	return res, err
}
