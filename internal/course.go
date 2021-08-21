package mooc

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")

// CourseID represents th course unique identifier
type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}
	return CourseID{
		value: v.String(),
	}, nil
}

func (cid CourseID) String() string {
	return cid.value
}

var ErrMissingCourseName = errors.New("missing Course Name")

type CourseName struct {
	value string
}

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrMissingCourseName
	}
	return CourseName{
		value: value,
	}, nil
}

func (cn CourseName) String() string {
	return cn.value
}

var ErrMissingCourseDuration = errors.New("missing Course Duration")

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrMissingCourseDuration
	}
	return CourseDuration{
		value: value,
	}, nil
}

func (cd CourseDuration) String() string {
	return cd.value
}

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	Retrieve() ([]Course, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

func (c Course) ID() CourseID {
	return c.id
}

func (c Course) Name() CourseName {
	return c.name
}

func (c Course) Duration() CourseDuration {
	return c.duration
}
