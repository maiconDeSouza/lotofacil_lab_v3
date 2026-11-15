package server

import (
	"fmt"
	"net/http"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/env"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/handler"
)

type Server struct {
	handler handler.HandlerInterface
	mux     *http.ServeMux
	env     *env.Env
}

func NewServer(handler handler.HandlerInterface, env *env.Env) *Server {
	return &Server{
		handler: handler,
		mux:     http.NewServeMux(),
		env:     env,
	}
}

func (s *Server) Run() *errapp.ErrAppRes {
	s.routes()
	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), s.mux)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao subir servidor", 0)
		return errAppRes
	}

	return nil
}

func (s *Server) routes() {
	// s.mux.HandleFunc("GET /", s.handler.Home)

	s.mux.HandleFunc("POST /fixed", s.handler.CreateNewGameFixed)
	s.mux.HandleFunc("PUT /fixed/{number}", s.handler.UpdateGameFixed)
	s.mux.HandleFunc("GET /fixed/{number}", s.handler.GetGameFixed)
	s.mux.HandleFunc("GET /fixed/all", s.handler.GetAllGamesFixed)

	s.mux.HandleFunc("POST /frequency", s.handler.CreateNewGameFrequency)
	s.mux.HandleFunc("PUT /frequency/{number}", s.handler.UpdateGameFrequency)
	s.mux.HandleFunc("GET /frequency/{number}", s.handler.GetGameFrequency)
	s.mux.HandleFunc("GET /frequency/all", s.handler.GetAllGamesFrequency)

	s.mux.HandleFunc("POST /random", s.handler.CreateNewGameRandom)
	s.mux.HandleFunc("PUT /random/{number}", s.handler.UpdateGameRandom)
	s.mux.HandleFunc("GET /random/{number}", s.handler.GetGameRandom)
	s.mux.HandleFunc("GET /random/all", s.handler.GetAllGamesRandom)

	s.mux.HandleFunc("PUT /scored", s.handler.UpdateScored)

	s.mux.HandleFunc("GET /result/{number}", s.handler.ConfirmResult)

}
