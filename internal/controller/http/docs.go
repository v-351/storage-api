package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/v-351/storage-api/docs"
)

func (s *Server) registerSwaggerRoute() {
	s.server.GET("/swagger/*", echoSwagger.WrapHandler)
}
