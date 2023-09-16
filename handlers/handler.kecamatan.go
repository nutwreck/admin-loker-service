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

type handleKecamatan struct {
	kecamatan entities.EntityKecamatan
}

func NewHandlerKecamatan(kecamatan entities.EntityKecamatan) *handleKecamatan {
	return &handleKecamatan{kecamatan: kecamatan}
}

/**
* =====================================
* Handler Create New Kecamatan Teritory
*======================================
 */
// CreateDataKecamatan godoc
// @Summary		Create Data Kecamatan
// @Description	Create Data Kecamatan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		Kecamatan body schemes.SchemeKecamatanRequest true "Create Data Kecamatan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kecamatan/create [post]
func (h *handleKecamatan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeKecamatan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKecamatan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kecamatan.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Kecamatan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Kecamatan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Kecamatan successfully", http.StatusCreated, nil)
}

/**
* ======================================
* Handler Results All Kecamatan Teritory
*=======================================
 */
// GetListKecamatan godoc
// @Summary		Get List Kecamatan
// @Description	Get List Kecamatan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : negara.code_negara, negara.name, provinsi.code_provinsi, provinsi.name, kabupaten.code_kabupaten, kabupaten.name, kecamatan.code_kecamatan, kecamatan.name, default is kecamatan.name ASC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param search query string false "Search for data that matches the input from all columns"
// @Param code_kecamatan query string false "Search by Code Kecamatan"
// @Param name query string false "Search by Name Kecamatan"
// @Param parent_code_kabupaten query string false "Search by Code Kabupaten"
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
// @Router /api/v1/wilayah/kecamatan/results [get]
func (h *handleKecamatan) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeKecamatan
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
	codeKecamatanParam := ctx.DefaultQuery("code_kecamatan", constants.EMPTY_VALUE)
	if codeKecamatanParam != constants.EMPTY_VALUE {
		body.CodeKecamatan = codeKecamatanParam
	}
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	parentCodeParam := ctx.DefaultQuery("parent_code_kabupaten", constants.EMPTY_VALUE)
	if parentCodeParam != constants.EMPTY_VALUE {
		body.ParentCodeKabupaten = parentCodeParam
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
	nameKabupatenParam := ctx.DefaultQuery("name_kabupaten", constants.EMPTY_VALUE)
	if nameKabupatenParam != constants.EMPTY_VALUE {
		body.NameKabupaten = nameKabupatenParam
	}

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if parentCodeParam == constants.EMPTY_VALUE && nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name Kecamatan & Parent Code Kabupaten Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
			return
		}
	}

	res, totalData, error := h.kecamatan.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Kecamatan data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Kecamatan data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

func (h *handleKecamatan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeKecamatan
	codes := ctx.Param("code_kecamatan")
	body.CodeKecamatan = codes

	errors, code := ValidatorKecamatan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kecamatan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kecamatan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Kecamatan data already to use", http.StatusOK, res)
}

/**
* =======================================
* Handler Delete Kecamatan By ID Teritory
*========================================
 */
// GetDeleteKecamatan godoc
// @Summary		Get Delete Kecamatan
// @Description	Get Delete Kecamatan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kecamatan path string true "Delete Kecamatan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kecamatan/delete/{code_kecamatan} [delete]
func (h *handleKecamatan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeKecamatan
	codes := ctx.Param("code_kecamatan")
	body.CodeKecamatan = codes

	errors, code := ValidatorKecamatan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kecamatan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kecamatan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Kecamatan data for this code %v failed", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Kecamatan data for this code %s success", codes), http.StatusOK, res)
}

/**
* =======================================
* Handler Update Kecamatan By ID Teritory
*========================================
 */
// GetUpdateKecamatan godoc
// @Summary		Get Update Kecamatan
// @Description	Get Update Kecamatan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kecamatan path string true "Update Kecamatan"
// @Param		Kecamatan body schemes.SchemeKecamatanRequestUpdate true "Update Kecamatan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kecamatan/update/{code_kecamatan} [put]
func (h *handleKecamatan) HandlerUpdate(ctx *gin.Context) {
	var (
		body schemes.SchemeKecamatan
	)
	codes := ctx.Param("code_kecamatan")
	body.CodeKecamatan = codes
	body.ParentCodeKabupaten = ctx.PostForm("parent_code_kabupaten")
	body.Name = ctx.PostForm("name")

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKecamatan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kecamatan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kecamatan data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Kecamatan data failed for this code %s", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Kecamatan data success for this code %s", codes), http.StatusOK, nil)
}

/**
* =======================================
*  All Validator User Input For Kecamatan
*========================================
 */

func ValidatorKecamatan(ctx *gin.Context, input schemes.SchemeKecamatan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKecamatan",
					Message: "CodeKecamatan is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeKabupaten",
					Message: "ParentCodeKabupaten is required on body",
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
					Field:   "CodeKecamatan",
					Message: "CodeKecamatan is required on param",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKecamatan",
					Message: "CodeKecamatan is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeKabupaten",
					Message: "ParentCodeKabupaten is required on body",
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
