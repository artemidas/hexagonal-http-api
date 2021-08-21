package server

import (
	"fmt"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/courses"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		engine:           gin.New(),
		courseRepository: courseRepository,
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
	s.engine.GET("/courses", courses.RetrieveCourses(s.courseRepository))
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
