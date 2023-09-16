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

type handleKategoriPekerjaan struct {
	kategoriPekerjaan entities.EntityKategoriPekerjaan
}

func NewHandlerKategoriPekerjaan(kategoriPekerjaan entities.EntityKategoriPekerjaan) *handleKategoriPekerjaan {
	return &handleKategoriPekerjaan{kategoriPekerjaan: kategoriPekerjaan}
}

/**
* ===============================================
* Handler Ping Status Kategori Pekerjaan Teritory
*================================================
 */

func (h *handleKategoriPekerjaan) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Kategori Pekerjaan", http.StatusOK, nil)
}

/**
* ==============================================
* Handler Create New Kategori Pekerjaan Teritory
*===============================================
 */
// CreateKategoriPekerjaan godoc
// @Summary		Create Kategori Pekerjaan
// @Description	Create Kategori Pekerjaan
// @Tags		Kategori Pekerjaan
// @Accept		json
// @Produce		json
// @Param		KategoriPekerjaan body schemes.SchemeKategoriPekerjaanRequest true "Create Kategori Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/kategori-pekerjaan/create [post]
func (h *handleKategoriPekerjaan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeKategoriPekerjaan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKategoriPekerjaan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kategoriPekerjaan.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Kategori Pekerjaan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Kategori Pekerjaan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Kategori Pekerjaan successfully", http.StatusCreated, nil)
}

/**
* ===============================================
* Handler Results All Kategori Pekerjaan Teritory
*================================================
 */
// GetListKategoriPekerjaan godoc
// @Summary		Get List Kategori Pekerjaan
// @Description	Get List Kategori Pekerjaan
// @Tags		Kategori Pekerjaan
// @Accept		json
// @Produce		json
// @Param sort query string false "Use ASC or DESC | Available column sort : id, name, active, created_at, updated_at, default is created_at DESC | If you don't want to use it, fill it blank"
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param name query string false "Search by name using LIKE pattern"
// @Param id query string false "Search by ID"
// @Success 200 {object} schemes.SchemeResponsesPagination
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/kategori-pekerjaan/results [get]
func (h *handleKategoriPekerjaan) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeKategoriPekerjaan
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)

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
	nameParam := ctx.DefaultQuery("name", constants.EMPTY_VALUE)
	if nameParam != constants.EMPTY_VALUE {
		body.Name = nameParam
	}
	idParam := ctx.DefaultQuery("id", constants.EMPTY_VALUE)
	if idParam != constants.EMPTY_VALUE {
		body.ID = idParam
	}

	res, totalData, error := h.kategoriPekerjaan.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Kategori Pekerjaan data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Kategori Pekerjaan data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

func (h *handleKategoriPekerjaan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeKategoriPekerjaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorKategoriPekerjaan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kategoriPekerjaan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kategori Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Kategori Pekerjaan data already to use", http.StatusOK, res)
}

/**
* ================================================
* Handler Delete Kategori Pekerjaan By ID Teritory
*=================================================
 */
// GetDeleteKategoriPekerjaan godoc
// @Summary		Get Delete Kategori Pekerjaan
// @Description	Get Delete Kategori Pekerjaan
// @Tags		Kategori Pekerjaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Kategori Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/kategori-pekerjaan/delete/{id} [delete]
func (h *handleKategoriPekerjaan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeKategoriPekerjaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorKategoriPekerjaan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.kategoriPekerjaan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kategori Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Kategori Pekerjaan data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Kategori Pekerjaan data for this id %s success", id), http.StatusOK, res)
}

/**
* ================================================
* Handler Update Kategori Pekerjaan By ID Teritory
*=================================================
 */
// GetUpdateKategoriPekerjaan godoc
// @Summary		Get Update Kategori Pekerjaan
// @Description	Get Update Kategori Pekerjaan
// @Tags		Kategori Pekerjaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Kategori Pekerjaan"
// @Param		KategoriPekerjaan body schemes.SchemeKategoriPekerjaanRequest true "Update Kategori Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/kategori-pekerjaan/update/{id} [put]
func (h *handleKategoriPekerjaan) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeKategoriPekerjaan
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

	errors, code := ValidatorKategoriPekerjaan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.kategoriPekerjaan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Kategori Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Kategori Pekerjaan data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Kategori Pekerjaan data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ================================================
*  All Validator User Input For Kategori Pekerjaan
*=================================================
 */

func ValidatorKategoriPekerjaan(ctx *gin.Context, input schemes.SchemeKategoriPekerjaan, Type string) (interface{}, int) {
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
