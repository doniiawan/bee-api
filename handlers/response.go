package handlers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
type ErrorController struct {
	beego.Controller
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (ec *ErrorController) Error400() {
	resp := &Response{Code: 400, Message: "Bad Request"}
	// res, _ := json.Marshal(resp)
	ec.Data["json"] = &resp
	ec.ServeJSON()
}

func (ec *ErrorController) Error404() {
	resp := &Response{Code: 404, Message: "Not Foundz"}
	// res, _ := json.Marshal(resp)
	ec.Data["json"] = &resp
	ec.ServeJSON()
}

func (ec *ErrorController) Error401() {
	resp := &Response{Code: 401, Message: "Unauthorized"}
	res, _ := json.Marshal(resp)
	ec.Data["json"] = res
	ec.ServeJSON()
}

func (ec *ErrorController) Error500() {
	resp := &Response{Code: 500, Message: "Internal Server Error"}
	res, _ := json.Marshal(resp)
	ec.Data["json"] = res
	ec.ServeJSON()
}
