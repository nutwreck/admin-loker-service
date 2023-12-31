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

type repositoryJenisPerusahaan struct {
	db *gorm.DB
}

func NewRepositoryJenisPerusahaan(db *gorm.DB) *repositoryJenisPerusahaan {
	return &repositoryJenisPerusahaan{db: db}
}

/**
* ===============================================
* Repository Create New Jenis Perusahaan Teritory
*================================================
 */

func (r *repositoryJenisPerusahaan) EntityCreate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan models.ModelJenisPerusahaan
	jenisPerusahaan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&jenisPerusahaan)

	checkJenisPerusahaanName := db.Debug().First(&jenisPerusahaan, "name = ?", jenisPerusahaan.Name)

	if checkJenisPerusahaanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &jenisPerusahaan, <-err
	}

	addJenisPerusahaan := db.Debug().Create(&jenisPerusahaan).Commit()

	if addJenisPerusahaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &jenisPerusahaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &jenisPerusahaan, <-err
}

/**
* ================================================
* Repository Results All Jenis Perusahaan Teritory
*=================================================
 */

func (r *repositoryJenisPerusahaan) EntityResults(input *schemes.SchemeJenisPerusahaan) (*[]models.ModelJenisPerusahaan, int64, schemes.SchemeDatabaseError) {
	var (
		jenisPerusahaan []models.ModelJenisPerusahaan
		totalData       int64
		sort            string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&jenisPerusahaan)

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

	checkJenisPerusahaan := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&jenisPerusahaan)

	if checkJenisPerusahaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &jenisPerusahaan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelJenisPerusahaan{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &jenisPerusahaan, totalData, <-err
}

/**
* =================================================
* Repository Result Jenis Perusahaan By ID Teritory
*==================================================
 */

func (r *repositoryJenisPerusahaan) EntityResult(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan models.ModelJenisPerusahaan
	jenisPerusahaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&jenisPerusahaan)

	checkJenisPerusahaanId := db.Debug().First(&jenisPerusahaan)

	if checkJenisPerusahaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &jenisPerusahaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &jenisPerusahaan, <-err
}

/**
* =================================================
* Repository Delete Jenis Perusahaan By ID Teritory
*==================================================
 */

func (r *repositoryJenisPerusahaan) EntityDelete(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan models.ModelJenisPerusahaan
	jenisPerusahaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&jenisPerusahaan)

	checkJenisPerusahaanId := db.Debug().First(&jenisPerusahaan)

	if checkJenisPerusahaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &jenisPerusahaan, <-err
	}

	deleteJenisPerusahaan := db.Debug().Delete(&jenisPerusahaan)

	if deleteJenisPerusahaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &jenisPerusahaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &jenisPerusahaan, <-err
}

/**
* =================================================
* Repository Update Jenis Perusahaan By ID Teritory
*==================================================
 */

func (r *repositoryJenisPerusahaan) EntityUpdate(input *schemes.SchemeJenisPerusahaan) (*models.ModelJenisPerusahaan, schemes.SchemeDatabaseError) {
	var jenisPerusahaan models.ModelJenisPerusahaan
	jenisPerusahaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&jenisPerusahaan)

	checkJenisPerusahaanId := db.Debug().First(&jenisPerusahaan)

	if checkJenisPerusahaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &jenisPerusahaan, <-err
	}

	jenisPerusahaan.Name = input.Name
	jenisPerusahaan.Active = input.Active

	updateJenisPerusahaan := db.Debug().Updates(&jenisPerusahaan)

	if updateJenisPerusahaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &jenisPerusahaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &jenisPerusahaan, <-err
}
