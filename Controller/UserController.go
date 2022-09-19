package Controller

import (
	"MetaWebServer/Service/UserService"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type UserContoller struct {
	Ctx     iris.Context
	Model   *UserService.UserBaseResponse
	Session *sessions.Session
}
type UserRestFulContoller struct {
	Session *sessions.Session
}

func (c *UserContoller) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Surprise MotherFucker!</h1>",
	}
}

func (c *UserContoller) GetHello() interface{} {
	params := c.Ctx.URLParams()
	var mapContent []map[string]string
	for i, v := range params {
		mapContent = append(mapContent, map[string]string{i: v})
	}
	return mapContent
}

func (c *UserContoller) GetAll() interface{} {
	service := UserService.UserBaseService(c.Model)
	return service.Get()
}

func (c *UserContoller) HandleError(ctx iris.Context, err error) {
	if iris.IsErrPath(err) {
		return
	}
	ctx.StopWithError(iris.StatusBadRequest, err)
}
