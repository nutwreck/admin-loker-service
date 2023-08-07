package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"

	"github.com/nutwreck/admin-loker-service/schemes"
)

func Sign(configs *schemes.JWtMetaRequest) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * configs.Options.ExpiredAt).Unix()

	claims := jwt.MapClaims{}
	claims["jwt"] = configs.Data
	claims["exp"] = (24 * 60) * expiredAt
	claims["audience"] = configs.Options.Audience
	claims["authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(configs.SecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyToken(accessToken, SecretPublicKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func RefreshToken(existingToken string, secretKey string) (string, error) {
	token, err := jwt.ParseWithClaims(existingToken, &schemes.JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(*schemes.JwtCustomClaims); ok && token.Valid {
		now := time.Now().Unix()

		if now >= claims.Expiration {
			newExpiryTime := now + (24 * 60 * 60) // New expiration time (e.g., 24 hours from now)

			newClaims := schemes.JwtCustomClaims{
				Jwt:           claims.Jwt,
				Expiration:    newExpiryTime,
				Audience:      claims.Audience,
				Authorization: claims.Authorization,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: newExpiryTime,
				},
			}

			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
			refreshedToken, err := newToken.SignedString([]byte(secretKey))
			if err != nil {
				return "", fmt.Errorf("error signing new token: %v", err)
			}

			return refreshedToken, nil
		}

		return existingToken, nil
	}

	return "", fmt.Errorf("invalid token")
}
