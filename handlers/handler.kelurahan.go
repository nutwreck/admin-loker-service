package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/configs"
	"github.com/nutwreck/admin-loker-service/constants"
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/helpers"
	"github.com/nutwreck/admin-loker-service/pkg"
	"github.com/nutwreck/admin-loker-service/schemes"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handleKelurahan struct {
	kelurahan entities.EntityKelurahan
}

func NewHandlerKelurahan(kelurahan entities.EntityKelurahan) *handleKelurahan {
	return &handleKelurahan{kelurahan: kelurahan}
}

/**
* =====================================
* Handler Create New Kelurahan Teritory
*======================================
 */
// CreateDataKelurahan godoc
// @Summary		Create Data Kelurahan
// @Description	Create Data Kelurahan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		Kelurahan body schemes.SchemeKelurahanRequest true "Create Data Kelurahan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kelurahan/create [post]
func (h *handleKelurahan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeKelurahan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKelurahan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kelurahan.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Kelurahan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Kelurahan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Kelurahan successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Kelurahan Teritory
*=======================================
 */
// GetListKelurahan godoc
// @Summary		Get List Kelurahan
// @Description	Get List Kelurahan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : negara.code_negara, negara.name, provinsi.code_provinsi, provinsi.name, kabupaten.code_kabupaten, kabupaten.name, kecamatan.code_kecamatan, kecamatan.name, kelurahan.code_kelurahan, kelurahan.name, default is kelurahan.name ASC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param search query string false "Search for data that matches the input from all columns"
// @Param code_kelurahan query string false "Search by Code Kelurahan"
// @Param name query string false "Search by Name Kelurahan"
// @Param parent_code_kecamatan query string false "Search by Code Kecamatan"
// @Param name_kecamatan query string false "Search by Name Kecamatan"
// @Param code_kabupaten query string false "Search by Code Kabupaten"
// @Param name_kabupaten query string false "Search by Name Kabupaten"
// @Param code_provinsi query string false "Search by Code Provinsi"
// @Param name_provinsi query string false "Search by Name Provinsi"
// @Param code_negara query string false "Search by Code Negara"
// @Param name_negara query string false "Search by Name Negara"
// @Success 200 {object} schemes.SchemeResponsesPagination
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kelurahan/results [get]
func (h *handleKelurahan) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeKelurahan
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)

	searchParam := ctx.DefaultQuery("search", constants.EMPTY_VALUE)
	if searchParam != constants.EMPTY_VALUE {
		body.Search = searchParam
	}
	sortParam := ctx.DefaultQuery("sort", constants.EMPTY_VALUE)
	if sortParam != constants.EMPTY_VALUE {
		body.Sort = sortParam
	}
	pageParam := ctx.DefaultQuery("page", constants.EMPTY_VALUE)
	body.Page = reqPage
	if pageParam != constants.EMPTY_VALUE {
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			helpers.APIResponsePagination(ctx, "Convert Params Failed", http.StatusInternalServerError, nil, pages, perPages, totalPages, totalDatas)
			return
		}
		reqPage = page
		body.Page = page
	}
	perPageParam := ctx.DefaultQuery("perpage", constants.EMPTY_VALUE)
	body.PerPage = reqPerPage
	if perPageParam != constants.EMPTY_VALUE {
		perPage, err := strconv.Atoi(perPageParam)
		if err != nil {
			helpers.APIResponsePagination(ctx, "Convert Params Failed", http.StatusInternalServerError, nil, pages, perPages, totalPages, totalDatas)
			return
		}
		reqPerPage = perPage
		body.PerPage = perPage
	}
	codeKelurahanParam := ctx.DefaultQuery("code_kelurahan", constants.EMPTY_VALUE)
	if codeKelurahanParam != constants.EMPTY_VALUE {
		body.CodeKelurahan = codeKelurahanParam
	}
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	parentCodeParam := ctx.DefaultQuery("parent_code_kecamatan", constants.EMPTY_VALUE)
	if parentCodeParam != constants.EMPTY_VALUE {
		body.ParentCodeKecamatan = parentCodeParam
	}
	nameKecamatanParam := ctx.DefaultQuery("name_kecamatan", constants.EMPTY_VALUE)
	if nameKecamatanParam != constants.EMPTY_VALUE {
		body.NameKecamatan = nameKecamatanParam
	}
	codeNegaraParam := ctx.DefaultQuery("code_negara", constants.EMPTY_VALUE)
	if codeNegaraParam != constants.EMPTY_VALUE {
		body.CodeNegara = codeNegaraParam
	}
	nameNegaraParam := ctx.DefaultQuery("name_negara", constants.EMPTY_VALUE)
	if nameNegaraParam != constants.EMPTY_VALUE {
		body.NameNegara = nameNegaraParam
	}
	codeProvinsiParam := ctx.DefaultQuery("code_provinsi", constants.EMPTY_VALUE)
	if codeProvinsiParam != constants.EMPTY_VALUE {
		body.CodeProvinsi = codeProvinsiParam
	}
	nameProvinsiParam := ctx.DefaultQuery("name_provinsi", constants.EMPTY_VALUE)
	if nameProvinsiParam != constants.EMPTY_VALUE {
		body.NameProvinsi = nameProvinsiParam
	}
	codeKabupatenParam := ctx.DefaultQuery("code_kabupaten", constants.EMPTY_VALUE)
	if codeKabupatenParam != constants.EMPTY_VALUE {
		body.CodeKabupaten = codeKabupatenParam
	}
	nameKabupatenParam := ctx.DefaultQuery("name_kabupaten", constants.EMPTY_VALUE)
	if nameKabupatenParam != constants.EMPTY_VALUE {
		body.NameKabupaten = nameKabupatenParam
	}

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if parentCodeParam == constants.EMPTY_VALUE && nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name Kelurahan & Code Kecamatan Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
			return
		}
	}

	res, totalData, error := h.kelurahan.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Kelurahan data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Kelurahan data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

func (h *handleKelurahan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeKelurahan
	codes := ctx.Param("code_kelurahan")
	body.CodeKelurahan = codes

	errors, code := ValidatorKelurahan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kelurahan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kelurahan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Kelurahan data already to use", http.StatusOK, res)
}

/**
* =======================================
* Handler Delete Kelurahan By ID Teritory
*========================================
 */
// GetDeleteKelurahan godoc
// @Summary		Get Delete Kelurahan
// @Description	Get Delete Kelurahan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kelurahan path string true "Delete Kelurahan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kelurahan/delete/{code_kelurahan} [delete]
func (h *handleKelurahan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeKelurahan
	codes := ctx.Param("code_kelurahan")
	body.CodeKelurahan = codes

	errors, code := ValidatorKelurahan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kelurahan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kelurahan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Kelurahan data for this code %v failed", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Kelurahan data for this code %s success", codes), http.StatusOK, res)
}

/**
* =======================================
* Handler Update Kelurahan By ID Teritory
*========================================
 */
// GetUpdateKelurahan godoc
// @Summary		Get Update Kelurahan
// @Description	Get Update Kelurahan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kelurahan path string true "Update Kelurahan"
// @Param		Kelurahan body schemes.SchemeKelurahanRequestUpdate true "Update Kelurahan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kelurahan/update/{code_kelurahan} [put]
func (h *handleKelurahan) HandlerUpdate(ctx *gin.Context) {
	var (
		body schemes.SchemeKelurahan
	)
	codes := ctx.Param("code_kelurahan")
	body.CodeKelurahan = codes
	body.ParentCodeKecamatan = ctx.PostForm("parent_code_kecamatan")
	body.Name = ctx.PostForm("name")

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKelurahan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kelurahan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kelurahan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Kelurahan data failed for this code %s", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Kelurahan data success for this code %s", codes), http.StatusOK, nil)
}

/**
* =======================================
*  All Validator User Input For Kelurahan
*========================================
 */

func ValidatorKelurahan(ctx *gin.Context, input schemes.SchemeKelurahan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKelurahan",
					Message: "CodeKelurahan is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeKecamatan",
					Message: "ParentCodeKecamatan is required on body",
				},
				{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "uppercase",
					Field:   "Name",
					Message: "Name must be uppercase",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKelurahan",
					Message: "CodeKelurahan is required on param",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKelurahan",
					Message: "CodeKelurahan is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeKecamatan",
					Message: "ParentCodeKecamatan is required on body",
				},
				{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "uppercase",
					Field:   "Name",
					Message: "Name must be uppercase",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
