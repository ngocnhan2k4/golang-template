package student

import "github.com/gin-gonic/gin"

type Resource struct {
	service Service
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := Resource{service: service}

	r.GET("/students", res.query)
}



func (r Resource) query(c *gin.Context) {
	students, err := r.service.Query(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query students"})
		return
	}
	c.JSON(200, students)
}
