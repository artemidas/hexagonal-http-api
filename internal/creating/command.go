package creating

import (
	"context"
	"errors"
	"github.com/artemidas/hexagonal-http-api/kit/command"
)

const CourseCommandType command.Type = "command.creating.course"

// CourseCommand is the command dispatched to create a new course
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command.Type {
	return CourseCommandType
}

type CourseCommandHandler struct {
	service CourseService
}

func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

func (h CourseCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createCourseCmd, ok := cmd.(CourseCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateCourse(
		ctx,
		createCourseCmd.id,
		createCourseCmd.name,
		createCourseCmd.duration,
	)
}
