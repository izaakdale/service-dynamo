package handlers

import "github.com/izaakdale/service-dynamo/middleware"

func (s *Service) WithRoutes() *Service {
	s.Router.HandleFunc("/store", s.StoreStockPrice).Methods("POST")
	s.Router.Use(middleware.LoggingMiddleware)
	return s
}
