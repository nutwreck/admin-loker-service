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

type repositoryKeahlian struct {
	db *gorm.DB
}

func NewRepositoryKeahlian(db *gorm.DB) *repositoryKeahlian {
	return &repositoryKeahlian{db: db}
}

/**
* =======================================
* Repository Create New Keahlian Teritory
*========================================
 */

func (r *repositoryKeahlian) EntityCreate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian models.ModelKeahlian
	keahlian.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&keahlian)

	checkKeahlianName := db.Debug().First(&keahlian, "name = ?", keahlian.Name)

	if checkKeahlianName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &keahlian, <-err
	}

	addKeahlian := db.Debug().Create(&keahlian).Commit()

	if addKeahlian.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &keahlian, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &keahlian, <-err
}

/**
* ========================================
* Repository Results All Keahlian Teritory
*=========================================
 */

func (r *repositoryKeahlian) EntityResults(input *schemes.SchemeKeahlian) (*[]models.ModelKeahlian, int64, schemes.SchemeDatabaseError) {
	var (
		keahlian  []models.ModelKeahlian
		totalData int64
		sort      string = configs.SortByDefault + " " + configs.OrderByDefault
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&keahlian)

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

	checkKeahlian := db.Debug().Order(sort).Offset(offset).Limit(int(input.PerPage)).Find(&keahlian)

	if checkKeahlian.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &keahlian, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelKeahlian{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &keahlian, totalData, <-err
}

/**
* =========================================
* Repository Result Keahlian By ID Teritory
*==========================================
 */

func (r *repositoryKeahlian) EntityResult(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian models.ModelKeahlian
	keahlian.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&keahlian)

	checkKeahlianId := db.Debug().First(&keahlian)

	if checkKeahlianId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &keahlian, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &keahlian, <-err
}

/**
* =========================================
* Repository Delete Keahlian By ID Teritory
*==========================================
 */

func (r *repositoryKeahlian) EntityDelete(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian models.ModelKeahlian
	keahlian.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&keahlian)

	checkKeahlianId := db.Debug().First(&keahlian)

	if checkKeahlianId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &keahlian, <-err
	}

	deleteKeahlian := db.Debug().Delete(&keahlian)

	if deleteKeahlian.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &keahlian, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &keahlian, <-err
}

/**
* =========================================
* Repository Update Keahlian By ID Teritory
*==========================================
 */

func (r *repositoryKeahlian) EntityUpdate(input *schemes.SchemeKeahlian) (*models.ModelKeahlian, schemes.SchemeDatabaseError) {
	var keahlian models.ModelKeahlian
	keahlian.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&keahlian)

	checkKeahlianId := db.Debug().First(&keahlian)

	if checkKeahlianId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &keahlian, <-err
	}

	keahlian.Name = input.Name
	keahlian.Active = input.Active

	updateKeahlian := db.Debug().Updates(&keahlian)

	if updateKeahlian.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &keahlian, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &keahlian, <-err
}
