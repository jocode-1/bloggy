package routing

import (
	"blog/pkg/config"
	"fmt"
)

func Serve() {

	r := GetRouter()

	configs := config.Get()

	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

}
