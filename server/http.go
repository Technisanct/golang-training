package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang-training/config"
	"golang-training/handler"
)

type HTTPServer struct {
	router *gin.Engine
	port   uint16
}

// NewHTTPServer ...
func NewHTTPServer() *HTTPServer {
	corsConfig := configureCORS()

	router := gin.New()
	router.Use(cors.New(corsConfig))

	handler.InitPublicRoutes(router)

	return &HTTPServer{
		router: router,
		port:   config.Get().PortHTTP,
	}
}

func configureCORS() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowOrigins = config.Get().Cors.AllowedOrigins
	corsConfig.AllowMethods = config.Get().Cors.AllowedMethods
	corsConfig.AddAllowHeaders(config.Get().Cors.AllowedHeaders...)
	corsConfig.AddExposeHeaders(config.Get().Cors.ExposeHeaders...)

	return corsConfig
}

// Start ... the server
func (s *HTTPServer) Start() {
	port := fmt.Sprintf(":%v", s.port)
	if err := s.router.Run(port); err != nil {
		panic("Server Startup Failed!")
	}
}
