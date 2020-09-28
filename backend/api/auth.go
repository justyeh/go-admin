package api

import (
	"backend/models"
	"backend/util"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

const JWT_SECRET = "go-admin"

type UserClaims struct {
	ID string `json:"Id"`
	jwt.StandardClaims
}

func (uc UserClaims) Valid() error {
	return nil
}

// 登录
func Login(c *gin.Context) {
	var loginUser models.LoginUser
	if err := c.ShouldBind(&loginUser); err != nil {
		util.ResponseBindError(c, err)
		return
	}

	if err := loginUser.UserWithAccountAndPassword(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	if len(loginUser.ID) == 0 {
		util.ResponseError(c, "用户名或密码错误")
		return
	}

	user := models.User{ID: loginUser.ID}
	if err := user.UserInfoWithID(); err != nil {
		util.ResponseError(c, err.Error())
		return
	}

	token, err := createUserToken(UserClaims{user.ID, jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * 10).Unix()}})
	if err != nil {
		util.ResponseError(c, "生成token失败："+err.Error())
		return
	}
	util.ResponseSuccess(c, gin.H{
		"message":  "登录成功",
		"token":    token,
		"userInfo": user,
	})
}

// 退出登录
func Logout(c *gin.Context) {

}

// 更新密码
func UpdatePassword(c *gin.Context) {

}

// 验证码
func Captcha(c *gin.Context) {
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
		util.ResponseError(c, "生成验证码错误："+err.Error())
		return
	}

	// 生成验证码token
	token, err := createCaptchaToken(content)
	if err != nil {
		util.ResponseError(c, "生成验证码UUID错误："+err.Error())
		return
	}

	util.ResponseSuccess(c, gin.H{
		"image":         item.EncodeB64string(),
		"uuid":          token,
		"captchaWidth":  captchaWidth,
		"captchaHeight": captchaHeight,
	})
}

func createCaptchaToken(content string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * 10).Unix()})
	return token.SignedString([]byte(content))
}

func createUserToken(uc UserClaims) (string, error) {
	/* token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	res, err := token.SignedString(j.SigningKey)
	return res, err */

	return "", nil
}

func parseTolen(tokenString string) string {
	return ""
}
