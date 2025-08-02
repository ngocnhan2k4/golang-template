package address

import (
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
	r.GET("/provinces", res.getProvinces)
	r.GET("/districts/:id", res.getDistricts)
	r.GET("/wards/:id", res.getWards)
	r.GET("/countries", res.getCountries)

}

type queryParams struct {
	Depth int `form:"depth" binding:"required"`
}

func (r resource) getProvinces(c *gin.Context) {
	provinces := r.service.GetProvinces(c.Request.Context())
	c.String(http.StatusOK, provinces)
}

func (r resource) getDistricts(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}
	districts := r.service.GetDistricts(c.Request.Context(), c.Param("id"), params.Depth)
	c.String(http.StatusOK, districts)
}

func (r resource) getWards(c *gin.Context) {
	var params queryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameters")
		return
	}
	wards := r.service.GetWards(c.Request.Context(), c.Param("id"), params.Depth)
	c.String(http.StatusOK, wards)
}

func (r resource) getCountries(c *gin.Context) {
	countries := r.service.GetCountries(c.Request.Context())
	c.String(http.StatusOK, countries)
}
