package Controller

import (
	"MetaWebServer/Service/UserService"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
}

func (c *UserController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Surprise MotherFucker!</h1>",
	}
}

func (c *UserController) GetHello(ctx iris.Context) interface{} {
	params := ctx.URLParams()
	var mapContent []map[string]string
	for i, v := range params {
		mapContent = append(mapContent, map[string]string{i: v})
	}
	return mapContent
}

func (c *UserController) GetAll() interface{} {
	service := UserService.UserBaseService(&UserService.UserBaseResponse{})
	return service.Get()
}

func (c *UserController) PostRawJson(ctx iris.Context) interface{} {
	var userParams map[string]interface{}
	err := ctx.ReadJSON(&userParams)
	if err != nil {
		return mvc.Response{
			ContentType: "text/html",
			Text:        err.Error(),
		}
	}
	return mvc.Response{
		ContentType: "text/html",
		//Text:        fmt.Sprintf("<h1>%s gotten,id is %s!</h1>", userPram.UserName, userPram.UserID),
		Text: fmt.Sprintf("<h1>%s gotten,id is %s!</h1>", userParams["UserName"], userParams["UserID"]),
	}
}

/*
etheir using application/x-www-form-urlencoded or multipart/form-data;
*/
func (c *UserController) PostFormJson(ctx iris.Context) interface{} {

	UserName := ctx.Request().PostFormValue("UserName")
	UserID := ctx.Request().PostFormValue("UserID")
	return mvc.Response{
		ContentType: "text/html",
		//Text:        fmt.Sprintf("<h1>%s gotten,id is %s!</h1>", userPram.UserName, userPram.UserID),
		Text: fmt.Sprintf("<h1>%s gotten,id is %s!</h1>", UserName, UserID),
	}
}

func (c *UserController) HandleError(ctx iris.Context, err error) {
	if iris.IsErrPath(err) {
		return
	}
	ctx.StopWithError(iris.StatusBadRequest, err)
}
