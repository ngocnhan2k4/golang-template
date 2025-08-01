package program

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
	logger  log.Logger
}

func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.GET("/", res.query)
	r.POST("/", res.create)
	r.DELETE("/:id", res.delete)
	r.PUT("/:id", res.update)
}

func (r resource) query(c *gin.Context) {
	res := r.service.Query(c.Request.Context())
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
		return
	}
	c.JSON(200, entity.Success("", res.Data, nil))
}

func (r resource) create(c *gin.Context) {
	var program CreateProgramRequest
	err := c.ShouldBind(&program)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := r.service.Create(c.Request.Context(), program)
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
	var program UpdateProgramRequest
	err := c.ShouldBind(&program)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := r.service.Update(c.Request.Context(), id, program)
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
