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

	if error.Type == "error_update_01" {
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
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param name query string false "Search by name using LIKE pattern"
// @Param parent_code_kecamatan query string false "Search by Code Kecamatan"
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
	parentCodeParam := ctx.DefaultQuery("parent_code_kecamatan", "")
	if parentCodeParam != constants.EMPTY_VALUE {
		body.ParentCodeKecamatan = parentCodeParam
	}

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if parentCodeParam == constants.EMPTY_VALUE && nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name & Code Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
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

/**
* =========================================
* Handler Result Kelurahan By Code Teritory
*==========================================
 */
// GetByCodeKelurahan godoc
// @Summary		Get By Code Kelurahan
// @Description	Get By Code Kelurahan
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_kelurahan path string true "Get Code Code Kelurahan"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/kelurahan/result/{code_kelurahan} [get]
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
