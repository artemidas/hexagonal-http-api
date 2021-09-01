package creating

import (
	"context"
	"errors"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/artemidas/hexagonal-http-api/internal/platform/storage/storagemocks"
	"github.com/artemidas/hexagonal-http-api/kit/event"
	"github.com/artemidas/hexagonal-http-api/kit/event/eventmocks"
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

	eventBusMock := new(eventmocks.Bus)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)

	err = courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_EventBusError(t *testing.T) {
	courseID := "4b75b47f-cfa1-496a-aa3a-90ff20110d00"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(errors.New("something unexpected happened"))

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)
	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Success(t *testing.T) {
	courseID := "4b75b47f-cfa1-496a-aa3a-90ff20110d00"
	courseName := "Test Course"
	courseDuration := "10 months"

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	eventBusMock := new(eventmocks.Bus)
	eventBusMock.On("Publish", mock.Anything, mock.MatchedBy(func(events []event.Event) bool {
		evt := events[0].(mooc.CourseCreatedEvent)
		return evt.CourseName() == courseName
	})).Return(nil)
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	courseService := NewCourseService(courseRepositoryMock, eventBusMock)
	err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
