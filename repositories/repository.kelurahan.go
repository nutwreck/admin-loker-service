package repositories

import (
	"net/http"
	"strings"

	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryKelurahan struct {
	db *gorm.DB
}

func NewRepositoryKelurahan(db *gorm.DB) *repositoryKelurahan {
	return &repositoryKelurahan{db: db}
}

/**
* ========================================
* Repository Create New Kelurahan Teritory
*=========================================
 */

func (r *repositoryKelurahan) EntityCreate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan models.ModelKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan
	kelurahan.ParentCodeKecamatan = input.ParentCodeKecamatan
	kelurahan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	checkData := db.Debug().First(&kelurahan, "name = ?", kelurahan.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &kelurahan, <-err
	}

	addData := db.Debug().Create(&kelurahan).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &kelurahan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kelurahan, <-err
}

/**
* =========================================
* Repository Results All Kelurahan Teritory
*==========================================
 */

func (r *repositoryKelurahan) EntityResults(input *schemes.SchemeKelurahan) (*[]models.ModelKelurahan, int64, schemes.SchemeDatabaseError) {
	var (
		kelurahan []models.ModelKelurahan
		totalData int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	if input.Name != "" {
		db = db.Where("model_kelurahans.name LIKE ?", "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeKecamatan != "" {
		db = db.Where("model_kelurahans.parent_code_kecamatan", input.ParentCodeKecamatan)
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkData := db.Debug().Order("model_kelurahans.name ASC").Offset(offset).Limit(int(input.PerPage)).Find(&kelurahan)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &kelurahan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelKelurahan{}).Count(&totalData)

	if input.Page == constants.EMPTY_NUMBER || input.PerPage == constants.EMPTY_NUMBER { //Off Pagination
		db.Debug().
			Preload("Kecamatan").
			Preload("Kecamatan.Kabupaten").
			Preload("Kecamatan.Kabupaten.Provinsi").
			Preload("Kecamatan.Kabupaten.Provinsi.Negara").
			Find(&kelurahan)
	} else {
		db.Debug().
			Offset(offset).
			Limit(int(input.PerPage)).
			Preload("Kecamatan").
			Preload("Kecamatan.Kabupaten").
			Preload("Kecamatan.Kabupaten.Provinsi").
			Preload("Kecamatan.Kabupaten.Provinsi.Negara").
			Find(&kelurahan)
	}

	err <- schemes.SchemeDatabaseError{}
	return &kelurahan, totalData, <-err
}

/**
* ==========================================
* Repository Result Kelurahan By ID Teritory
*===========================================
 */

func (r *repositoryKelurahan) EntityResult(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan models.ModelKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	checkId := db.Debug().First(&kelurahan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kelurahan, <-err
	}

	db.Debug().
		Preload("Kecamatan").
		Preload("Kecamatan.Kabupaten").
		Preload("Kecamatan.Kabupaten.Provinsi").
		Preload("Kecamatan.Kabupaten.Provinsi.Negara").
		First(&kelurahan)

	err <- schemes.SchemeDatabaseError{}
	return &kelurahan, <-err
}

/**
* ==========================================
* Repository Delete Kelurahan By ID Teritory
*===========================================
 */

func (r *repositoryKelurahan) EntityDelete(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan models.ModelKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	checkId := db.Debug().First(&kelurahan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kelurahan, <-err
	}

	deleteData := db.Debug().Delete(&kelurahan)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kelurahan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kelurahan, <-err
}

/**
* ==========================================
* Repository Update Kelurahan By ID Teritory
*===========================================
 */

func (r *repositoryKelurahan) EntityUpdate(input *schemes.SchemeKelurahan) (*models.ModelKelurahan, schemes.SchemeDatabaseError) {
	var kelurahan models.ModelKelurahan
	kelurahan.CodeKelurahan = input.CodeKelurahan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	checkId := db.Debug().First(&kelurahan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kelurahan, <-err
	}

	kelurahan.ParentCodeKecamatan = input.ParentCodeKecamatan
	kelurahan.Name = input.Name

	updateData := db.Debug().Updates(&kelurahan)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kelurahan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kelurahan, <-err
}
