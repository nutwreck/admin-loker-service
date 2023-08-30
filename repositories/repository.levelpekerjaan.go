package repositories

import (
	"net/http"

	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryLevelPekerjaan struct {
	db *gorm.DB
}

func NewRepositoryLevelPekerjaan(db *gorm.DB) *repositoryLevelPekerjaan {
	return &repositoryLevelPekerjaan{db: db}
}

/**
* ==============================================
* Repository Create New Level Pekerjaan Teritory
*===============================================
 */

func (r *repositoryLevelPekerjaan) EntityCreate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan models.ModelLevelPekerjaan
	levelPekerjaan.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&levelPekerjaan)

	checkLevelPekerjaanName := db.Debug().First(&levelPekerjaan, "name = ?", levelPekerjaan.Name)

	if checkLevelPekerjaanName.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &levelPekerjaan, <-err
	}

	addLevelPekerjaan := db.Debug().Create(&levelPekerjaan).Commit()

	if addLevelPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &levelPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &levelPekerjaan, <-err
}

/**
* ===============================================
* Repository Results All Level Pekerjaan Teritory
*================================================
 */

func (r *repositoryLevelPekerjaan) EntityResults(input *schemes.SchemeLevelPekerjaan) (*[]models.ModelLevelPekerjaan, int64, schemes.SchemeDatabaseError) {
	var (
		levelPekerjaan []models.ModelLevelPekerjaan
		totalData      int64
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&levelPekerjaan)

	if input.Name != "" {
		db = db.Where("name LIKE ?", "%"+input.Name+"%")
	}

	if input.ID != "" {
		db = db.Where("id LIKE ?", "%"+input.ID+"%")
	}

	offset := int((input.Page - 1) * input.PerPage)

	checkLevelPekerjaan := db.Debug().Order("created_at DESC").Offset(offset).Limit(int(input.PerPage)).Find(&levelPekerjaan)

	if checkLevelPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &levelPekerjaan, totalData, <-err
	}

	// Menghitung total data yang diambil
	db.Model(&models.ModelLevelPekerjaan{}).Count(&totalData)

	err <- schemes.SchemeDatabaseError{}
	return &levelPekerjaan, totalData, <-err
}

/**
* ================================================
* Repository Result Level Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryLevelPekerjaan) EntityResult(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan models.ModelLevelPekerjaan
	levelPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&levelPekerjaan)

	checkLevelPekerjaanId := db.Debug().First(&levelPekerjaan)

	if checkLevelPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &levelPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &levelPekerjaan, <-err
}

/**
* ================================================
* Repository Delete Level Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryLevelPekerjaan) EntityDelete(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan models.ModelLevelPekerjaan
	levelPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&levelPekerjaan)

	checkLevelPekerjaanId := db.Debug().First(&levelPekerjaan)

	if checkLevelPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &levelPekerjaan, <-err
	}

	deleteLevelPekerjaan := db.Debug().Delete(&levelPekerjaan)

	if deleteLevelPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &levelPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &levelPekerjaan, <-err
}

/**
* ================================================
* Repository Update Level Pekerjaan By ID Teritory
*=================================================
 */

func (r *repositoryLevelPekerjaan) EntityUpdate(input *schemes.SchemeLevelPekerjaan) (*models.ModelLevelPekerjaan, schemes.SchemeDatabaseError) {
	var levelPekerjaan models.ModelLevelPekerjaan
	levelPekerjaan.ID = input.ID

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&levelPekerjaan)

	checkLevelPekerjaanId := db.Debug().First(&levelPekerjaan)

	if checkLevelPekerjaanId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &levelPekerjaan, <-err
	}

	levelPekerjaan.Name = input.Name
	levelPekerjaan.Active = input.Active

	updateLevelPekerjaan := db.Debug().Updates(&levelPekerjaan)

	if updateLevelPekerjaan.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &levelPekerjaan, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &levelPekerjaan, <-err
}
