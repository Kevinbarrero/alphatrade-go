package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kevinbarrero/alphatrade-go/db/sqlc"
	"github.com/kevinbarrero/alphatrade-go/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.GET("/coin", server.getCoin())
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
