package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/artemidas/hexagonal-http-api/internal/creating"
	"github.com/artemidas/hexagonal-http-api/internal/platform/bus/inmemory"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server"
	"github.com/artemidas/hexagonal-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "root"
	dbPass = "password"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely_courses"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
	)
	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)
	//retrievingCourseService := retrieving.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
