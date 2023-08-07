package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/helpers"
	"github.com/nutwreck/admin-loker-service/pkg"
	"github.com/nutwreck/admin-loker-service/schemes"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type handleKeahlian struct {
	keahlian entities.EntityKeahlian
}

func NewHandlerKeahlian(keahlian entities.EntityKeahlian) *handleKeahlian {
	return &handleKeahlian{keahlian: keahlian}
}

/**
* =====================================
* Handler Ping Status Keahlian Teritory
*======================================
 */

func (h *handleKeahlian) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Keahlian", http.StatusOK, nil)
}

/**
* ====================================
* Handler Create New Keahlian Teritory
*=====================================
 */
// CreateKeahlian godoc
// @Summary		Create Keahlian
// @Description	Create Keahlian
// @Tags		Keahlian
// @Accept		json
// @Produce		json
// @Param		Keahlian body schemes.SchemeKeahlianRequest true "Create Keahlian"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/keahlian/create [post]
func (h *handleKeahlian) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeKeahlian
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorKeahlian(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.keahlian.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Keahlian name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Keahlian failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Keahlian successfully", http.StatusCreated, nil)
}

/**
* =====================================
* Handler Results All Keahlian Teritory
*======================================
 */
// GetListKeahlian godoc
// @Summary		Get List Keahlian
// @Description	Get List Keahlian
// @Tags		Keahlian
// @Accept		json
// @Produce		json
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/keahlian/results [get]
func (h *handleKeahlian) HandlerResults(ctx *gin.Context) {
	res, error := h.keahlian.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Keahlian data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Keahlian data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Result Keahlian By ID Teritory
*=======================================
 */
// GetByIDKeahlian godoc
// @Summary		Get By ID Keahlian
// @Description	Get By ID Keahlian
// @Tags		Keahlian
// @Accept		json
// @Produce		json
// @Param		id path string true "Get By ID Keahlian"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/keahlian/result/{id} [get]
func (h *handleKeahlian) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeKeahlian
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorKeahlian(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.keahlian.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Keahlian data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Keahlian data already to use", http.StatusOK, res)
}

/**
* ======================================
* Handler Delete Keahlian By ID Teritory
*=======================================
 */
// GetDeleteKeahlian godoc
// @Summary		Get Delete Keahlian
// @Description	Get Delete Keahlian
// @Tags		Keahlian
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Keahlian"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/keahlian/delete/{id} [delete]
func (h *handleKeahlian) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeKeahlian
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorKeahlian(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.keahlian.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Keahlian data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Keahlian data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Keahlian data for this id %s success", id), http.StatusOK, res)
}

/**
* ======================================
* Handler Update Keahlian By ID Teritory
*=======================================
 */
// GetUpdateKeahlian godoc
// @Summary		Get Update Keahlian
// @Description	Get Update Keahlian
// @Tags		Keahlian
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Keahlian"
// @Param		Keahlian body schemes.SchemeKeahlianRequest true "Update Keahlian"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/keahlian/update/{id} [put]
func (h *handleKeahlian) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeKeahlian
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

	errors, code := ValidatorKeahlian(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.keahlian.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Keahlian data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Keahlian data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Keahlian data success for this id %s", id), http.StatusCreated, nil)
}

/**
* =============================================
*  All Validator User Input For Keahlian
*==============================================
 */

func ValidatorKeahlian(ctx *gin.Context, input schemes.SchemeKeahlian, Type string) (interface{}, int) {
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
