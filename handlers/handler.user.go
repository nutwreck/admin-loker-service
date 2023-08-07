package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gpc "github.com/restuwahyu13/go-playground-converter"

	"github.com/nutwreck/admin-loker-service/entities"
	"github.com/nutwreck/admin-loker-service/helpers"
	"github.com/nutwreck/admin-loker-service/pkg"
	"github.com/nutwreck/admin-loker-service/schemes"
)

type handlerUser struct {
	user entities.EntityUser
}

func NewHandlerUser(user entities.EntityUser) *handlerUser {
	return &handlerUser{user: user}
}

/**
* ==================================
* Handler Ping User Status
*==================================
 */

func (h *handlerUser) HandlerPing(ctx *gin.Context) {
	helpers.APIResponse(ctx, "Ping User", http.StatusOK, nil)
}

/**
* ======================================
* Handler Register New Account
*======================================-
 */

// RegisterUser godoc
// @Summary		Register User
// @Description	add by json user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user body schemes.SchemeAddUser true "Add User"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/auth/register [post]
func (h *handlerUser) HandlerRegister(ctx *gin.Context) {
	var body schemes.SchemeUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "register")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	_, error := h.user.EntityRegister(&body)

	if error.Type == "error_register_01" {
		helpers.APIResponse(ctx, "Email already taken", error.Code, nil)
		return
	}

	if error.Type == "error_register_02" {
		helpers.APIResponse(ctx, "Register new user account failed", error.Code, nil)
		return
	}

	helpers.APIResponse(ctx, "Register new user account success", http.StatusOK, nil)
}

/**
* =================================
* Handler Login Auth Account
*==================================
 */

// LoginUser godoc
// @Summary		Login User
// @Description	login user
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user body schemes.SchemeLoginUser true "Login User"
// @Success 200 {object} schemes.SchemeResponses
// @Failure 400 {object} schemes.SchemeResponses400Example
// @Failure 403 {object} schemes.SchemeResponses403Example
// @Failure 404 {object} schemes.SchemeResponses404Example
// @Failure 409 {object} schemes.SchemeResponses409Example
// @Failure 500 {object} schemes.SchemeResponses500Example
// @Router /api/v1/auth/login [post]
func (h *handlerUser) HandlerLogin(ctx *gin.Context) {
	var body schemes.SchemeUser
	err := ctx.ShouldBindJSON(&body)

	if err != nil {
		helpers.APIResponse(ctx, "Parse json data from body failed", http.StatusBadRequest, nil)
		return
	}

	errors, code := ValidatorUser(ctx, body, "login")

	if code > 0 {
		helpers.ErrorResponse(ctx, errors)
		return
	}

	res, error := h.user.EntityLogin(&body)

	if error.Type == "error_login_01" {
		helpers.APIResponse(ctx, "User account is not never registered", error.Code, nil)
		return
	}

	if error.Type == "error_login_02" {
		helpers.APIResponse(ctx, "Email or Password is wrong", error.Code, nil)
		return
	}

	accessToken, errorJwt := pkg.Sign(&schemes.JWtMetaRequest{
		Data:      gin.H{"id": res.ID, "email": res.Email, "role": res.Role},
		SecretKey: pkg.GodotEnv("JWT_SECRET_KEY"),
		Options:   schemes.JwtMetaOptions{Audience: pkg.GodotEnv("JWT_AUD"), ExpiredAt: 1},
	})

	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * 1).Local()

	if errorJwt != nil {
		helpers.APIResponse(ctx, "Generate access token failed", http.StatusBadRequest, nil)
		return
	}

	helpers.APIResponse(ctx, "Login successfully", http.StatusOK, gin.H{"accessToken": accessToken, "expiredAt": expiredAt})
}

func (h *handlerUser) HandlerRefreshToken(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	token := strings.Split(bearer, " ")
	existingToken := strings.TrimSpace(token[1])
	secretKey := pkg.GodotEnv("JWT_SECRET_KEY")

	refreshedToken, err := pkg.RefreshToken(existingToken, secretKey)
	if err != nil {
		helpers.APIResponse(ctx, "Error refreshing token", 500, nil)
		return
	}
	expiredAt := time.Now().Add(time.Duration(time.Minute) * (24 * 60) * 1).Local()

	helpers.APIResponse(ctx, "Refresh Token successfully", http.StatusOK, gin.H{"accessToken": refreshedToken, "expiredAt": expiredAt})
}

/**
* ======================================
*  All Validator User Input For User
*=======================================
 */

func ValidatorUser(ctx *gin.Context, input schemes.SchemeUser, Type string) (interface{}, int) {
	var schema gpc.ErrorConfig

	if Type == "register" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "FirstName",
					Message: "FirstName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "FirstName",
					Message: "FirstName must be lowercase",
				},
				{
					Tag:     "required",
					Field:   "LastName",
					Message: "LastName is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "LastName",
					Message: "LastName must be lowercase",
				},
				{
					Tag:     "required",
					Field:   "Email",
					Message: "Email is required on body",
				},
				{
					Tag:     "email",
					Field:   "Email",
					Message: "Email format is not valid",
				},
				{
					Tag:     "password",
					Field:   "Password",
					Message: "Password is required on body",
				},
				{
					Tag:     "gte",
					Field:   "Password",
					Message: "Password must be greater than equal 8 character",
				},
				{
					Tag:     "required",
					Field:   "Role",
					Message: "Role is required on body",
				},
				{
					Tag:     "lowercase",
					Field:   "Role",
					Message: "Role must be lowercase",
				},
			},
		}
	}

	if Type == "login" {
		schema = gpc.ErrorConfig{
			Options: []gpc.ErrorMetaConfig{
				{
					Tag:     "required",
					Field:   "Email",
					Message: "Email is required on body",
				},
				{
					Tag:     "email",
					Field:   "Email",
					Message: "Email format is not valid",
				},
				{
					Tag:     "required",
					Field:   "Password",
					Message: "Password is required on body",
				},
			},
		}
	}

	err, code := pkg.GoValidator(&input, schema.Options)
	return err, code
}
