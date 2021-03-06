package router

import (
	"goshop/admin-api/controller"
	"goshop/admin-api/pkg/core/routerhelper"
	"goshop/admin-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product", new(controller.Product), r, middleware.VerifyToken())
		g.Get("/index")
		g.Post("/add")
		g.Post("/edit")
		g.Post("/delete")
	})
}
