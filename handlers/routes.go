package handlers

func (s *Service) WithRoutes() *Service {
	s.Router.HandleFunc("/store", s.StoreStockPrice).Methods("POST")
	return s
}
