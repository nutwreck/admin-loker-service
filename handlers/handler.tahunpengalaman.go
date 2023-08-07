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

type handleTahunPengalaman struct {
	tahunPengalaman entities.EntityTahunPengalaman
}

func NewHandlerTahunPengalaman(tahunPengalaman entities.EntityTahunPengalaman) *handleTahunPengalaman {
	return &handleTahunPengalaman{tahunPengalaman: tahunPengalaman}
}

/**
* =============================================
* Handler Ping Status Tahun Pengalaman Teritory
*==============================================
 */

func (h *handleTahunPengalaman) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping Tahun Pengalaman", http.StatusOK, nil)
}

/**
* ============================================
* Handler Create New Tahun Pengalaman Teritory
*=============================================
 */
// CreateTahunPengalaman godoc
// @Summary		Create Tahun Pengalaman
// @Description	Create Tahun Pengalaman
// @Tags		Tahun Pengalaman
// @Accept		json
// @Produce		json
// @Param		TahunPengalaman body schemes.SchemeTahunPengalamanRequest true "Create Tahun Pengalaman"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/tahun-pengalaman/create [post]
func (h *handleTahunPengalaman) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeTahunPengalaman
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorTahunPengalaman(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.tahunPengalaman.EntityCreate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, "Tahun Pengalaman name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Tahun Pengalaman failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Tahun Pengalaman successfully", http.StatusCreated, nil)
}

/**
* =============================================
* Handler Results All Tahun Pengalaman Teritory
*==============================================
 */
// GetListTahunPengalaman godoc
// @Summary		Get List Tahun Pengalaman
// @Description	Get List Tahun Pengalaman
// @Tags		Tahun Pengalaman
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
// @Router /api/v1/tahun-pengalaman/results [get]
func (h *handleTahunPengalaman) HandlerResults(ctx *gin.Context) {
	res, error := h.tahunPengalaman.EntityResults()

	if error.Type == "error_results_01" {
		helpers.APIResponse(ctx, "Tahun Pengalaman data not found", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Tahun Pengalaman data already to use", http.StatusOK, res)
}

/**
* ==============================================
* Handler Result Tahun Pengalaman By ID Teritory
*===============================================
 */
// GetByIDTahunPengalaman godoc
// @Summary		Get By ID Tahun Pengalaman
// @Description	Get By ID Tahun Pengalaman
// @Tags		Tahun Pengalaman
// @Accept		json
// @Produce		json
// @Param		id path string true "Get By ID Tahun Pengalaman"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/tahun-pengalaman/result/{id} [get]
func (h *handleTahunPengalaman) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeTahunPengalaman
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorTahunPengalaman(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.tahunPengalaman.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Tahun Pengalaman data not found for this id %s ", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Tahun Pengalaman data already to use", http.StatusOK, res)
}

/**
* ==============================================
* Handler Delete Tahun Pengalaman By ID Teritory
*===============================================
 */
// GetDeleteTahunPengalaman godoc
// @Summary		Get Delete Tahun Pengalaman
// @Description	Get Delete Tahun Pengalaman
// @Tags		Tahun Pengalaman
// @Accept		json
// @Produce		json
// @Param		id path string true "Delete Tahun Pengalaman"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/tahun-pengalaman/delete/{id} [delete]
func (h *handleTahunPengalaman) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeTahunPengalaman
	id := ctx.Param("id")
	body.ID = id

	errors, code := ValidatorTahunPengalaman(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.tahunPengalaman.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Tahun Pengalaman data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Tahun Pengalaman data for this id %v failed", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Tahun Pengalaman data for this id %s success", id), http.StatusOK, res)
}

/**
* ==============================================
* Handler Update Tahun Pengalaman By ID Teritory
*===============================================
 */
// GetUpdateTahunPengalaman godoc
// @Summary		Get Update Tahun Pengalaman
// @Description	Get Update Tahun Pengalaman
// @Tags		Tahun Pengalaman
// @Accept		json
// @Produce		json
// @Param		id path string true "Update Tahun Pengalaman"
// @Param		TahunPengalaman body schemes.SchemeTahunPengalamanRequest true "Update Tahun Pengalaman"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Security	ApiKeyAuth
// @Router /api/v1/tahun-pengalaman/update/{id} [put]
func (h *handleTahunPengalaman) HandlerUpdate(ctx *gin.Context) {
	var (
		body      schemes.SchemeTahunPengalaman
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

	errors, code := ValidatorTahunPengalaman(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.tahunPengalaman.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Tahun Pengalaman data not found for this id %s ", id), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Tahun Pengalaman data failed for this id %s", id), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Tahun Pengalaman data success for this id %s", id), http.StatusCreated, nil)
}

/**
* ==============================================
*  All Validator User Input For Tahun Pengalaman
*===============================================
 */

func ValidatorTahunPengalaman(ctx *gin.Context, input schemes.SchemeTahunPengalaman, Type string) (interface{}, int) {
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
