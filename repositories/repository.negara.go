package repositories

import (
	"net/http"
	"strings"

	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryNegara struct {
	db *gorm.DB
}

func NewRepositoryNegara(db *gorm.DB) *repositoryNegara {
	return &repositoryNegara{db: db}
}

/**
* =====================================
* Repository Create New Negara Teritory
*======================================
 */

func (r *repositoryNegara) EntityCreate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara models.ModelNegara
	negara.CodeNegara = input.CodeNegara
	negara.ParentCode = input.ParentCode
	negara.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&negara)

	checkData := db.Debug().First(&negara, "name = ?", negara.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &negara, <-err
	}

	addData := db.Debug().Create(&negara).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &negara, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &negara, <-err
}

/**
* =======================================
* Repository Results All Negara Teritory
*========================================
 */

func (r *repositoryNegara) EntityResults(input *schemes.SchemeNegara) (*[]models.ModelNegara, int64, schemes.SchemeDatabaseError) {
	var (
		negara    []models.ModelNegara
		totalData int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&negara)

	if input.Name != "" {
		db = db.Where("name LIKE ?", "%"+strings.ToUpper(input.Name)+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkData := db.Debug().Order("name ASC").Offset(offset).Limit(int(input.PerPage)).Find(&negara)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &negara, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelNegara{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &negara, totalData, <-err
}

/**
* =======================================
* Repository Result Negara By ID Teritory
*========================================
 */

func (r *repositoryNegara) EntityResult(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara models.ModelNegara
	negara.CodeNegara = input.CodeNegara

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&negara)

	checkId := db.Debug().First(&negara)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &negara, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &negara, <-err
}

/**
* =======================================
* Repository Delete Negara By ID Teritory
*========================================
 */

func (r *repositoryNegara) EntityDelete(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara models.ModelNegara
	negara.CodeNegara = input.CodeNegara

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&negara)

	checkId := db.Debug().First(&negara)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &negara, <-err
	}

	deleteData := db.Debug().Delete(&negara)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &negara, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &negara, <-err
}

/**
* =======================================
* Repository Update Negara By ID Teritory
*========================================
 */

func (r *repositoryNegara) EntityUpdate(input *schemes.SchemeNegara) (*models.ModelNegara, schemes.SchemeDatabaseError) {
	var negara models.ModelNegara
	negara.CodeNegara = input.CodeNegara

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&negara)

	checkId := db.Debug().First(&negara)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &negara, <-err
	}

	negara.ParentCode = input.ParentCode
	negara.Name = input.Name

	updateData := db.Debug().Updates(&negara)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &negara, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &negara, <-err
}
