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

type handlePendidikan struct {
	pendidikan entities.EntityPendidikan
}

func NewHandlerPendidikan(pendidikan entities.EntityPendidikan) *handlePendidikan {
	return &handlePendidikan{pendidikan: pendidikan}
}

/**
* =======================================
* Handler Ping Status Pendidikan Teritory
*========================================
 */

func (h *handlePendidikan) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Pendidikan", http.StatusOK, nil)
}

/**
* ======================================
* Handler Create New Pendidikan Teritory
*=======================================
 */
// CreatePendidikan godoc
// @Summary		Create Pendidikan
// @Description	Create Pendidikan
// @Tags		Pendidikan
// @Accept		json
// @Produce		json
// @Param		Pendidikan body schemes.SchemePendidikanRequest true "Create Pendidikan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/pendidikan/create [post]
func (h *handlePendidikan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemePendidikan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorPendidikan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.pendidikan.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Pendidikan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Pendidikan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Pendidikan successfully", http.StatusCreated, nil)
}

/**
* =======================================
* Handler Results All Pendidikan Teritory
*========================================
 */
// GetListPendidikan godoc
// @Summary		Get List Pendidikan
// @Description	Get List Pendidikan
// @Tags		Pendidikan
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
// @Router /api/v1/pendidikan/results [get]
func (h *handlePendidikan) HandlerResults(ctx *gin.Context) {
	var (
		body       schemes.SchemePendidikan
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

	res, totalData, error := h.pendidikan.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Pendidikan data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	totalPagesDiv := float64(totalData) / float64(reqPerPage)
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Pendidikan data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

/**
* ========================================
* Handler Result Pendidikan By ID Teritory
*=========================================
 */
// GetByIDPendidikan godoc
// @Summary		Get By ID Pendidikan
// @Description	Get By ID Pendidikan
// @Tags		Pendidikan
// @Accept		json
// @Produce		json
// @Param		id path string true "Get By ID Pendidikan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/pendidikan/result/{id} [get]
func (h *handlePendidikan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemePendidikan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorPendidikan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.pendidikan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pendidikan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Pendidikan data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Pendidikan By ID Teritory
*=======================================
 */
// GetDeletePendidikan godoc
// @Summary		Get Delete Pendidikan
// @Description	Get Delete Pendidikan
// @Tags		Pendidikan
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Pendidikan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/pendidikan/delete/{id} [delete]
func (h *handlePendidikan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemePendidikan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorPendidikan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.pendidikan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pendidikan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Pendidikan data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Pendidikan data for this id %s success", id), http.StatusOK, res)
}

/**
* ========================================
* Handler Update Pendidikan By ID Teritory
*=========================================
 */
// GetUpdatePendidikan godoc
// @Summary		Get Update Pendidikan
// @Description	Get Update Pendidikan
// @Tags		Pendidikan
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Pendidikan"
// @Param		Pendidikan body schemes.SchemePendidikanRequest true "Update Pendidikan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/pendidikan/update/{id} [put]
func (h *handlePendidikan) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemePendidikan
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

	errors, code := ValidatorPendidikan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.pendidikan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Pendidikan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Pendidikan data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Pendidikan data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ========================================
*  All Validator User Input For Pendidikan
*=========================================
 */

func ValidatorPendidikan(ctx *gin.Context, input schemes.SchemePendidikan, Type string) (interface{}, int) {
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
					Message: "Name maximal 75 character",
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
					Message: "Name maximal 75 character",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
