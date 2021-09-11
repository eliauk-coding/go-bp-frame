package main

import (
	"gobpframe/apps/routers"
	_ "gobpframe/docs"
	"gobpframe/mwares"
	"gobpframe/server"
	"gobpframe/utils/logger"
)

// @title GoBpFrame接口服务
// @version 1.0.0
// @description	GoBpFrame接口服务文档
// @schemes	http

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl http://example
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	server.InitDB()
	dbConn := server.DBConn()
	defer dbConn.Close()

	server.InitServer()
	svr := server.Server()
	svr.Use(mwares.Logger, mwares.CORS, mwares.Catch)
	svr.SetStatic("/assets", "./statics")
	svr.SetSwagger("/swagger/doc.json")
	svr.LoadViews("views/*.html")

	routers.Public()
	//svr.Use(mwares.JWT)
	routers.Protected()

	logger.Infof("listening on %s:%d...", svr.Host, svr.Port)
	logger.Infof("db connected to %s:%d...", dbConn.Host, dbConn.Port)
	if err := svr.Run(); err != nil {
		logger.Fatal("start server failed,", err)
	}
}
