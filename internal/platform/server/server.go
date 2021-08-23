package server

import (
	"fmt"
	"github.com/artemidas/hexagonal-http-api/internal/creating"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/courses"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/health"
	"github.com/artemidas/hexagonal-http-api/internal/retrieving"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	// deps
	cs creating.CourseService
	rs retrieving.CourseService
}

func New(host string, port uint, cs creating.CourseService, rs retrieving.CourseService) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
		cs:       cs,
		rs:       rs,
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
	s.engine.GET("/courses", courses.RetrieveCourses(s.rs))
	s.engine.POST("/courses", courses.CreateHandler(s.cs))
}
