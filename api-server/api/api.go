package api

import (
	"database/sql"
	"gateway-router/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *sqlc.Queries
	router  *gin.Engine
}

func New(db *sql.DB) *Server {
	server := &Server{
		queries: sqlc.New(db),
	}
	router := gin.Default()
	// validate := validator.New(validator.WithRequiredStructEnabled())
	// // validate.RegisterValidation("cidr", utils.IsCIDR)
	// // if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// // 	v.RegisterValidation("cidr", utils.IsCIDR)
	// // }

	router.GET("/vxlan", server.ListVxlanTunnel)
	router.GET("/vxlan/:id", server.GetVxlanTunnel)
	router.POST("/vxlan", server.CreateVxlanTunnel)
	router.DELETE("/vxlan/:id", server.DeleteVxlanTunnel)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
