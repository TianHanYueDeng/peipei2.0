package main

import (
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"os"
	"peipei2/router"
)

func init(){
	// 日志格式化为 JSON 而不是默认的 ASCII
	log.SetFormatter(&log.JSONFormatter{})
	// 输出 stdout 而不是默认的 stderr，也可以是一个文件
	log.SetOutput(os.Stdout)
	// 只记录严重或以上警告
	log.SetLevel(log.WarnLevel)
}

func main() {


	app := router.NewApp()
	err := app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}

}
