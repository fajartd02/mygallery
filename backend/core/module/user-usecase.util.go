package module

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretkey               = "asdfasdfasdf"
	ErrUserNotFound         = errors.New("user error: ")
	ErrRecordReportNotFound = errors.New("report error: ")
)

func GenerateJWT(email string, userid uint) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// GeneratehashPassword take password as input and generate new hash password from it
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// IsAuthorized check whether user is authorized or not
func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.Request.Header["Token"] == nil {
			responseWithError(c, 403, "Token is not found")
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			responseWithError(c, 403, err.Error())
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Next()
		}

	}
}
