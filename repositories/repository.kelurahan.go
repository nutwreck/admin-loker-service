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

func (r *repositoryKelurahan) EntityResults(input *schemes.SchemeKelurahan) (*[]schemes.SchemeGetDataKelurahan, int64, schemes.SchemeDatabaseError) {
	var (
		kelurahan       []models.ModelKelurahan
		result          []schemes.SchemeGetDataKelurahan
		countData       schemes.SchemeCountData
		args            []interface{}
		totalData       int64
		sortData        string = "kelurahan.name ASC"
		queryCountData  string = constants.EMPTY_VALUE
		queryData       string = constants.EMPTY_VALUE
		queryAdditional string = constants.EMPTY_VALUE
	)

	err := make(chan schemes.SchemeDatabaseError, 1)

	db := r.db.Model(&kelurahan)

	if input.Sort != constants.EMPTY_VALUE {
		unScape, _ := url.QueryUnescape(input.Sort)
		sortData = strings.Replace(unScape, "'", constants.EMPTY_VALUE, -1)
	}

	offset := int((input.Page - 1) * input.PerPage)

	//Untuk mengambil jumlah data tanpa limit
	queryCountData = `
		SELECT
			COUNT(kelurahan.*) AS count_data
		FROM model_kelurahans AS kelurahan
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
			kecamatan.name AS name_kecamatan,
			kelurahan.code_kelurahan AS code_kelurahan,
			kelurahan.name AS name_kelurahan
		FROM model_kelurahans AS kelurahan
	`

	queryAdditional = `
		JOIN model_kecamatans AS kecamatan ON kelurahan.parent_code_kecamatan = kecamatan.code_kecamatan
		JOIN model_kabupatens AS kabupaten ON kecamatan.parent_code_kabupaten = kabupaten.code_kabupaten
		JOIN model_provinsis AS provinsi ON kabupaten.parent_code_provinsi = provinsi.code_provinsi
		JOIN model_negaras AS negara ON provinsi.parent_code_negara = negara.code_negara
	`

	queryAdditional += ` WHERE TRUE`

	if input.CodeKelurahan != constants.EMPTY_VALUE {
		queryAdditional += ` AND kelurahan.code_kelurahan = ?`
		args = append(args, input.CodeKelurahan)
	}

	if input.Name != constants.EMPTY_VALUE {
		queryAdditional += ` AND kelurahan.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.Name)+"%")
	}

	if input.ParentCodeKecamatan != constants.EMPTY_VALUE {
		queryAdditional += ` AND kelurahan.parent_code_kecamatan = ?`
		args = append(args, input.ParentCodeKecamatan)
	}

	if input.NameKecamatan != constants.EMPTY_VALUE {
		queryAdditional += ` AND kecamatan.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameKecamatan)+"%")
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

	if input.CodeKabupaten != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.code_kabupaten = ?`
		args = append(args, input.CodeKabupaten)
	}

	if input.NameKabupaten != constants.EMPTY_VALUE {
		queryAdditional += ` AND kabupaten.name LIKE ?`
		args = append(args, "%"+strings.ToUpper(input.NameKabupaten)+"%")
	}

	if input.Search != constants.EMPTY_VALUE {
		queryAdditional += ` AND (negara.name LIKE ? OR provinsi.name LIKE ? OR kabupaten.name LIKE ? OR kecamatan.name LIKE ? OR kelurahan.name LIKE ?)`
		args = append(args,
			"%"+strings.ToUpper(input.Search)+"%",
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
