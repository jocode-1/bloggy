package routes

import (
	homeCtrl "blog/internal/modules/home/controllers"
	"blog/pkg/html"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {

	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
			"title": "About Page",
		})
	})

}
