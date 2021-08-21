package courses

import (
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveCourses(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		courses, err := courseRepository.Retrieve()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, courses)
	}
}
