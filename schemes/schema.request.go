package schemes

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtMetaOptions struct {
	Audience  string
	ExpiredAt time.Duration
}

type JWtMetaRequest struct {
	Data      map[string]interface{}
	SecretKey string
	Options   JwtMetaOptions
}

type JwtCustomClaims struct {
	Jwt           string `json:"jwt"`
	Expiration    int64  `json:"exp"`
	Audience      string `json:"audience"`
	Authorization bool   `json:"authorization"`
	jwt.StandardClaims
}
