package controllers

func (s *Server) initializeRoutes()  {
	s.Router.HandleFunc("/", s.Home).Methods("GET")
	
}