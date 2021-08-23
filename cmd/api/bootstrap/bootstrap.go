package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/artemidas/hexagonal-http-api/internal/creating"
	"github.com/artemidas/hexagonal-http-api/internal/platform/server"
	"github.com/artemidas/hexagonal-http-api/internal/platform/storage/mysql"
	"github.com/artemidas/hexagonal-http-api/internal/retrieving"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

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

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)
	retrievingCourseService := retrieving.NewCourseService(courseRepository)

	srv := server.New(host, port, creatingCourseService, retrievingCourseService)
	return srv.Run()
}
