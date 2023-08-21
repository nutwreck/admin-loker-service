package repositories

import (
	"net/http"
	"strings"

	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryKabupaten struct {
	db *gorm.DB
}

func NewRepositoryKabupaten(db *gorm.DB) *repositoryKabupaten {
	return &repositoryKabupaten{db: db}
}

/**
* ========================================
* Repository Create New Kabupaten Teritory
*=========================================
 */

func (r *repositoryKabupaten) EntityCreate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten
	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkData := db.Debug().First(&kabupaten, "name = ?", kabupaten.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &kabupaten, <-err
	}

	addData := db.Debug().Create(&kabupaten).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* =========================================
* Repository Results All Kabupaten Teritory
*==========================================
 */

func (r *repositoryKabupaten) EntityResults(input *schemes.SchemeKabupaten) (*[]models.ModelKabupaten, int64, schemes.SchemeDatabaseError) {
	var (
		kabupaten []models.ModelKabupaten
		totalData int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	if input.Name != "" {
		db = db.Where("model_kabupatens.name LIKE ?", "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeProvinsi != "" {
		db = db.Where("model_kabupatens.parent_code_provinsi", input.ParentCodeProvinsi)
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkData := db.Debug().Order("model_kabupatens.name ASC").Offset(offset).Limit(int(input.PerPage)).Find(&kabupaten)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &kabupaten, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelKabupaten{}).Count(&totalData)

	if input.Page == constants.EMPTY_NUMBER || input.PerPage == constants.EMPTY_NUMBER { //Off Pagination
		db.Debug().
			Preload("Provinsi").
			Preload("Provinsi.Negara").
			Find(&kabupaten)
	} else {
		db.Debug().
			Offset(offset).
			Limit(int(input.PerPage)).
			Preload("Provinsi").
			Preload("Provinsi.Negara").
			Find(&kabupaten)
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, totalData, <-err
}

/**
* ==========================================
* Repository Result Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityResult(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kabupaten, <-err
	}

	db.Debug().
		Preload("Provinsi").
		Preload("Provinsi.Negara").
		First(&kabupaten)

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* ==========================================
* Repository Delete Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityDelete(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kabupaten, <-err
	}

	deleteData := db.Debug().Delete(&kabupaten)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* ==========================================
* Repository Update Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityUpdate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kabupaten, <-err
	}

	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	updateData := db.Debug().Updates(&kabupaten)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}
