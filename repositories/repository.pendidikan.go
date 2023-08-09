package repositories

import (
	"net/http"

	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryPendidikan struct {
	db *gorm.DB
}

func NewRepositoryPendidikan(db *gorm.DB) *repositoryPendidikan {
	return &repositoryPendidikan{db: db}
}

/**
* ==========================================
* Repository Create New Pendidikan Teritory
*===========================================
 */

func (r *repositoryPendidikan) EntityCreate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan models.ModelPendidikan
	pendidikan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&pendidikan)

	checkPendidikanName := db.Debug().First(&pendidikan, "name = ?", pendidikan.Name)

	if checkPendidikanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &pendidikan, <-err
	}

	addPendidikan := db.Debug().Create(&pendidikan).Commit()

	if addPendidikan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &pendidikan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &pendidikan, <-err
}

/**
* ==========================================
* Repository Results All Pendidikan Teritory
*===========================================
 */

func (r *repositoryPendidikan) EntityResults(input *schemes.SchemePendidikan) (*[]models.ModelPendidikan, int64, schemes.SchemeDatabaseError) {
	var (
		pendidikan []models.ModelPendidikan
		totalData  int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&pendidikan)

	if input.Name != "" {
		db = db.Where("name LIKE ?", "%"+input.Name+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkPendidikan := db.Debug().Order("created_at DESC").Offset(offset).Limit(int(input.PerPage)).Find(&pendidikan)

	if checkPendidikan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &pendidikan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelPendidikan{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &pendidikan, totalData, <-err
}

/**
* ==========================================
* Repository Result Pendidikan By ID Teritory
*===========================================
 */

func (r *repositoryPendidikan) EntityResult(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan models.ModelPendidikan
	pendidikan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&pendidikan)

	checkPendidikanId := db.Debug().First(&pendidikan)

	if checkPendidikanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &pendidikan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &pendidikan, <-err
}

/**
* ==========================================
* Repository Delete Pendidikan By ID Teritory
*===========================================
 */

func (r *repositoryPendidikan) EntityDelete(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan models.ModelPendidikan
	pendidikan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&pendidikan)

	checkPendidikanId := db.Debug().First(&pendidikan)

	if checkPendidikanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &pendidikan, <-err
	}

	deletePendidikan := db.Debug().Delete(&pendidikan)

	if deletePendidikan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &pendidikan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &pendidikan, <-err
}

/**
* ==========================================
* Repository Update Pendidikan By ID Teritory
*===========================================
 */

func (r *repositoryPendidikan) EntityUpdate(input *schemes.SchemePendidikan) (*models.ModelPendidikan, schemes.SchemeDatabaseError) {
	var pendidikan models.ModelPendidikan
	pendidikan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&pendidikan)

	checkPendidikanId := db.Debug().First(&pendidikan)

	if checkPendidikanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &pendidikan, <-err
	}

	pendidikan.Name = input.Name
	pendidikan.Active = input.Active

	updatePendidikan := db.Debug().Updates(&pendidikan)

	if updatePendidikan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &pendidikan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &pendidikan, <-err
}
