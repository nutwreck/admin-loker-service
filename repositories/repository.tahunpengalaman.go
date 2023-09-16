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

type repositoryTahunPengalaman struct {
	db *gorm.DB
}

func NewRepositoryTahunPengalaman(db *gorm.DB) *repositoryTahunPengalaman {
	return &repositoryTahunPengalaman{db: db}
}

/**
* ===============================================
* Repository Create New Tahun Pengalaman Teritory
*================================================
 */

func (r *repositoryTahunPengalaman) EntityCreate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman models.ModelTahunPengalaman
	tahunPengalaman.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tahunPengalaman)

	checkTahunPengalamanName := db.Debug().First(&tahunPengalaman, "name = ?", tahunPengalaman.Name)

	if checkTahunPengalamanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &tahunPengalaman, <-err
	}

	addTahunPengalaman := db.Debug().Create(&tahunPengalaman).Commit()

	if addTahunPengalaman.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &tahunPengalaman, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tahunPengalaman, <-err
}

/**
* ================================================
* Repository Results All Tahun Pengalaman Teritory
*=================================================
 */

func (r *repositoryTahunPengalaman) EntityResults(input *schemes.SchemeTahunPengalaman) (*[]models.ModelTahunPengalaman, int64, schemes.SchemeDatabaseError) {
	var (
		tahunPengalaman []models.ModelTahunPengalaman
		totalData       int64
		sort            string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tahunPengalaman)

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

	checkTahunPengalaman := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&tahunPengalaman)

	if checkTahunPengalaman.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &tahunPengalaman, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelTahunPengalaman{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &tahunPengalaman, totalData, <-err
}

/**
* =================================================
* Repository Result Tahun Pengalaman By ID Teritory
*==================================================
 */

func (r *repositoryTahunPengalaman) EntityResult(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman models.ModelTahunPengalaman
	tahunPengalaman.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tahunPengalaman)

	checkTahunPengalamanId := db.Debug().First(&tahunPengalaman)

	if checkTahunPengalamanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &tahunPengalaman, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tahunPengalaman, <-err
}

/**
* =================================================
* Repository Delete Tahun Pengalaman By ID Teritory
*==================================================
 */

func (r *repositoryTahunPengalaman) EntityDelete(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman models.ModelTahunPengalaman
	tahunPengalaman.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tahunPengalaman)

	checkTahunPengalamanId := db.Debug().First(&tahunPengalaman)

	if checkTahunPengalamanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &tahunPengalaman, <-err
	}

	deleteTahunPengalaman := db.Debug().Delete(&tahunPengalaman)

	if deleteTahunPengalaman.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &tahunPengalaman, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tahunPengalaman, <-err
}

/**
* =================================================
* Repository Update Tahun Pengalaman By ID Teritory
*==================================================
 */

func (r *repositoryTahunPengalaman) EntityUpdate(input *schemes.SchemeTahunPengalaman) (*models.ModelTahunPengalaman, schemes.SchemeDatabaseError) {
	var tahunPengalaman models.ModelTahunPengalaman
	tahunPengalaman.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&tahunPengalaman)

	checkTahunPengalamanId := db.Debug().First(&tahunPengalaman)

	if checkTahunPengalamanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &tahunPengalaman, <-err
	}

	tahunPengalaman.Name = input.Name
	tahunPengalaman.Active = input.Active

	updateTahunPengalaman := db.Debug().Updates(&tahunPengalaman)

	if updateTahunPengalaman.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &tahunPengalaman, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &tahunPengalaman, <-err
}
