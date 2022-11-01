package main

import (
	"MetaWebServer/Controller"
	"MetaWebServer/DataReflect/Config"
	"MetaWebServer/Utils"
	"fmt"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"log"
	"time"

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
func loadCache() (string, *redis.Database) {
	cacheMgr := Utils.CacheManager(&Utils.CacheService{})
	cacheMgr.Get(Utils.AU_G_KEY)
	propsCache, res := cacheMgr.Get(Utils.PROP_G_KEY)
	var port string
	var props *Config.PropsConfig
	if res {
		props = propsCache.(*Config.PropsConfig)
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
	sessionDB := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      fmt.Sprintf("%s:%d", props.Redis[0].Host, props.Redis[0].Port),
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Username:  "",
		Password:  "",
		Database:  "",
		Prefix:    props.Redis[0].User,
		Driver:    redis.GoRedis(), // defaults.
	})
	return port, sessionDB
}

/*
初始化
@return inited application object
@return server port
*/
func newApp() (*iris.Application, string) {
	serverPort, sessionDB := loadCache()
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	cacheMgr := Utils.CacheManager(&Utils.CacheService{})
	defer sessionDB.Close() // close the database connection if application errored.
	sess := sessions.New(sessions.Config{
		Cookie:          "access_token",
		Expires:         0,
		AllowReclaim:    true,
		CookieSecureTLS: true,
	})
	sess.UseDatabase(sessionDB)

	app.RegisterView(iris.HTML("./Statics", ".html"))
	app.HandleDir("/css", iris.Dir("./Statics/css"))
	app.HandleDir("/js", iris.Dir("./Statics/js"))
	app.Use(sess.Handler())
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
