package repositories

import (
	"net/http"
	"strings"

	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryProvinsi struct {
	db *gorm.DB
}

func NewRepositoryProvinsi(db *gorm.DB) *repositoryProvinsi {
	return &repositoryProvinsi{db: db}
}

/**
* =======================================
* Repository Create New Provinsi Teritory
*========================================
 */

func (r *repositoryProvinsi) EntityCreate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi
	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkData := db.Debug().First(&provinsi, "name = ?", provinsi.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &provinsi, <-err
	}

	addData := db.Debug().Create(&provinsi).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* ========================================
* Repository Results All Provinsi Teritory
*=========================================
 */

func (r *repositoryProvinsi) EntityResults(input *schemes.SchemeProvinsi) (*[]models.ModelProvinsi, int64, schemes.SchemeDatabaseError) {
	var (
		provinsi  []models.ModelProvinsi
		totalData int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	if input.Name != "" {
		db = db.Where("model_provinsis.name LIKE ?", "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeNegara != "" {
		db = db.Where("model_provinsis.parent_code_negara", input.ParentCodeNegara)
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkData := db.Debug().Order("model_provinsis.name ASC").Offset(offset).Limit(int(input.PerPage)).Find(&provinsi)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &provinsi, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelProvinsi{}).Count(&totalData)

	if input.Page == constants.EMPTY_NUMBER || input.PerPage == constants.EMPTY_NUMBER { //Off Pagination
		db.Debug().
			Preload("Negara").
			Find(&provinsi)
	} else {
		db.Debug().
			Offset(offset).
			Limit(int(input.PerPage)).
			Preload("Negara").
			Find(&provinsi)
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, totalData, <-err
}

/**
* =========================================
* Repository Result Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityResult(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &provinsi, <-err
	}

	db.Debug().
		Preload("Negara").
		First(&provinsi)

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* =========================================
* Repository Delete Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityDelete(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &provinsi, <-err
	}

	deleteData := db.Debug().Delete(&provinsi)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* =========================================
* Repository Update Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityUpdate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &provinsi, <-err
	}

	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	updateData := db.Debug().Updates(&provinsi)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}
