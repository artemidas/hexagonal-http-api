package courses

import (
	"errors"
	mooc "github.com/artemidas/hexagonal-http-api/internal"
	"github.com/artemidas/hexagonal-http-api/internal/creating"
	"github.com/artemidas/hexagonal-http-api/kit/command"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewCourseCommand(
			req.ID,
			req.Name,
			req.Duration,
		))

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrMissingCourseName),
				errors.Is(err, mooc.ErrMissingCourseDuration):
				ctx.JSON(http.StatusBadRequest, err.Error())
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
