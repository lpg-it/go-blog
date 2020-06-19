package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-log"
	"ihome/ihomeWeb/handler"
	"net/http"

	"github.com/micro/go-web"
	_ "ihome/ihomeWeb/models"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.ihomeWeb"),
		web.Version("latest"),
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))

	// register html handler
	service.Handle("/", rou)
	// 获取地区信息
	rou.GET("/api/v1.0/areas", handler.GetArea)
	// 获取 session
	rou.GET("/api/v1.0/session", handler.GetSession)
	// 获取首页轮播图
	rou.GET("/api/v1.0/house/index", handler.GetIndex)
	// 获取图片验证码
	rou.GET("/api/v1.0/imagecode/:uuid", handler.GetImageCd)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}