package course

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
	logger  log.Logger
}

type queryParams struct {
	Page      int   `form:"page" binding:"required"`
	Limit     int   `form:"limit" binding:"required"`
	FacultyId *int  `form:"facultyId"`
	CourseId  *int  `form:"courseId"`
	IsDeleted *bool `form:"isDeleted"`
}

func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.GET("/", res.query)
	r.GET("/:id", res.get)
	r.POST("/", res.create)
	r.DELETE("/:id", res.delete)
	r.PUT("/:id", res.update)
}

func (r resource) query(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}
	res := r.service.Query(c.Request.Context(), params.Page, params.Limit, params.FacultyId, params.CourseId, params.IsDeleted)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	data := res.Data.(GetAllCourseRequest)
	c.JSON(200, data)
}

func (r resource) get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	res := r.service.Get(c.Request.Context(), id)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	data := res.Data.(GetCourseRequest)
	c.JSON(200, data)
}

func (r resource) create(c *gin.Context) {
	var course CreateCourseRequest
	err := c.ShouldBind(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := r.service.Create(c.Request.Context(), course)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	c.JSON(http.StatusCreated, entity.Success("", res.Data, nil))
}

func (r resource) update(c *gin.Context) {
	id := c.Param("id")
	var course UpdateCourseRequest
	err := c.ShouldBind(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := r.service.Update(c.Request.Context(), id, course)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	c.JSON(http.StatusOK, entity.Success("", res.Data, nil))
}

func (r resource) delete(c *gin.Context) {
	id := c.Param("id")
	res := r.service.Delete(c.Request.Context(), id)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	c.JSON(http.StatusOK, entity.Success("", res.Data, res.Message))
}
