package api

import (
	"backend/tools"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func Login(c *gin.Context) {
	// loginUser := model.LoginUser
}

func Logout(c *gin.Context) {

}

func UpdatePassword(c *gin.Context) {

}

func GenerateCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	_, content, _ := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		tools.ResponseError(c, http.StatusInternalServerError, "生成验证码错误："+err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: 1500,
	})
	uuid, err := token.SignedString([]byte(content))
	if err != nil {
		tools.ResponseError(c, http.StatusInternalServerError, "生成验证码UUID错误："+err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{
		"image": item.EncodeB64string(),
		"uuid":  uuid,
	})
}
