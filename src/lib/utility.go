package lib

import (
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

func MakeSlug(text string) string {
	str := []byte(strings.ToLower(text))

	regE := regexp.MustCompile("[[:space:]]")
	str = regE.ReplaceAll(str, []byte("-"))

	regE = regexp.MustCompile("[[:blank:]]")
	str = regE.ReplaceAll(str, []byte(""))

	return string(str)
}

func GetSignedUser(c *gin.Context) (uuid.UUID, error) {
	user := c.MustGet("user").(jwt.MapClaims)
	return uuid.Parse(user["identifier"].(string))
}

func IsPasswordValid(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func UploadFile(c *gin.Context, file *multipart.FileHeader, destination string) {
	if errUpload := c.SaveUploadedFile(file, destination); errUpload != nil {
		if _, err := os.Stat("/assets/images"); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
				Code:    http.StatusInternalServerError,
				Data:    nil,
				Message: "path not exist " + err.Error(),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, Response{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: "Internal Server Error. Failed while upload icon" + errUpload.Error(),
		})
		return
	}
}
