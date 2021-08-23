package creating

import (
	"context"
	"errors"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/artemidas/hexagonal-http-api/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseService_CreateCourse_RepositoryError(t *testing.T) {
	courseID := "6bb74a6a-2757-4626-8f76-7d15837fb0e0"
	courseName := "Test Course"
	courseDuration := "10 months"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, course).
		Return(errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock)
	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Success(t *testing.T) {
	courseID := "4b75b47f-cfa1-496a-aa3a-90ff20110d00"
	courseName := "Test Course"
	courseDuration := "10 months"

	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, course).
		Return(nil)

	courseService := NewCourseService(courseRepositoryMock)
	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
