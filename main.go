package main

import (
	"fmt"
	"gin_example/pkg/setting"
	"gin_example/routers"
	"net/http"
)

func main() {

	//router := gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	// Group Routers
	routerInit := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routerInit,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
