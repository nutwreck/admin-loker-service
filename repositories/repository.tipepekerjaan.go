package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/nutwreck/admin-loker-service/configs"
	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryTipePekerjaan struct {
	db *gorm.DB
}

func NewRepositoryTipePekerjaan(db *gorm.DB) *repositoryTipePekerjaan {
	return &repositoryTipePekerjaan{db: db}
}

/**
* ==============================================
* Repository Create New Tipe Pekerjaan Teritory
*===============================================
 */

func (r *repositoryTipePekerjaan) EntityCreate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan models.ModelTipePekerjaan
	tipePekerjaan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tipePekerjaan)

	checkTipePekerjaanName := db.Debug().First(&tipePekerjaan, "name = ?", tipePekerjaan.Name)

	if checkTipePekerjaanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &tipePekerjaan, <-err
	}

	addTipePekerjaan := db.Debug().Create(&tipePekerjaan).Commit()

	if addTipePekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &tipePekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tipePekerjaan, <-err
}

/**
* ===============================================
* Repository Results All Tipe Pekerjaan Teritory
*================================================
 */

func (r *repositoryTipePekerjaan) EntityResults(input *schemes.SchemeTipePekerjaan) (*[]models.ModelTipePekerjaan, int64, schemes.SchemeDatabaseError) {
	var (
		tipePekerjaan []models.ModelTipePekerjaan
		totalData     int64
		sort          string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tipePekerjaan)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sort = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	if input.Name != constants.EMPTY_VALUE {
		db = db.Where("name LIKE ?", "%"+input.Name+"%")
	}

	if input.ID != constants.EMPTY_VALUE {
		db = db.Where("id LIKE ?", "%"+input.ID+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkTipePekerjaan := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&tipePekerjaan)

	if checkTipePekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &tipePekerjaan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelTipePekerjaan{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &tipePekerjaan, totalData, <-err
}

/**
* ================================================
* Repository Result Tipe Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryTipePekerjaan) EntityResult(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan models.ModelTipePekerjaan
	tipePekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tipePekerjaan)

	checkTipePekerjaanId := db.Debug().First(&tipePekerjaan)

	if checkTipePekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &tipePekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tipePekerjaan, <-err
}

/**
* ================================================
* Repository Delete Tipe Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryTipePekerjaan) EntityDelete(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan models.ModelTipePekerjaan
	tipePekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tipePekerjaan)

	checkTipePekerjaanId := db.Debug().First(&tipePekerjaan)

	if checkTipePekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &tipePekerjaan, <-err
	}

	deleteTipePekerjaan := db.Debug().Delete(&tipePekerjaan)

	if deleteTipePekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &tipePekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tipePekerjaan, <-err
}

/**
* ================================================
* Repository Update Tipe Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryTipePekerjaan) EntityUpdate(input *schemes.SchemeTipePekerjaan) (*models.ModelTipePekerjaan, schemes.SchemeDatabaseError) {
	var tipePekerjaan models.ModelTipePekerjaan
	tipePekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tipePekerjaan)

	checkTipePekerjaanId := db.Debug().First(&tipePekerjaan)

	if checkTipePekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &tipePekerjaan, <-err
	}

	tipePekerjaan.Name = input.Name
	tipePekerjaan.Active = input.Active

	updateTipePekerjaan := db.Debug().Updates(&tipePekerjaan)

	if updateTipePekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &tipePekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tipePekerjaan, <-err
}
