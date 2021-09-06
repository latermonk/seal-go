package main

import "github.com/gin-gonic/gin"

//vender
//1. 一定要将代码新建到gopath目录之下的src
//2. 要记得设置GO111MODULE=off, 开始go modules要记得GO111MODULE=on
// 先查找 gopath/src 这个目录之下的包是否有 goroot/src目录之下找
// 其实就是不做包管理
// 包管理 - 异常处理 泛型

// 能用go modules就用modules 不用去考虑以前的开发模式
// 即使你使用了以前的模式 也可以自动设置为现在的modules模式
// go modules是一个统一的做法
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
