package v1

import (
	caddy "github.com/TheLazarusNetwork/LazarusTunnel/api/v1/caddy"
	nginx "github.com/TheLazarusNetwork/LazarusTunnel/api/v1/nginx"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		caddy.ApplyRoutes(v1)
		nginx.ApplyRoutes(v1)
	}
}
