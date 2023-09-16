package services

import (
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type serviceNegara struct {
	negara entities.EntityNegara
}

func NewServiceNegara(negara entities.EntityNegara) *serviceNegara {
	return &serviceNegara{negara: negara}
}

/**
* ==================================
* Service Create New Negara Teritory
*===================================
 */

func (s *serviceNegara) EntityCreate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara schemes.SchemeNegara
	negara.CodeNegara = input.CodeNegara
	negara.ParentCode = input.ParentCode
	negara.Name = input.Name

	res, err := s.negara.EntityCreate(&negara)
	return res, err
}

/**
* ===================================
* Service Results All Negara Teritory
*====================================
 */

func (s *serviceNegara) EntityResults(input *schemes.SchemeNegara) (*[]models.ModelNegara, int64, schemes.SchemeDatabaseError) {
	var negara schemes.SchemeNegara
	negara.SortBy = input.SortBy
	negara.OrderBy = input.OrderBy
	negara.Page = input.Page
	negara.PerPage = input.PerPage
	negara.Name = input.Name

	res, totalData, err := s.negara.EntityResults(&negara)
	return res, totalData, err
}

/**
* ====================================
* Service Result Negara By ID Teritory
*=====================================
 */

func (s *serviceNegara) EntityResult(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara schemes.SchemeNegara
	negara.CodeNegara = input.CodeNegara

	res, err := s.negara.EntityResult(&negara)
	return res, err
}

/**
* ====================================
* Service Delete Negara By ID Teritory
*=====================================
 */

func (s *serviceNegara) EntityDelete(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara schemes.SchemeNegara
	negara.CodeNegara = input.CodeNegara

	res, err := s.negara.EntityDelete(&negara)
	return res, err
}

/**
* ====================================
* Service Update Negara By ID Teritory
*=====================================
 */

func (s *serviceNegara) EntityUpdate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara schemes.SchemeNegara
	negara.CodeNegara = input.CodeNegara
	negara.ParentCode = input.ParentCode
	negara.Name = input.Name

	res, err := s.negara.EntityUpdate(&negara)
	return res, err
}
