package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceTipePekerjaan struct {
	tipePekerjaan entities.EntityTipePekerjaan
}

func NewServiceTipePekerjaan(tipePekerjaan entities.EntityTipePekerjaan) *serviceTipePekerjaan {
	return &serviceTipePekerjaan{tipePekerjaan: tipePekerjaan}
}

/**
* ===========================================
* Service Create New Tipe Pekerjaan Teritory
*============================================
 */

func (s *serviceTipePekerjaan) EntityCreate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan schemes.SchemeTipePekerjaan
	tipePekerjaan.Name = input.Name

	res, err := s.tipePekerjaan.EntityCreate(&tipePekerjaan)
	return res, err
}

/**
* ============================================
* Service Results All Tipe Pekerjaan Teritory
*=============================================
 */

func (s *serviceTipePekerjaan) EntityResults(input *schemes.SchemeTipePekerjaan) (*[]models.ModelTipePekerjaan, int64, schemes.SchemeDatabaseError) {
	var tipePekerjaan schemes.SchemeTipePekerjaan
	tipePekerjaan.Sort = input.Sort
	tipePekerjaan.Page = input.Page
	tipePekerjaan.PerPage = input.PerPage
	tipePekerjaan.Name = input.Name
	tipePekerjaan.ID = input.ID

	res, totalData, err := s.tipePekerjaan.EntityResults(&tipePekerjaan)
	return res, totalData, err
}

/**
* =============================================
* Service Result Tipe Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceTipePekerjaan) EntityResult(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan schemes.SchemeTipePekerjaan
	tipePekerjaan.ID = input.ID

	res, err := s.tipePekerjaan.EntityResult(&tipePekerjaan)
	return res, err
}

/**
* =============================================
* Service Delete Tipe Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceTipePekerjaan) EntityDelete(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan schemes.SchemeTipePekerjaan
	tipePekerjaan.ID = input.ID

	res, err := s.tipePekerjaan.EntityDelete(&tipePekerjaan)
	return res, err
}

/**
* =============================================
* Service Update Tipe Pekerjaan By ID Teritory
*==============================================
 */

func (s *serviceTipePekerjaan) EntityUpdate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan schemes.SchemeTipePekerjaan
	tipePekerjaan.ID = input.ID
	tipePekerjaan.Name = input.Name
	tipePekerjaan.Active = input.Active

	res, err := s.tipePekerjaan.EntityUpdate(&tipePekerjaan)
	return res, err
}
