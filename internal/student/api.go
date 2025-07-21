package student

import (
	utils "Template/pkg/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resource struct {
	service Service
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := Resource{service: service}

	r.GET("/", res.query)
}

// ListStudents godoc
//
//	@Summary		List students
//	@Description	get students
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		Student
//	@Failure		500	{object}	utils.ErrorResponse
//	@Router			/students [get]
func (r Resource) query(c *gin.Context) {
	students, err := r.service.Query(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ErrorResponse{
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,})
		return
	}
	c.JSON(200, students)
}
