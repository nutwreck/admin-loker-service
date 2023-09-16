package repositories

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/schemes"
	"gorm.io/gorm"
)

type repositoryProvinsi struct {
	db *gorm.DB
}

func NewRepositoryProvinsi(db *gorm.DB) *repositoryProvinsi {
	return &repositoryProvinsi{db: db}
}

/**
* =======================================
* Repository Create New Provinsi Teritory
*========================================
 */

func (r *repositoryProvinsi) EntityCreate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi
	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkData := db.Debug().First(&provinsi, "name = ?", provinsi.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &provinsi, <-err
	}

	addData := db.Debug().Create(&provinsi).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* ========================================
* Repository Results All Provinsi Teritory
*=========================================
 */

func (r *repositoryProvinsi) EntityResults(input *schemes.SchemeProvinsi) (*[]schemes.SchemeGetDataProvinsi, int64, schemes.SchemeDatabaseError) {
	var (
		provinsi        []models.ModelProvinsi
		result          []schemes.SchemeGetDataProvinsi
		countData       schemes.SchemeCountData
		args            []interface{}
		totalData       int64
		sortData        string = "provinsi.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(provinsi.*) AS count_data
		FROM model_provinsis AS provinsi
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			negara.code_negara AS code_negara,
			negara.name AS name_negara,
			provinsi.code_provinsi AS code_provinsi,
			provinsi.name AS name_provinsi
		FROM model_provinsis AS provinsi
	`

	queryAdditional = `
		JOIN model_negaras AS negara ON provinsi.parent_code_negara = negara.code_negara
	`

	queryAdditional += ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND provinsi.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.code_negara = ?`
		args = append(args, input.ParentCodeNegara)
	}

	if input.CodeProvinsi != constants.EMPTY_VALUE {
		queryAdditional += ` AND provinsi.code_provinsi = ?`
		args = append(args, input.CodeProvinsi)
	}

	if input.NameNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameNegara)+"%")
	}

	if input.Search != constants.EMPTY_VALUE {
		queryAdditional += ` AND (negara.name LIKE ? OR provinsi.name LIKE ?)`
		args = append(args,
			"%"+strings.ToUpper(input.Search)+"%",
			"%"+strings.ToUpper(input.Search)+"%")
	}

	//Eksekusi query ambil jumlah data tanpa limit
	db.Raw(queryCountData+queryAdditional, args...).Scan(&countData)

	queryAdditional += ` ORDER BY ` + sortData

	if input.Page != constants.EMPTY_NUMBER || input.PerPage != constants.EMPTY_NUMBER {
		queryAdditional += ` LIMIT ?`
		args = append(args, int(input.PerPage))

		queryAdditional += ` OFFSET ?`
		args = append(args, offset)
	}

	getDatas := db.Raw(queryData+queryAdditional, args...).Scan(&result)

	if getDatas.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
		return &result, totalData, <-err
	}

	// Menghitung total data yang diambil
	totalData = countData.CountData

	err <- schemes.SchemeDatabaseError{}
	return &result, totalData, <-err
}

/**
* =========================================
* Repository Result Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityResult(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &provinsi, <-err
	}

	db.Debug().
		Preload("Negara").
		First(&provinsi)

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* =========================================
* Repository Delete Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityDelete(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &provinsi, <-err
	}

	deleteData := db.Debug().Delete(&provinsi)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}

/**
* =========================================
* Repository Update Provinsi By ID Teritory
*==========================================
 */

func (r *repositoryProvinsi) EntityUpdate(input *schemes.SchemeProvinsi) (*models.ModelProvinsi, schemes.SchemeDatabaseError) {
	var provinsi models.ModelProvinsi
	provinsi.CodeProvinsi = input.CodeProvinsi

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&provinsi)

	checkId := db.Debug().First(&provinsi)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &provinsi, <-err
	}

	provinsi.ParentCodeNegara = input.ParentCodeNegara
	provinsi.Name = input.Name

	updateData := db.Debug().Updates(&provinsi)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &provinsi, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &provinsi, <-err
}
