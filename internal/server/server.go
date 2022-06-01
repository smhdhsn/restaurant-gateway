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
	uRes   *resource.UserResource
	router *gin.Engine
}

// New creates a new HTTP server.
func New(er *resource.EdibleResource, or *resource.OrderResource, ur *resource.UserResource) *Server {
	s := &Server{
		router: gin.New(),
		eRes:   er,
		oRes:   or,
		uRes:   ur,
	}

	s.router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiGroup := s.router.Group("/api")
	s.mapEdibleRoutes(apiGroup)
	s.mapOrderRoutes(apiGroup)
	s.mapUserRoutes(apiGroup)

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

// mapUserRoutes is responsible for mapping user's routes.
func (s *Server) mapUserRoutes(r *gin.RouterGroup) {
	uRouter := r.Group("/users")

	uRouter.POST("/", s.uRes.SourceHand.Store)
	uRouter.GET("/:userID", s.uRes.SourceHand.Find)
	uRouter.PUT("/:userID", s.uRes.SourceHand.Update)
	uRouter.DELETE("/:userID", s.uRes.SourceHand.Destroy)
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(conf *config.ServerConf) error {
	log.Info(fmt.Sprintf("server started listening on port <%d>", conf.Port))
	return s.router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
}
