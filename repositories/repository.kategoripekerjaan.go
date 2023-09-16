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

type repositoryKategoriPekerjaan struct {
	db *gorm.DB
}

func NewRepositoryKategoriPekerjaan(db *gorm.DB) *repositoryKategoriPekerjaan {
	return &repositoryKategoriPekerjaan{db: db}
}

/**
* =================================================
* Repository Create New Kategori Pekerjaan Teritory
*==================================================
 */

func (r *repositoryKategoriPekerjaan) EntityCreate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan models.ModelKategoriPekerjaan
	kategoriPekerjaan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kategoriPekerjaan)

	checkKategoriPekerjaanName := db.Debug().First(&kategoriPekerjaan, "name = ?", kategoriPekerjaan.Name)

	if checkKategoriPekerjaanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &kategoriPekerjaan, <-err
	}

	addKategoriPekerjaan := db.Debug().Create(&kategoriPekerjaan).Commit()

	if addKategoriPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &kategoriPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kategoriPekerjaan, <-err
}

/**
* ==================================================
* Repository Results All Kategori Pekerjaan Teritory
*===================================================
 */

func (r *repositoryKategoriPekerjaan) EntityResults(input *schemes.SchemeKategoriPekerjaan) (*[]models.ModelKategoriPekerjaan, int64, schemes.SchemeDatabaseError) {
	var (
		kategoriPekerjaan []models.ModelKategoriPekerjaan
		totalData         int64
		sort              string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kategoriPekerjaan)

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

	checkKategoriPerusahaan := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&kategoriPekerjaan)

	if checkKategoriPerusahaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &kategoriPekerjaan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelKategoriPekerjaan{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &kategoriPekerjaan, totalData, <-err
}

/**
* ===================================================
* Repository Result Kategori Pekerjaan By ID Teritory
*====================================================
 */

func (r *repositoryKategoriPekerjaan) EntityResult(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan models.ModelKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kategoriPekerjaan)

	checkKategoriPekerjaanId := db.Debug().First(&kategoriPekerjaan)

	if checkKategoriPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kategoriPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kategoriPekerjaan, <-err
}

/**
* ===================================================
* Repository Delete Kategori Pekerjaan By ID Teritory
*====================================================
 */

func (r *repositoryKategoriPekerjaan) EntityDelete(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan models.ModelKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kategoriPekerjaan)

	checkKategoriPekerjaanId := db.Debug().First(&kategoriPekerjaan)

	if checkKategoriPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kategoriPekerjaan, <-err
	}

	deleteKategoriPekerjaan := db.Debug().Delete(&kategoriPekerjaan)

	if deleteKategoriPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kategoriPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kategoriPekerjaan, <-err
}

/**
* ===================================================
* Repository Update Kategori Pekerjaan By ID Teritory
*====================================================
 */

func (r *repositoryKategoriPekerjaan) EntityUpdate(input *schemes.SchemeKategoriPekerjaan) (*models.ModelKategoriPekerjaan, schemes.SchemeDatabaseError) {
	var kategoriPekerjaan models.ModelKategoriPekerjaan
	kategoriPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kategoriPekerjaan)

	checkKategoriPekerjaanId := db.Debug().First(&kategoriPekerjaan)

	if checkKategoriPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kategoriPekerjaan, <-err
	}

	kategoriPekerjaan.Name = input.Name
	kategoriPekerjaan.Active = input.Active

	updatekategoriPekerjaan := db.Debug().Updates(&kategoriPekerjaan)

	if updatekategoriPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kategoriPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kategoriPekerjaan, <-err
}
