package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/victoryeo/golang-restapi/models"
)

type loginst struct {
	Username string `json:"username,omitempty"`
}

func GetSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("No header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(GetSecretKey()), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func JWTTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()
}

func Login(c *gin.Context) {

	loginParams := loginst{}
	c.ShouldBindJSON(&loginParams)
	fmt.Print("Login params ", loginParams, "\n")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": loginParams.Username,
		// let the token be valid for one year
		"nbf": time.Date(2022, 01, 01, 12, 0, 0, 0, time.UTC).Unix(), //nbf: not before
		"exp": time.Date(2023, 01, 01, 12, 0, 0, 0, time.UTC).Unix(), //exp: expire
	})

	tokenStr, err := token.SignedString([]byte(GetSecretKey()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SignedResponse{
		Token:   tokenStr,
		Message: "logged in",
	})
	return
}
