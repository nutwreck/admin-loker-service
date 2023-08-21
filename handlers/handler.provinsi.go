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

type handleProvinsi struct {
	provinsi entities.EntityProvinsi
}

func NewHandlerProvinsi(provinsi entities.EntityProvinsi) *handleProvinsi {
	return &handleProvinsi{provinsi: provinsi}
}

/**
* ====================================
* Handler Create New Provinsi Teritory
*=====================================
 */
// CreateDataProvinsi godoc
// @Summary		Create Data Provinsi
// @Description	Create Data Provinsi
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		Provinsi body schemes.SchemeProvinsiRequest true "Create Data Provinsi"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/provinsi/create [post]
func (h *handleProvinsi) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeProvinsi
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorProvinsi(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.provinsi.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Provinsi name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Provinsi failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Provinsi successfully", http.StatusCreated, nil)
}

/**
* =====================================
* Handler Results All Provinsi Teritory
*======================================
 */
// GetListProvinsi godoc
// @Summary		Get List Provinsi
// @Description	Get List Provinsi
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param name query string false "Search by name using LIKE pattern"
// @Param parent_code_negara query string false "Search by Code Negara"
// @Success 200 {object} schemes.SchemeResponsesPagination
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/provinsi/results [get]
func (h *handleProvinsi) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeProvinsi
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)
	pageParam := ctx.DefaultQuery("page", "")
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
	perPageParam := ctx.DefaultQuery("perpage", "")
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
	nameParam := ctx.DefaultQuery("name", "")
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	parentCodeParam := ctx.DefaultQuery("parent_code_negara", "")
	if parentCodeParam != constants.EMPTY_VALUE {
		body.ParentCodeNegara = parentCodeParam
	}

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if parentCodeParam == constants.EMPTY_VALUE && nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name & Code Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
			return
		}
	}

	res, totalData, error := h.provinsi.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Provinsi data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Provinsi data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ========================================
* Handler Result Provinsi By Code Teritory
*=========================================
 */
// GetByCodeProvinsi godoc
// @Summary		Get By Code Provinsi
// @Description	Get By Code Provinsi
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_provinsi path string true "Get Code Code Provinsi"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/provinsi/result/{code_provinsi} [get]
func (h *handleProvinsi) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeProvinsi
	codes := ctx.Param("code_provinsi")
	body.CodeProvinsi = codes

	errors, code := ValidatorProvinsi(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.provinsi.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Provinsi data not found for this code %s ", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Provinsi data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Provinsi By ID Teritory
*=======================================
 */
// GetDeleteProvinsi godoc
// @Summary		Get Delete Provinsi
// @Description	Get Delete Provinsi
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_provinsi path string true "Delete Provinsi"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/provinsi/delete/{code_provinsi} [delete]
func (h *handleProvinsi) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeProvinsi
	codes := ctx.Param("code_provinsi")
	body.CodeProvinsi = codes

	errors, code := ValidatorProvinsi(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.provinsi.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Provinsi data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Provinsi data for this code %v failed", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Provinsi data for this code %s success", codes), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Provinsi By ID Teritory
*=======================================
 */
// GetUpdateProvinsi godoc
// @Summary		Get Update Provinsi
// @Description	Get Update Provinsi
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_provinsi path string true "Update Provinsi"
// @Param		provinsi body schemes.SchemeProvinsiRequestUpdate true "Update Provinsi"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/provinsi/update/{code_provinsi} [put]
func (h *handleProvinsi) HandlerUpdate(ctx *gin.Context) {
	var (
		body schemes.SchemeProvinsi
	)
	codes := ctx.Param("code_provinsi")
	body.CodeProvinsi = codes
	body.ParentCodeNegara = ctx.PostForm("parent_code_negara")
	body.Name = ctx.PostForm("name")

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorProvinsi(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.provinsi.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Provinsi data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Provinsi data failed for this code %s", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Provinsi data success for this code %s", codes), http.StatusOK, nil)
}

/**
* ======================================
*  All Validator User Input For Provinsi
*=======================================
 */

func ValidatorProvinsi(ctx *gin.Context, input schemes.SchemeProvinsi, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeProvinsi",
					Message: "CodeProvinsi is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeNegara",
					Message: "ParentCodeNegara is required on body",
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
					Field:   "CodeProvinsi",
					Message: "CodeProvinsi is required on param",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeProvinsi",
					Message: "CodeProvinsi is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCodeNegara",
					Message: "ParentCodeNegara is required on body",
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
