package main

import (
	"fmt"

	"go-copyright-p2/configs"
	"go-copyright-p2/routes"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

var EchoObj *echo.Echo //echo框架对象全局定义
//静态html文件处理
func staticFile() {

	EchoObj.Static("/", "static/pc/home") //根目录设置
	EchoObj.Static("/static", "static")   //全路径处理
	EchoObj.Static("/upload", "static/pc/upload")
	EchoObj.Static("/css", "static/pc/css")
	EchoObj.Static("/assets", "static/pc/assets")
	EchoObj.Static("/user", "static/pc/user")
}

func main() {

	fmt.Printf("get config %v ,%v\n", configs.Config.Common.Port, configs.Config.Db.Connstr)
	EchoObj = echo.New()             //创建echo对象
	EchoObj.Use(middleware.Logger()) //安装日志中间件
	EchoObj.Use(middleware.Recover())
	EchoObj.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	EchoObj.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	staticFile() //静态文件处理调用

	EchoObj.GET("/ping", routes.PingHandler)                        //路由测试函数
	EchoObj.POST("/account", routes.Register)                       //注册账户
	EchoObj.GET("/session", routes.GetSession)                      //session获取
	EchoObj.POST("/login", routes.Login)                            //登陆
	EchoObj.Logger.Fatal(EchoObj.Start(configs.Config.Common.Port)) //启动服务
}
