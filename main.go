package golang

import (
	"fmt"
	"golang/router"
)

func main() {

	//设置并获取路由
	r := router.SetupRouter()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err.Error())
	}

}
