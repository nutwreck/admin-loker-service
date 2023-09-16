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

type repositoryKabupaten struct {
	db *gorm.DB
}

func NewRepositoryKabupaten(db *gorm.DB) *repositoryKabupaten {
	return &repositoryKabupaten{db: db}
}

/**
* ========================================
* Repository Create New Kabupaten Teritory
*=========================================
 */

func (r *repositoryKabupaten) EntityCreate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten
	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkData := db.Debug().First(&kabupaten, "name = ?", kabupaten.Name)

	if checkData.RowsAffected > 0 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &kabupaten, <-err
	}

	addData := db.Debug().Create(&kabupaten).Commit()

	if addData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* =========================================
* Repository Results All Kabupaten Teritory
*==========================================
 */

func (r *repositoryKabupaten) EntityResults(input *schemes.SchemeKabupaten) (*[]schemes.SchemeGetDataKabupaten, int64, schemes.SchemeDatabaseError) {
	var (
		kabupaten       []models.ModelKabupaten
		result          []schemes.SchemeGetDataKabupaten
		countData       schemes.SchemeCountData
		args            []interface{}
		totalData       int64
		sortData        string = "kabupaten.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(kabupaten.*) AS count_data
		FROM model_kabupatens AS kabupaten
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			negara.code_negara AS code_negara,
			negara.name AS name_negara,
			provinsi.code_provinsi AS code_provinsi,
			provinsi.name AS name_provinsi,
			kabupaten.code_kabupaten AS code_kabupaten,
			kabupaten.name AS name_kabupaten
		FROM model_kabupatens AS kabupaten
	`

	queryAdditional = `
		JOIN model_provinsis AS provinsi ON kabupaten.parent_code_provinsi = provinsi.code_provinsi
		JOIN model_negaras AS negara ON provinsi.parent_code_negara = negara.code_negara
	`

	queryAdditional += ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeProvinsi != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.parent_code_provinsi = ?`
		args = append(args, input.ParentCodeProvinsi)
	}

	if input.CodeKabupaten != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.code_kabupaten = ?`
		args = append(args, input.CodeKabupaten)
	}

	if input.CodeNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.code_negara = ?`
		args = append(args, input.CodeNegara)
	}

	if input.NameNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameNegara)+"%")
	}

	if input.NameProvinsi != constants.EMPTY_VALUE {
		queryAdditional += ` AND provinsi.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameProvinsi)+"%")
	}

	if input.Search != constants.EMPTY_VALUE {
		queryAdditional += ` AND (negara.name LIKE ? OR provinsi.name LIKE ? OR kabupaten.name LIKE ?)`
		args = append(args,
			"%"+strings.ToUpper(input.Search)+"%",
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
* ==========================================
* Repository Result Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityResult(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &kabupaten, <-err
	}

	db.Debug().
		Preload("Provinsi").
		Preload("Provinsi.Negara").
		First(&kabupaten)

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* ==========================================
* Repository Delete Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityDelete(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &kabupaten, <-err
	}

	deleteData := db.Debug().Delete(&kabupaten)

	if deleteData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}

/**
* ==========================================
* Repository Update Kabupaten By ID Teritory
*===========================================
 */

func (r *repositoryKabupaten) EntityUpdate(input *schemes.SchemeKabupaten) (*models.ModelKabupaten, schemes.SchemeDatabaseError) {
	var kabupaten models.ModelKabupaten
	kabupaten.CodeKabupaten = input.CodeKabupaten

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kabupaten)

	checkId := db.Debug().First(&kabupaten)

	if checkId.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &kabupaten, <-err
	}

	kabupaten.ParentCodeProvinsi = input.ParentCodeProvinsi
	kabupaten.Name = input.Name

	updateData := db.Debug().Updates(&kabupaten)

	if updateData.RowsAffected < 1 {
		err <- schemes.SchemeDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &kabupaten, <-err
	}

	err <- schemes.SchemeDatabaseError{}
	return &kabupaten, <-err
}
