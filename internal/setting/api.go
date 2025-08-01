package setting

import (
	"Template/pkg/log"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service Service
	logger  log.Logger
}

func RegisterHandlers(r *gin.RouterGroup, service Service, logger log.Logger) {
	res := resource{service, logger}
	r.GET("/", res.get)
	r.PUT("/", res.update)
}

func (r resource) update(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.New().Error("Failed to read request body", "error", err)
		return
	}
	domain := string(body)
	res := r.service.Update(c.Request.Context(), domain)
	if !res {
		c.JSON(http.StatusBadRequest, "Invalid domain format")
		return
	}
	c.JSON(http.StatusOK, "Email domain has been updated successfully")
}

func (r resource) get(c *gin.Context) {
	setting := r.service.Get(c.Request.Context())
	c.JSON(http.StatusOK, setting)
}
