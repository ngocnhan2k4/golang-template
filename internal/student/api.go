package student

import (
	"Template/internal/entity"
	"Template/pkg/log"
	"fmt"
	"net/http"

	//"strconv"

	//"strconv"
	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
	logger  log.Logger
}

type queryParams struct {
	Page     int     `form:"page" `
	Limit    int     `form:"limit" `
	Year     *int    `form:"year"`
	ClassId  *string `form:"classId"`
	Semester *int    `form:"semester"`
}

func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	//r.GET("/", res.query)
	// r.GET("/scores", res.get)
	r.POST("/", res.create)
	//r.DELETE("/:id", res.delete)
	//r.PUT("/:id", res.update)
	// r.PUT("/scoresd", res.updateScores)
}

// func (r resource) query(c *gin.Context) {
// 	var params queryParams
// 	if err := c.ShouldBindQuery(&params); err != nil {
// 		c.String(http.StatusBadRequest, "Invalid query parameters")
// 		return
// 	}
// 	var intClassId *int
// 	if params.ClassId != nil {
// 		id, _ := strconv.Atoi(*params.ClassId)
// 		intClassId = &id
// 	}
// 	res := r.service.Query(c.Request.Context(), params.Page, params.Limit, intClassId, params.Semester, params.Year)
// 	if res.ErrorMessage != "" {
// 		err := entity.APIError{
// 			Code:    res.ErrorCode,
// 			Message: res.ErrorMessage,
// 		}
// 		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
// 		return
// 	}
// 	data := res.Data.(GetAllClassRequest)
// 	c.JSON(200, data)
// }

func (r resource) create(c *gin.Context) {
	var students []CreateStudentRequest
	err := c.ShouldBind(&students)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Students:", students)
	res := r.service.Create(c.Request.Context(), students)
	if res.ErrorMessage != "" {
		err := entity.APIError{
			Code:    res.ErrorCode,
			Message: res.ErrorMessage,
		}
		var errors []entity.APIError
		for _, e := range res.Errors {
			errors = append(errors, entity.APIError{
				Code:    e.Code,
				Message: e.Message,
				Index:   e.Index,
			})
		}
		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, errors))
		return
	}
	c.JSON(http.StatusCreated, entity.Success("", res.Data, nil))
}

// func (r resource) update(c *gin.Context) {
// 	id := c.Param("id")
// 	var course UpdateClassRequest
// 	err := c.ShouldBind(&course)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	res := r.service.Update(c.Request.Context(), id, course)
// 	if res.ErrorMessage != "" {
// 		err := entity.APIError{
// 			Code:    res.ErrorCode,
// 			Message: res.ErrorMessage,
// 		}
// 		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
// 		return
// 	}
// 	c.JSON(http.StatusOK, entity.Success("", res.Data, nil))
// }

// func (r resource) delete(c *gin.Context) {
// 	id := c.Param("id")
// 	res := r.service.Delete(c.Request.Context(), id)
// 	if res.ErrorMessage != "" {
// 		err := entity.APIError{
// 			Code:    res.ErrorCode,
// 			Message: res.ErrorMessage,
// 		}
// 		c.JSON(http.StatusBadRequest, entity.BadRequest("", nil, &err, nil, nil))
// 		return
// 	}
// 	c.JSON(http.StatusOK, entity.Success("", res.Data, nil))
// }
