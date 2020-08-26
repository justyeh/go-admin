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
	var loginUser models.loginUser
	if err := c.ShouldBind(&loginUser); err != nil {
		tools.ResponseBindError(c, err)
		return
	}
	if len(loginUser.ID) == 0 {
		tools.ResponseError(c, "用户名或密码错误")
		return
	}

	user := models.User{ID: loginUser.ID}
	user.UserInfo()
	tools.ResponseSuccess(c, gin.H{
		"message":  "登录成功",
		"token":    "",
		"userInfo": user
	})
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

	// captchaLength：验证码长度 captchaMaxSkew：图片倾斜程度 captchaDotCount：噪点数量，越低越清晰
	captchaLength := 5
	captchaMaxSkew := 0.3
	captchaDotCount := 1
	driver := base64Captcha.NewDriverDigit(captchaHeight, captchaWidth, captchaLength, captchaMaxSkew, captchaDotCount)
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
