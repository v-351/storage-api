package http

func (s *Server) registerRoutes() {
	s.server.GET("warehouse/:warehouse_id/total", s.getWarehouseTotal)
	s.server.POST("/reserve", s.reserveProduct)
	s.server.POST("/release", s.releaseReserveProcuct)
}
