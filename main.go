package main

import (
	"MetaWebServer/Controller"
	"MetaWebServer/DataReflect/Config"
	"MetaWebServer/Utils"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

type Page struct {
	Title string
	Host  string
}

/*
加载缓存配置
@return server port
*/
func loadCache() string {
	cacheMgr := Utils.CacheManager(&Utils.CacheService{})
	cacheMgr.Get(Utils.AU_G_KEY)
	propsCache, res := cacheMgr.Get(Utils.PROP_G_KEY)
	var port string
	if res {
		props := propsCache.(*Config.PropsConfig)
		port = props.ServerPort
	} else {
		port = ""
	}
	//配置多数据源及连接池信息
	DBInfoCache, res := cacheMgr.Get(Utils.DB_G_KEY)
	if res {
		DBInfo := DBInfoCache.(*Config.DBConfig)
		for _, v := range DBInfo.DBInfo {
			Utils.CreateDBConns(v)
		}
	}
	return port
}

/*
初始化
@return inited application object
@return server port
*/
func newApp() (*iris.Application, string) {
	serverPort := loadCache()
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	cacheMgr := Utils.CacheManager(&Utils.CacheService{})
	app.RegisterView(iris.HTML("./Statics", ".html"))
	app.HandleDir("/css", iris.Dir("./Statics/css"))
	app.HandleDir("/js", iris.Dir("./Statics/js"))
	app.Use(func(c iris.Context) {
		var authCnf Config.AuthConfig
		authCnfCache, res := cacheMgr.Get(Utils.AU_G_KEY)
		if res {
			authCnf = authCnfCache.(Config.AuthConfig)
		}
		var Authed bool
		header := c.Request().Header
		forwardIp := header.Values("X-Forwarded-For")
		if authCnf.NeedAuth {
			for _, v := range authCnf.List {
				if forwardIp == nil || forwardIp[0] == v {
					log.SetPrefix(Utils.AUTH_PRX)
					log.Printf("IP Authed:%#v", forwardIp[0])
					Authed = true
					break
				}
			}
			if !Authed {
				err := c.View("403.html", Page{"Mario Warning", forwardIp[0]})
				if err != nil {
					return
				}
				c.EndRequest()
			} else {
				c.Next()
			}
		} else {
			c.Next()
		}
	})
	app.Handle("GET", "/", func(ctx iris.Context) {
		err := ctx.ServeFile("./Statics/index.html")
		if err != nil {
			return
		}
	})
	mvc.Configure(app.Party("/user"), func(a *mvc.Application) {
		a.Handle(new(Controller.UserController))
	})
	return app, serverPort
}

func main() {
	app, port := newApp()
	log.Printf("server running at:%s\n", port)
	err := app.Run(iris.Addr(port), iris.WithConfiguration(iris.YAML("./Config/app.yaml")))
	if err != nil {
		return
	}
}
