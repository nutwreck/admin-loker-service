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

func (r *repositoryKecamatan) EntityResults(input *schemes.SchemeKecamatan) (*[]schemes.SchemeGetDataKecamatan, int64, schemes.SchemeDatabaseError) {
	var (
		kecamatan       []models.ModelKecamatan
		result          []schemes.SchemeGetDataKecamatan
		countData       schemes.SchemeCountData
		args            []interface{}
		totalData       int64
		sortData        string = "kecamatan.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kecamatan)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(kecamatan.*) AS count_data
		FROM model_kecamatans AS kecamatan
	`

	//Untuk mengambil detail data
	queryData = `
		SELECT
			negara.code_negara AS code_negara,
			negara.name AS name_negara,
			provinsi.code_provinsi AS code_provinsi,
			provinsi.name AS name_provinsi,
			kabupaten.code_kabupaten AS code_kabupaten,
			kabupaten.name AS name_kabupaten,
			kecamatan.code_kecamatan AS code_kecamatan,
			kecamatan.name AS name_kecamatan
		FROM model_kecamatans AS kecamatan
	`

	queryAdditional = `
		JOIN model_kabupatens AS kabupaten ON kecamatan.parent_code_kabupaten = kabupaten.code_kabupaten
		JOIN model_provinsis AS provinsi ON kabupaten.parent_code_provinsi = provinsi.code_provinsi
		JOIN model_negaras AS negara ON provinsi.parent_code_negara = negara.code_negara
	`

	queryAdditional += ` WHERE TRUE`

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND kecamatan.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeKabupaten != constants.EMPTY_VALUE {
		queryAdditional += ` AND kecamatan.parent_code_kabupaten = ?`
		args = append(args, input.ParentCodeKabupaten)
	}

	if input.CodeKecamatan != constants.EMPTY_VALUE {
		queryAdditional += ` AND kecamatan.code_kecamatan = ?`
		args = append(args, input.CodeKecamatan)
	}

	if input.CodeNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.code_negara = ?`
		args = append(args, input.CodeNegara)
	}

	if input.NameNegara != constants.EMPTY_VALUE {
		queryAdditional += ` AND negara.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameNegara)+"%")
	}

	if input.CodeProvinsi != constants.EMPTY_VALUE {
		queryAdditional += ` AND provinsi.code_provinsi = ?`
		args = append(args, input.CodeProvinsi)
	}

	if input.NameProvinsi != constants.EMPTY_VALUE {
		queryAdditional += ` AND provinsi.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameProvinsi)+"%")
	}

	if input.NameKabupaten != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameKabupaten)+"%")
	}

	if input.Search != constants.EMPTY_VALUE {
		queryAdditional += ` AND (negara.name LIKE ? OR provinsi.name LIKE ? OR kabupaten.name LIKE ? OR kecamatan.name LIKE ?)`
		args = append(args,
			"%"+strings.ToUpper(input.Search)+"%",
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
