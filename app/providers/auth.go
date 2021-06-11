package providers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Auth Wrapper for authentication
type Auth struct{}

//Authenticate ..validates token
func (auth *Auth) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("token")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("An error occured")
					}
					return authKey, nil
				})
				if error != nil {
					c.AbortWithError(http.StatusBadRequest, error)
				}
				if !token.Valid {
					c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized access or token is expired"))
				} else {
					c.Next()
				}
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
