package server

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gobpframe/config"
)

type HttpServer struct {
	Host          string
	Port          int
	MaxBodySize   int64
	MaxHeaderSize int
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration

	engine *gin.Engine
}

var (
	server      *HttpServer
	serverMutex sync.RWMutex
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func InitServer() {
	serverMutex.Lock()
	defer serverMutex.Unlock()

	if server != nil {
		return
	}

	engine := gin.New()
	engine.MaxMultipartMemory = 8 << 20 // 8MB
	port := config.GetInt("Server.Port")
	if port == 0 {
		port = 8999
	}
	server = &HttpServer{
		Host:          config.GetStr("Server.Host"),
		Port:          port,
		ReadTimeout:   config.GetDuration("Server.ReadTimeout") * time.Second,
		WriteTimeout:  config.GetDuration("Server.WriteTimeout") * time.Second,
		MaxBodySize:   config.GetInt64("Server.MaxBodySize"),
		MaxHeaderSize: config.GetInt("Server.MaxHeaderSize"),
		engine:        engine,
	}
}

func Server() *HttpServer {
	serverMutex.RLock()
	defer serverMutex.RUnlock()
	return server
}

func (s *HttpServer) Run() error {
	port := 8080
	if s.Port > 0 {
		port = s.Port
	}

	host := ""
	if s.Host != "" {
		host = s.Host
	}

	svr := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        s.Router(),
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 2 << 10, // 2MB
	}

	if s.MaxBodySize > 0 {
		s.engine.MaxMultipartMemory = s.MaxBodySize
	}

	if s.ReadTimeout > 0 {
		svr.ReadTimeout = s.ReadTimeout
	}

	if s.WriteTimeout > 0 {
		svr.WriteTimeout = s.WriteTimeout
	}

	if s.MaxHeaderSize > 0 {
		svr.MaxHeaderBytes = s.MaxHeaderSize
	}

	return svr.ListenAndServe()
}

func (s *HttpServer) Use(middlewares ...gin.HandlerFunc) {
	s.engine.Use(middlewares...)
}

func (s *HttpServer) Router() *gin.Engine {
	return s.engine
}

func (s *HttpServer) LoadViews(views string) {
	s.engine.LoadHTMLGlob(views)
}

func (s *HttpServer) GroupRouter(group string) *gin.RouterGroup {
	return s.engine.Group(group)
}

func (s *HttpServer) SetSwagger(path string) {
	url := ginSwagger.URL(config.GetStr("Swagger.ApiHost") + path)
	s.Router().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func (s *HttpServer) SetStatic(relativePath string, localDir string) {
	s.engine.Static(relativePath, localDir)
}
