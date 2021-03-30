package services

import (
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
)

// JWT default/fixed claims
var (
	iss = "api"
	sub = "uid-"
	aud = "client"
	exp = 24 * 30 * time.Hour // 30 days
	nbf = 30 * time.Second    // 30 seconds
)

// GenerateToken by hmac
func GenerateToken(uid int) (token string, err error) {
	secretkey, _ := beego.AppConfig.String("secretkey")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = uid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return token, nil
}
