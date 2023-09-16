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

type handleKabupaten struct {
	kabupaten entities.EntityKabupaten
}

func NewHandlerKabupaten(kabupaten entities.EntityKabupaten) *handleKabupaten {
	return &handleKabupaten{kabupaten: kabupaten}
}

/**
* =====================================
* Handler Create New Kabupaten Teritory
*======================================
 */
// CreateDataKabupaten godoc
// @Summary		Create Data Kabupaten
// @Description	Create Data Kabupaten
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		Kabupaten body schemes.SchemeKabupatenRequest true "Create Data Kabupaten"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kabupaten/create [post]
func (h *handleKabupaten) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeKabupaten
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKabupaten(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kabupaten.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Kabupaten name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Kabupaten failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Kabupaten successfully", http.StatusCreated, nil)
}

/**
* =====================================
* Handler Results All Kabupaten Teritory
*======================================
 */
// GetListKabupaten godoc
// @Summary		Get List Kabupaten
// @Description	Get List Kabupaten
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : negara.code_negara, negara.name, provinsi.code_provinsi, provinsi.name, kabupaten.code_kabupaten, kabupaten.name, default is kabupaten.name ASC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param search query string false "Search for data that matches the input from all columns"
// @Param code_kabupaten query string false "Search by Code Kabupaten"
// @Param name query string false "Search by Name Kabupaten"
// @Param parent_code_provinsi query string false "Search by Code Provinsi"
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
// @Router /api/v1/wilayah/kabupaten/results [get]
func (h *handleKabupaten) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeKabupaten
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)

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
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	parentCodeParam := ctx.DefaultQuery("parent_code_provinsi", constants.EMPTY_VALUE)
	if parentCodeParam != constants.EMPTY_VALUE {
		body.ParentCodeProvinsi = parentCodeParam
	}
	searchParam := ctx.DefaultQuery("search", constants.EMPTY_VALUE)
	if searchParam != constants.EMPTY_VALUE {
		body.Search = searchParam
	}
	codeNegaraParam := ctx.DefaultQuery("code_negara", constants.EMPTY_VALUE)
	if codeNegaraParam != constants.EMPTY_VALUE {
		body.CodeNegara = codeNegaraParam
	}
	nameNegaraParam := ctx.DefaultQuery("name_negara", constants.EMPTY_VALUE)
	if nameNegaraParam != constants.EMPTY_VALUE {
		body.NameNegara = nameNegaraParam
	}
	nameProvinsiParam := ctx.DefaultQuery("name_provinsi", constants.EMPTY_VALUE)
	if nameProvinsiParam != constants.EMPTY_VALUE {
		body.NameProvinsi = nameProvinsiParam
	}
	codeKabupatenParam := ctx.DefaultQuery("code_kabupaten", constants.EMPTY_VALUE)
	if codeKabupatenParam != constants.EMPTY_VALUE {
		body.CodeKabupaten = codeKabupatenParam
	}
	sortParam := ctx.DefaultQuery("sort", constants.EMPTY_VALUE)
	if sortParam != constants.EMPTY_VALUE {
		body.Sort = sortParam
	}

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if parentCodeParam == constants.EMPTY_VALUE && nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name Kabupaten & Parent Code Provinsi Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
			return
		}
	}

	res, totalData, error := h.kabupaten.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Kabupaten data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != constants.EMPTY_NUMBER {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Kabupaten data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

func (h *handleKabupaten) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeKabupaten
	codes := ctx.Param("code_kabupaten")
	body.CodeKabupaten = codes

	errors, code := ValidatorKabupaten(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kabupaten.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kabupaten data not found for this code %s ", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Kabupaten data already to use", http.StatusOK, res)
}

/**
* =======================================
* Handler Delete Kabupaten By ID Teritory
*========================================
 */
// GetDeleteKabupaten godoc
// @Summary		Get Delete Kabupaten
// @Description	Get Delete Kabupaten
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kabupaten path string true "Delete Kabupaten"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kabupaten/delete/{code_kabupaten} [delete]
func (h *handleKabupaten) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeKabupaten
	codes := ctx.Param("code_kabupaten")
	body.CodeKabupaten = codes

	errors, code := ValidatorKabupaten(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kabupaten.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kabupaten data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Kabupaten data for this code %v failed", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Kabupaten data for this code %s success", codes), http.StatusOK, res)
}

/**
* =======================================
* Handler Update Kabupaten By ID Teritory
*========================================
 */
// GetUpdateKabupaten godoc
// @Summary		Get Update Kabupaten
// @Description	Get Update Kabupaten
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kabupaten path string true "Update Kabupaten"
// @Param		kabupaten body schemes.SchemeKabupatenRequestUpdate true "Update Kabupaten"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kabupaten/update/{code_kabupaten} [put]
func (h *handleKabupaten) HandlerUpdate(ctx *gin.Context) {
	var (
		body schemes.SchemeKabupaten
	)
	codes := ctx.Param("code_kabupaten")
	body.CodeKabupaten = codes
	body.ParentCodeProvinsi = ctx.PostForm("parent_code_provinsi")
	body.Name = ctx.PostForm("name")

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKabupaten(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kabupaten.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kabupaten data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Kabupaten data failed for this code %s", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Kabupaten data success for this code %s", codes), http.StatusOK, nil)
}

/**
* =======================================
*  All Validator User Input For Kabupaten
*========================================
 */

func ValidatorKabupaten(ctx *gin.Context, input schemes.SchemeKabupaten, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKabupaten",
					Message: "CodeKabupaten is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeProvinsi",
					Message: "ParentCodeProvinsi is required on body",
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
					Field:   "CodeKabupaten",
					Message: "CodeKabupaten is required on param",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeKabupaten",
					Message: "CodeKabupaten is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeProvinsi",
					Message: "ParentCodeProvinsi is required on body",
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
