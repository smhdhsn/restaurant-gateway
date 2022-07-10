package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-gateway/internal/config"
	"github.com/smhdhsn/restaurant-gateway/internal/server/resource"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
)

// Server contains server's services.
type Server struct {
	eRes   *resource.EdibleResource
	oRes   *resource.OrderResource
	router *gin.Engine
}

// New creates a new HTTP server.
func New(er *resource.EdibleResource, or *resource.OrderResource) *Server {
	s := &Server{
		router: gin.New(),
		eRes:   er,
		oRes:   or,
	}

	s.router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiGroup := s.router.Group("/api")
	s.mapEdibleRoutes(apiGroup)
	s.mapOrderRoutes(apiGroup)

	return s
}

// mapEdibleRoutes is responsible for mapping edible's routes.
func (s *Server) mapEdibleRoutes(r *gin.RouterGroup) {
	emRouter := r.Group("/menu")

	emRouter.GET("/", s.eRes.MenuHand.List)
}

// mapOrderRoutes is responsible for mapping order's routes.
func (s *Server) mapOrderRoutes(r *gin.RouterGroup) {
	osRouter := r.Group("/order")

	osRouter.POST("/:foodID", s.oRes.OrderSubmit.Submit)
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(conf *config.ServerConf) error {
	log.Info(fmt.Sprintf("server started listening on port <%d>", conf.Port))
	return s.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
