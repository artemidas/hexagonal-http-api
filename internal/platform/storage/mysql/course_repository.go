package mysql

import (
	"context"
	"database/sql"
	"fmt"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/huandu/go-sqlbuilder"
	"time"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation
type CourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository
func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the mooc.CourseRepository interface
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return err
}

func (r *CourseRepository) Retrieve() ([]mooc.Course, error) {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	courseSQL := courseSQLStruct.SelectFrom(sqlCourseTable)

	query, args := courseSQL.Build()
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return []mooc.Course{}, fmt.Errorf("error trying to retrieve courses from database: %v", err)
	}
	defer rows.Close()

	var courses []mooc.Course
	rows.Scan(courseSQLStruct.Addr(courses)...)

	return courses, nil
}
