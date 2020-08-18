package api

import (
	"backend/models"
	"backend/tools"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func Login(c *gin.Context) {
	var login models.Login

	if err := c.ShouldBind(&login); err != nil {
		tools.ResponseError(c, err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{"message": "登录成功"})
}

func Logout(c *gin.Context) {

}

func UpdatePassword(c *gin.Context) {

}

func GenerateCaptcha(c *gin.Context) {
	captchaWidth, err := strconv.Atoi(c.Query("w"))
	if err != nil {
		captchaWidth = 200
	}
	captchaHeight, err := strconv.Atoi(c.Query("h"))
	if err != nil {
		captchaHeight = 80
	}

	driver := base64Captcha.NewDriverDigit(captchaHeight, captchaWidth, 5, 0.7, 10)
	_, content, _ := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		tools.ResponseError(c, "生成验证码错误："+err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: 1500,
	})
	uuid, err := token.SignedString([]byte(content))
	if err != nil {
		tools.ResponseError(c, "生成验证码UUID错误："+err.Error())
		return
	}

	tools.ResponseSuccess(c, gin.H{
		"image":         item.EncodeB64string(),
		"uuid":          uuid,
		"captchaWidth":  captchaWidth,
		"captchaHeight": captchaHeight,
	})
}
