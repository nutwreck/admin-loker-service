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

type handleJenisPerusahaan struct {
	jenisPerusahaan entities.EntityJenisPerusahaan
}

func NewHandlerJenisPerusahaan(jenisPerusahaan entities.EntityJenisPerusahaan) *handleJenisPerusahaan {
	return &handleJenisPerusahaan{jenisPerusahaan: jenisPerusahaan}
}

/**
* =============================================
* Handler Ping Status Jenis Perusahaan Teritory
*==============================================
 */

func (h *handleJenisPerusahaan) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Jenis Perusahaan", http.StatusOK, nil)
}

/**
* ============================================
* Handler Create New Jenis Perusahaan Teritory
*=============================================
 */
// CreateJenisPerusahaan godoc
// @Summary		Create Jenis Perusahaan
// @Description	Create Jenis Perusahaan
// @Tags		Jenis Perusahaan
// @Accept		json
// @Produce		json
// @Param		jenisperusahaan body schemes.SchemeJenisPerusahaanRequest true "Create Jenis Perusahaan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/jenis-perusahaan/create [post]
func (h *handleJenisPerusahaan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeJenisPerusahaan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorJenisPerusahaan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.jenisPerusahaan.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Jenis Perusahaan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Jenis Perusahaan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Jenis Perusahaan successfully", http.StatusCreated, nil)
}

/**
* =============================================
* Handler Results All Jenis Perusahaan Teritory
*==============================================
 */
// GetListJenisPerusahaan godoc
// @Summary		Get List Jenis Perusahaan
// @Description	Get List Jenis Perusahaan
// @Tags		Jenis Perusahaan
// @Accept		json
// @Produce		json
// @Param page query int false "Page number for pagination, default is 1"
// @Param perpage query int false "Items per page for pagination, default is 10"
// @Param name query string false "Search by name using LIKE pattern"
// @Success 200 {object} schemes.SchemeResponsesPagination
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/jenis-perusahaan/results [get]
func (h *handleJenisPerusahaan) HandlerResults(ctx *gin.Context) {
	var (
		body       schemes.SchemeJenisPerusahaan
		reqPage    = configs.FirstPage
		reqPerPage = configs.TotalPerPage
		pages      int
		perPages   int
		totalPages int
		totalDatas int
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

	res, totalData, error := h.jenisPerusahaan.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Jenis Perusahaan data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	totalPagesDiv := float64(totalData) / float64(reqPerPage)
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Jenis Perusahaan data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ==============================================
* Handler Result Jenis Perusahaan By ID Teritory
*===============================================
 */
// GetByIDJenisPerusahaan godoc
// @Summary		Get By ID Jenis Perusahaan
// @Description	Get By ID Jenis Perusahaan
// @Tags		Jenis Perusahaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Get By ID Jenis Perusahaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/jenis-perusahaan/result/{id} [get]
func (h *handleJenisPerusahaan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeJenisPerusahaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorJenisPerusahaan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.jenisPerusahaan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Jenis Perusahaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Jenis Perusahaan data already to use", http.StatusOK, res)
}

/**
* ==============================================
* Handler Delete Jenis Perusahaan By ID Teritory
*===============================================
 */
// GetDeleteJenisPerusahaan godoc
// @Summary		Get Delete Jenis Perusahaan
// @Description	Get Delete Jenis Perusahaan
// @Tags		Jenis Perusahaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Jenis Perusahaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/jenis-perusahaan/delete/{id} [delete]
func (h *handleJenisPerusahaan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeJenisPerusahaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorJenisPerusahaan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.jenisPerusahaan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Jenis Perusahaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Jenis Perusahaan data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Jenis Perusahaan data for this id %s success", id), http.StatusOK, res)
}

/**
* ==============================================
* Handler Update Jenis Perusahaan By ID Teritory
*===============================================
 */
// GetUpdateJenisPerusahaan godoc
// @Summary		Get Update Jenis Perusahaan
// @Description	Get Update Jenis Perusahaan
// @Tags		Jenis Perusahaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Jenis Perusahaan"
// @Param		jenisperusahaan body schemes.SchemeJenisPerusahaanRequest true "Update Jenis Perusahaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/jenis-perusahaan/update/{id} [put]
func (h *handleJenisPerusahaan) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeJenisPerusahaan
		activeGet = false
	)
	id := ctx.Param("id")
	body.ID = id
	body.Name = ctx.PostForm("name")
	activeStr := ctx.PostForm("active")
	if activeStr == "true" {
		activeGet = true
	}
	body.Active = &activeGet

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorJenisPerusahaan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.jenisPerusahaan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Jenis Perusahaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Jenis Perusahaan data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Jenis Perusahaan data success for this id %s", id), http.StatusOK, nil)
}

/**
* ==============================================
*  All Validator User Input For Jenis Perusahaan
*===============================================
 */

func ValidatorJenisPerusahaan(ctx *gin.Context, input schemes.SchemeJenisPerusahaan, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
				},
				{
					Tag:     "max",
					Field:   "Name",
					Message: "Name maximal 200 character",
				},
			},
		}
	}

	if Type == "result" || Type == "delete" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "ID",
					Message: "ID is required on param",
				},
				{
					Tag:     "uuid",
					Field:   "ID",
					Message: "ID must be uuid",
				},
				{
					Tag:     "required",
					Field:   "Name",
					Message: "Name is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Name",
					Message: "Name must be lowercase",
				},
				{
					Tag:     "max",
					Field:   "Name",
					Message: "Name maximal 200 character",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
