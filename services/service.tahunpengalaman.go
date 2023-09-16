package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceTahunPengalaman struct {
	tahunPengalaman entities.EntityTahunPengalaman
}

func NewServiceTahunPengalaman(tahunPengalaman entities.EntityTahunPengalaman) *serviceTahunPengalaman {
	return &serviceTahunPengalaman{tahunPengalaman: tahunPengalaman}
}

/**
* ============================================
* Service Create New Tahun Pengalaman Teritory
*=============================================
 */

func (s *serviceTahunPengalaman) EntityCreate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman schemes.SchemeTahunPengalaman
	tahunPengalaman.Name = input.Name

	res, err := s.tahunPengalaman.EntityCreate(&tahunPengalaman)
	return res, err
}

/**
* =============================================
* Service Results All Tahun Pengalaman Teritory
*==============================================
 */

func (s *serviceTahunPengalaman) EntityResults(input *schemes.SchemeTahunPengalaman) (*[]models.ModelTahunPengalaman, int64, schemes.SchemeDatabaseError) {
	var tahunPengalaman schemes.SchemeTahunPengalaman
	tahunPengalaman.Sort = input.Sort
	tahunPengalaman.Page = input.Page
	tahunPengalaman.PerPage = input.PerPage
	tahunPengalaman.Name = input.Name
	tahunPengalaman.ID = input.ID

	res, totalData, err := s.tahunPengalaman.EntityResults(&tahunPengalaman)
	return res, totalData, err
}

/**
* ==============================================
* Service Result Tahun Pengalaman By ID Teritory
*===============================================
 */

func (s *serviceTahunPengalaman) EntityResult(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman schemes.SchemeTahunPengalaman
	tahunPengalaman.ID = input.ID

	res, err := s.tahunPengalaman.EntityResult(&tahunPengalaman)
	return res, err
}

/**
* ==============================================
* Service Delete Tahun Pengalaman By ID Teritory
*===============================================
 */

func (s *serviceTahunPengalaman) EntityDelete(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman schemes.SchemeTahunPengalaman
	tahunPengalaman.ID = input.ID

	res, err := s.tahunPengalaman.EntityDelete(&tahunPengalaman)
	return res, err
}

/**
* ==============================================
* Service Update Tahun Pengalaman By ID Teritory
*===============================================
 */

func (s *serviceTahunPengalaman) EntityUpdate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman schemes.SchemeTahunPengalaman
	tahunPengalaman.ID = input.ID
	tahunPengalaman.Name = input.Name
	tahunPengalaman.Active = input.Active

	res, err := s.tahunPengalaman.EntityUpdate(&tahunPengalaman)
	return res, err
}
