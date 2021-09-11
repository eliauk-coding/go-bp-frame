package routers

import (
	"gobpframe/apps/demo/router"
)

func Public() {
	router.Public()
}

func Protected() {
	router.Protected()
}
