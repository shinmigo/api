package router

import (
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
	
	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("auth", new(controller.Auth), r, middleware.Cors())
		g.Post("/login")
		
		g.Use(middleware.VerifyToken()).GET("/logout")
	})
}