package server

import (
	"fmt"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/courses"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/health"
	"github.com/artemidas/hexagonal-http-api/kit/command"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	// deps
	commandBus command.Bus
}

func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		engine:     gin.New(),
		commandBus: commandBus,
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
	//s.engine.GET("/courses", courses.RetrieveCourses(s.rs))
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
