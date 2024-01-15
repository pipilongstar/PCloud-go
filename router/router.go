package router

import (
	"github.com/gin-gonic/gin"
	"golang/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//处理文件相关接口的路由
	FileRouter(r.Group("/file"))

	return r
}

func FileRouter(r *gin.RouterGroup) {
	//定义中间件或拦截器

	//定义相关路由
	r.POST("/upload", controller.UploadHandler)

}
