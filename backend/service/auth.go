package service

import (
	"fmt"
	"image/color"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

func UserLogin(c *gin.Context) {

}

func UserLogout(c *gin.Context) {

}

func GenerateCaptcha(c *gin.Context) {
	source := "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"
	driver := base64Captcha.NewDriverString(50, 150, 0, base64Captcha.OptionShowHollowLine, 4, source, &color.RGBA{0, 0, 0, 0}, []string{})
	_, content, _ := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		fmt.Println("生成验证码错误", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{ExpiresAt: 15000})
	uuid, err := token.SignedString(content)
	if err != nil {
		fmt.Println("生成验证码UUID错误", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"image": item.EncodeB64string(),
		"uuid":  uuid,
	})
}
