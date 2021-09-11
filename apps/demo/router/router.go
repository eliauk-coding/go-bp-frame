package router

import (
	"gobpframe/apps/demo/actions/index"
	"gobpframe/server"
)

func Public() {
	r := server.Server().Router()
	idxCtrl := index.NewController()
	r.GET("/", idxCtrl.Index)
}

func Protected() {
	UserRouter.InitUserRoute()
}
