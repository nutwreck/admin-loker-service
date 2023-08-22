package repositories

import (
	"net/http"
	"strings"

	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryKecamatan struct {
	db *gorm.DB
}

func NewRepositoryKecamatan(db *gorm.DB) *repositoryKecamatan {
	return &repositoryKecamatan{db: db}
}

/**
* ========================================
* Repository Create New Kecamatan Teritory
*=========================================
 */

func (r *repositoryKecamatan) EntityCreate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan models.ModelKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan
	kecamatan.ParentCodeKabupaten = input.ParentCodeKabupaten
	kecamatan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	checkData := db.Debug().First(&kecamatan, "name = ?", kecamatan.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &kecamatan, <-err
	}

	addData := db.Debug().Create(&kecamatan).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &kecamatan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kecamatan, <-err
}

/**
* =========================================
* Repository Results All Kecamatan Teritory
*==========================================
 */

func (r *repositoryKecamatan) EntityResults(input *schemes.SchemeKecamatan) (*[]models.ModelKecamatan, int64, schemes.SchemeDatabaseError) {
	var (
		kecamatan []models.ModelKecamatan
		totalData int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	if input.Name != "" {
		db = db.Where("model_kecamatans.name LIKE ?", "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeKabupaten != "" {
		db = db.Where("model_kecamatans.parent_code_kabupaten", input.ParentCodeKabupaten)
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkData := db.Debug().Order("model_kecamatans.name ASC").Offset(offset).Limit(int(input.PerPage)).Find(&kecamatan)

	if checkData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &kecamatan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelKecamatan{}).Count(&totalData)

	if input.Page == constants.EMPTY_NUMBER || input.PerPage == constants.EMPTY_NUMBER { //Off Pagination
		db.Debug().
			Preload("Kabupaten").
			Preload("Kabupaten.Provinsi").
			Preload("Kabupaten.Provinsi.Negara").
			Find(&kecamatan)
	} else {
		db.Debug().
			Offset(offset).
			Limit(int(input.PerPage)).
			Preload("Kabupaten").
			Preload("Kabupaten.Provinsi").
			Preload("Kabupaten.Provinsi.Negara").
			Find(&kecamatan)
	}

	err <- schemes.SchemeDatabaseError{}
	return &kecamatan, totalData, <-err
}

/**
* ==========================================
* Repository Result Kecamatan By ID Teritory
*===========================================
 */

func (r *repositoryKecamatan) EntityResult(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan models.ModelKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	checkId := db.Debug().First(&kecamatan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kecamatan, <-err
	}

	db.Debug().
		Preload("Kabupaten").
		Preload("Kabupaten.Provinsi").
		Preload("Provinsi.Negara").
		First(&kecamatan)

	err <- schemes.SchemeDatabaseError{}
	return &kecamatan, <-err
}

/**
* ==========================================
* Repository Delete Kecamatan By ID Teritory
*===========================================
 */

func (r *repositoryKecamatan) EntityDelete(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan models.ModelKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	checkId := db.Debug().First(&kecamatan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kecamatan, <-err
	}

	deleteData := db.Debug().Delete(&kecamatan)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kecamatan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kecamatan, <-err
}

/**
* ==========================================
* Repository Update Kecamatan By ID Teritory
*===========================================
 */

func (r *repositoryKecamatan) EntityUpdate(input *schemes.SchemeKecamatan) (*models.ModelKecamatan, schemes.SchemeDatabaseError) {
	var kecamatan models.ModelKecamatan
	kecamatan.CodeKecamatan = input.CodeKecamatan

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	checkId := db.Debug().First(&kecamatan)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kecamatan, <-err
	}

	kecamatan.ParentCodeKabupaten = input.ParentCodeKabupaten
	kecamatan.Name = input.Name

	updateData := db.Debug().Updates(&kecamatan)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kecamatan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kecamatan, <-err
}
