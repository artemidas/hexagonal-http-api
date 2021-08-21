package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseRepository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDescription := "491f9a9b-615c-4c01-bd92-6672df835362", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDescription)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDescription).
		WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Success(t *testing.T) {
	courseID, courseName, courseDescription := "491f9a9b-615c-4c01-bd92-6672df835362", "Test Course", "10 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDescription)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec("INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDescription).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
