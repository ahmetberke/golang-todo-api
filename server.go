package main

import "net/http"

type Server struct {
	PORT 	string
	handler http.Handler
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{PORT: port, handler: handler}
}

func (s *Server) Run()  {
	_ = http.ListenAndServe(s.PORT, s.handler)
}

func (s *Server) NewApp(p string, h func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(p, h)
}