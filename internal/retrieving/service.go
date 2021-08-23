package retrieving

import mooc "github.com/artemidas/hexagonal-http-api/internal"

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

func (s CourseService) RetrieveCourses() ([]mooc.Course, error) {
	return s.courseRepository.Retrieve()
}
