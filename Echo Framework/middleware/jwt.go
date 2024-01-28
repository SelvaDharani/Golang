package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// ValidateToken validates the JWT token
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) != 2 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		}
		tokenString = splitToken[1]
		isValid, err := validateAndDecodeJWT(tokenString)
		if err != nil || !isValid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		return next(c)
	}
}

func validateAndDecodeJWT(tokenString string) (bool, error) {
	// Parse and validate the JWT token
	var mySigningKey = []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key for validation
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
