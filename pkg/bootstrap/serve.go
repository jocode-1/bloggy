package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/html"
	"blog/pkg/routing"
)

func Serve() {

	config.Set()
	routing.Init()
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()

	routing.Serve()

}
