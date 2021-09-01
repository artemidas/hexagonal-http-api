package creating

import (
	"context"
	"errors"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/artemidas/hexagonal-http-api/internal/increasing"
	"github.com/artemidas/hexagonal-http-api/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaseService increasing.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaseService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return e.increasingService.Increase(courseCreatedEvt.ID())
}
