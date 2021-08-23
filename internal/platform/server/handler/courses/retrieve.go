package courses

import (
	"github.com/artemidas/hexagonal-http-api/internal/retrieving"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveCourses(rs retrieving.CourseService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		courses, err := rs.RetrieveCourses()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, courses)
	}
}
