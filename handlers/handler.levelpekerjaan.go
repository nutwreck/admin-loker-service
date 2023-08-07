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

type handleLevelPekerjaan struct {
	levelPekerjaan entities.EntityLevelPekerjaan
}

func NewHandlerLevelPekerjaan(levelPekerjaan entities.EntityLevelPekerjaan) *handleLevelPekerjaan {
	return &handleLevelPekerjaan{levelPekerjaan: levelPekerjaan}
}

/**
* ============================================
* Handler Ping Status Level Pekerjaan Teritory
*=============================================
 */

func (h *handleLevelPekerjaan) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Level Pekerjaan", http.StatusOK, nil)
}

/**
* ===========================================
* Handler Create New Level Pekerjaan Teritory
*============================================
 */
// CreateLevelPekerjaan godoc
// @Summary		Create Level Pekerjaan
// @Description	Create Level Pekerjaan
// @Tags		Level Pekerjaan
// @Accept		json
// @Produce		json
// @Param		LevelPekerjaan body schemes.SchemeLevelPekerjaanRequest true "Create Level Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/level-pekerjaan/create [post]
func (h *handleLevelPekerjaan) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeLevelPekerjaan
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorLevelPekerjaan(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.levelPekerjaan.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Level Pekerjaan name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Level Pekerjaan failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Level Pekerjaan successfully", http.StatusCreated, nil)
}

/**
* ============================================
* Handler Results All Level Pekerjaan Teritory
*=============================================
 */
// GetListLevelPekerjaan godoc
// @Summary		Get List Level Pekerjaan
// @Description	Get List Level Pekerjaan
// @Tags		Level Pekerjaan
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
// @Router /api/v1/level-pekerjaan/results [get]
func (h *handleLevelPekerjaan) HandlerResults(ctx *gin.Context) {
	res, error := h.levelPekerjaan.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Level Pekerjaan data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Level Pekerjaan data already to use", http.StatusOK, res)
}

/**
* =============================================
* Handler Result Level Pekerjaan By ID Teritory
*==============================================
 */
// GetByIDLevelPekerjaan godoc
// @Summary		Get By ID Level Pekerjaan
// @Description	Get By ID Level Pekerjaan
// @Tags		Level Pekerjaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Get By ID Level Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/level-pekerjaan/result/{id} [get]
func (h *handleLevelPekerjaan) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeLevelPekerjaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorLevelPekerjaan(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.levelPekerjaan.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Level Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Level Pekerjaan data already to use", http.StatusOK, res)
}

/**
* =============================================
* Handler Delete Level Pekerjaan By ID Teritory
*==============================================
 */
// GetDeleteLevelPekerjaan godoc
// @Summary		Get Delete Level Pekerjaan
// @Description	Get Delete Level Pekerjaan
// @Tags		Level Pekerjaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Level Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/level-pekerjaan/delete/{id} [delete]
func (h *handleLevelPekerjaan) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeLevelPekerjaan
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorLevelPekerjaan(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.levelPekerjaan.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Level Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Level Pekerjaan data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Level Pekerjaan data for this id %s success", id), http.StatusOK, res)
}

/**
* =============================================
* Handler Update Level Pekerjaan By ID Teritory
*==============================================
 */
// GetUpdateLevelPekerjaan godoc
// @Summary		Get Update Level Pekerjaan
// @Description	Get Update Level Pekerjaan
// @Tags		Level Pekerjaan
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Level Pekerjaan"
// @Param		LevelPekerjaan body schemes.SchemeLevelPekerjaanRequest true "Update Level Pekerjaan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/level-pekerjaan/update/{id} [put]
func (h *handleLevelPekerjaan) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeLevelPekerjaan
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

	errors, code := ValidatorLevelPekerjaan(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.levelPekerjaan.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Level Pekerjaan data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Level Pekerjaan data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Level Pekerjaan data success for this id %s", id), http.StatusCreated, nil)
}

/**
* =============================================
*  All Validator User Input For Level Pekerjaan
*==============================================
 */

func ValidatorLevelPekerjaan(ctx *gin.Context, input schemes.SchemeLevelPekerjaan, Type string) (interface{}, int) {
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
					Message: "Name maximal 100 character",
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
					Message: "Name maximal 100 character",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
