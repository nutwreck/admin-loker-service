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

type handleNegara struct {
	negara entities.EntityNegara
}

func NewHandlerNegara(negara entities.EntityNegara) *handleNegara {
	return &handleNegara{negara: negara}
}

/**
* ==================================
* Handler Create New Negara Teritory
*===================================
 */
// CreateDataNegara godoc
// @Summary		Create Data Negara
// @Description	Create Data Negara
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		negara body schemes.SchemeNegaraRequest true "Create Data Negara"
// @Success 200 {object} schemes.SchemeResponses
// @Success 201 {object} schemes.SchemeResponses201Example
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/negara/create [post]
func (h *handleNegara) HandlerCreate(ctx *gin.Context) {
	var body schemes.SchemeNegara
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorNegara(ctx, body, "create")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.negara.EntityCreate(&body)

	if error.Type == "error_create_01" {
		helpers.APIResponse(ctx, "Negara name already exist", error.Code, nil)
		return
	}

	if error.Type == "error_create_02" {
		helpers.APIResponse(ctx, "Create new Negara failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Create new Negara successfully", http.StatusCreated, nil)
}

/**
* ===================================
* Handler Results All Negara Teritory
*====================================
 */
// GetListNegara godoc
// @Summary		Get List Negara
// @Description	Get List Negara
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param page query int false "Page number for pagination, default is 1 | if you want to disable pagination, fill it with the number 0"
// @Param perpage query int false "Items per page for pagination, default is 10 | if you want to disable pagination, fill it with the number 0"
// @Param name query string false "Search by name using LIKE pattern"
// @Success 200 {object} schemes.SchemeResponsesPagination
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/negara/results [get]
func (h *handleNegara) HandlerResults(ctx *gin.Context) {
	var (
		body          schemes.SchemeNegara
		reqPage       = configs.FirstPage
		reqPerPage    = configs.TotalPerPage
		pages         int
		perPages      int
		totalPagesDiv float64
		totalPages    int
		totalDatas    int
	)

	sortByParam := ctx.DefaultQuery("sortby", constants.EMPTY_VALUE)
	if sortByParam != constants.EMPTY_VALUE {
		body.SortBy = sortByParam
	}
	orderByParam := ctx.DefaultQuery("orderby", constants.EMPTY_VALUE)
	if orderByParam != constants.EMPTY_VALUE {
		body.OrderBy = orderByParam
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

	if reqPage == constants.EMPTY_NUMBER || reqPerPage == constants.EMPTY_NUMBER { //Jika Off Pagination tapi kolom pencarian dikosongkan
		if nameParam == constants.EMPTY_VALUE {
			helpers.APIResponsePagination(ctx, "Kolom Name Tidak Boleh Kosong Jika Pagination Dimatikan!", http.StatusBadRequest, nil, pages, perPages, totalPages, totalDatas)
			return
		}
	}

	res, totalData, error := h.negara.EntityResults(&body)

	if error.Type == "error_results_01" {
		helpers.APIResponsePagination(ctx, "Negara data not found", error.Code, nil, pages, perPages, totalPages, totalDatas)
		return
	}

	pages = reqPage
	perPages = reqPerPage
	if reqPerPage != 0 {
		totalPagesDiv = float64(totalData) / float64(reqPerPage)
	}
	totalPages = int(math.Ceil(totalPagesDiv))
	totalDatas = int(totalData)

	helpers.APIResponsePagination(ctx, "Negara data already to use", http.StatusOK, res, pages, perPages, totalPages, totalDatas)
}

func (h *handleNegara) HandlerResult(ctx *gin.Context) {
	var body schemes.SchemeNegara
	codes := ctx.Param("code_negara")
	body.CodeNegara = codes

	errors, code := ValidatorNegara(ctx, body, "result")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.negara.EntityResult(&body)

	if error.Type == "error_result_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Negara data not found for this code %s ", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Negara data already to use", http.StatusOK, res)
}

/**
* ====================================
* Handler Delete Negara By ID Teritory
*=====================================
 */
// GetDeleteNegara godoc
// @Summary		Get Delete Negara
// @Description	Get Delete Negara
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_negara path string true "Delete Negara"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/negara/delete/{code_negara} [delete]
func (h *handleNegara) HandlerDelete(ctx *gin.Context) {
	var body schemes.SchemeNegara
	codes := ctx.Param("code_negara")
	body.CodeNegara = codes

	errors, code := ValidatorNegara(ctx, body, "delete")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.negara.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Negara data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_delete_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Delete Negara data for this code %v failed", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Delete Negara data for this code %s success", codes), http.StatusOK, res)
}

/**
* ====================================
* Handler Update Negara By ID Teritory
*=====================================
 */
// GetUpdateNegara godoc
// @Summary		Get Update Negara
// @Description	Get Update Negara
// @Tags		Wilayah
// @Accept		json
// @Produce		json
// @Param		code_negara path string true "Update Negara"
// @Param		negara body schemes.SchemeNegaraRequestUpdate true "Update Negara"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 401 {object} schemes.SchemeResponses401Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/wilayah/negara/update/{code_negara} [put]
func (h *handleNegara) HandlerUpdate(ctx *gin.Context) {
	var (
		body schemes.SchemeNegara
	)
	codes := ctx.Param("code_negara")
	body.CodeNegara = codes
	body.ParentCode = ctx.PostForm("parent_code")
	body.Name = ctx.PostForm("name")

	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorNegara(ctx, body, "update")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.negara.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		helpers.APIResponse(ctx, fmt.Sprintf("Negara data not found for this code %s ", codes), error.Code, nil)
		return
	}

	if error.Type == "error_update_02" {
		helpers.APIResponse(ctx, fmt.Sprintf("Update Negara data failed for this code %s", codes), error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, fmt.Sprintf("Update Negara data success for this code %s", codes), http.StatusOK, nil)
}

/**
* ====================================
*  All Validator User Input For Negara
*=====================================
 */

func ValidatorNegara(ctx *gin.Context, input schemes.SchemeNegara, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "create" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeNegara",
					Message: "CodeNegara is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCode",
					Message: "ParentCode is required on body",
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
					Field:   "CodeNegara",
					Message: "CodeNegara is required on param",
				},
			},
		}
	}

	if Type == "update" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "CodeNegara",
					Message: "CodeNegara is required on body",
				},
				{
					Tag:     "required",
					Field:   "ParentCode",
					Message: "ParentCode is required on body",
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
