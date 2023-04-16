package helper

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "secret"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) { // return data user
	errResp := errors.New("failed to verify token")

	headerToken := c.Request.Header.Get("Authorization")
	isBearer := strings.HasPrefix(headerToken, "Bearer")

	if !isBearer {
		return nil, errResp
	}

	// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuYW5nQG1haWwuY29tIiwiaWQiOjF9.WOWONxV_iXvbXxByQXG0J4Lk0g81cOBkd5yp5mE
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // check signature method
		if !ok {
			return nil, errResp
		}

		return []byte(secretKey), nil
	})

	// validate
	_, ok := token.Claims.(jwt.MapClaims) // adakah error ketika convert token claim ke map claims
	if !ok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}
