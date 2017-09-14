package main

import "github.com/gin-gonic/gin"

func init() {
	router := gin.Default()

	// 前台路由
	router.GET("/", HandleIndex)

	// 后台路由
	admin := router.Group("/admin")
	admin.GET("/login", HandleLogin)
}
