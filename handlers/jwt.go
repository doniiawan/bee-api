package handlers

import (
	"bee-api/models"
	"encoding/json"
	"fmt"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"

)

func init() {
}

func TrimBearer(token string) string {
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:]
	}
	return token
}

func Jwt(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	var uri string = ctx.Input.URI()
	if uri == "/v1/auth" {
		return
	}

	if ctx.Input.Header("Authorization") == "" {
		ctx.Abort(401, "Unauthorized")
		ctx.Output.SetStatus(401)
		resBody, err := json.Marshal(models.APIResponse{Code: 400, Message: "Unauthorized"})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	var tokenString string = TrimBearer(ctx.Input.Header("Authorization"))
	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secretkey, _ := beego.AppConfig.String("secretkey")
		return []byte(secretkey), nil
	})

	if err != nil {
		fmt.Println(err)
		ctx.Output.SetStatus(403)
		var responseBody models.APIResponse = models.APIResponse{Code: 403, Message: err.Error()}
		resBytes, err := json.Marshal(responseBody)
		ctx.Output.Body(resBytes)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		return
	} else {
		ctx.Abort(400, "Unauthorized")
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(models.APIResponse{Code: 403, Message: ctx.Input.Header("Authorization")})
		ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}
}
