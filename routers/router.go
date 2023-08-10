package routers

import (
	"gin_example/pkg/setting"
	v1 "gin_example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// get tag list
		apiv1.GET("/tags", v1.GetTags)

		// add new tag
		apiv1.POST("/tags", v1.AddTag)

		// update a tag
		apiv1.PUT("/tags/:id", v1.EditTag)

		// delete a tag
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}
