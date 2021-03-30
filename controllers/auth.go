package controllers

import (

	// "bee-api/models"

	"bee-api/models"
	"bee-api/services"
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// Define a struct to return when user is authorized
type AuthorizedResponse struct {
	Message string       `json:"message"`
	User    *models.User `json:"user"`
	Token   string       `json:"token"`
}

// Define a struct to return when there is an error
type ErrorResponse struct {
	Message string `json:"message"`
}

// URLMapping ...
func (a *AuthController) URLMapping() {
	a.Mapping("Post", a.Post)
	a.Mapping("GetOne", a.GetOne)
	a.Mapping("GetAll", a.GetAll)
	a.Mapping("Put", a.Put)
	a.Mapping("Delete", a.Delete)
}

// Post ...
// @Title Create
// @Description create Auth
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 201 {object} models.Auth
// @Failure 403 body is empty
// @router / [post]
func (a *AuthController) Post() {
	// fmt.Println(a.Ctx.Input.RequestBody)
	var credential models.UserBasic

	json.Unmarshal(a.Ctx.Input.RequestBody, &credential)
	user, err := models.Login(credential.Username, credential.Password)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		a.Data["json"] = errResponse
	}

	token, err := services.GenerateToken(user.Id)
	if err != nil {
		errResponse := ErrorResponse{
			Message: err.Error(),
		}
		a.Data["json"] = errResponse
	}

	a.Data["json"] = token
	a.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Auth by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Auth
// @Failure 403 :id is empty
// @router /:id [get]
func (a *AuthController) GetOne() {
	a.Data["json"] = "pong"
	a.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Auth
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Auth
// @Failure 403
// @router / [get]
func (a *AuthController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Auth
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 200 {object} models.Auth
// @Failure 403 :id is not int
// @router /:id [put]
func (a *AuthController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Auth
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (a *AuthController) Delete() {

}
