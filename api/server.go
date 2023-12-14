package api

import (
	db "github.com/czarro/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serve HTTP request
type Server struct {
	store  db.SQLStore
	router *gin.Engine
}

func NewServer(q *db.Queries) *Server {
	server := &Server{store: db.SQLStore{Queries: q}}
	router := gin.Default()
	router.POST("/customers", server.CreateCustomer)
	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
