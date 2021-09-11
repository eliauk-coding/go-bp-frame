package router

import (
	"gobpframe/apps/demo/actions/user"
	"gobpframe/server"
)

var UserRouter = new(userRouter)

type userRouter struct {
}

func (s *userRouter) InitUserRoute() {
	r := server.Server().GroupRouter("/user")
	usrCtrl := user.NewController()
	r.POST("/login", usrCtrl.Login)
	r.PUT("/create", usrCtrl.Create)
}
