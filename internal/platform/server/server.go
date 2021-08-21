package server

import (
	"fmt"
	"github.com/artemidas/hexagonal-http-api/internal/platform/handler/health"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
