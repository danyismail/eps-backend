package utils

import (
	"eps-backend/model"
	"eps-backend/structs"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper             func(c echo.Context) bool
	FuncCheckTokenExist func(c echo.Context) (string, error)
	CustomJwtPayload    struct {
		jwt.RegisteredClaims
		Role     string `json:"role"`
		Username string `json:"username"`
	}
)

func GenerateJWT(secret string, user model.User, tokenLifeTime int) (string, error) {
	claims := CustomJwtPayload{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(tokenLifeTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "eps-backend",
			Subject:   user.ID.String(),
			Audience:  []string{user.Role},
		},
		Username: user.Username,
		Role:     user.Role,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, err
}

func ParseJWT(secret, token string) (*CustomJwtPayload, error) {
	claims := &CustomJwtPayload{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func CreateEchoMiddleware(jwtSecret interface{}) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = jwtSecret
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	extractor := CheckAuthRequest("Authorization", "Bearer")
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := extractor(c) //fetch token on header request
			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			//parse JWT to object JWT Library
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				//validate object token jwt
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})

			if err != nil {
				return c.JSON(http.StatusForbidden, structs.SimpleCommonResponse{
					StatusCode: http.StatusForbidden,
					Message:    err.Error(),
				})
			}
			//check if object token jwt valid and claims ok, then set information from claims to context
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				username := claims["username"].(string)
				role := claims["role"].(string)
				c.Set("username", username)
				c.Set("role", role)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, structs.SimpleCommonResponse{
				StatusCode: http.StatusForbidden,
				Message:    "invalid or expired jwt",
			})
		}
	}
}

/*
CheckAuthRequest use to check if any authorization methon on request,
returns a FuncCheckTokenExist type that contain logic to extracts token from the request header.
*/
func CheckAuthRequest(header string, authScheme string) FuncCheckTokenExist {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", errors.New("missing or malformed jwt")
	}
}
