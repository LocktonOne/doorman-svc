package service

import (
	"gitlab.com/tokene/doorman/internal/config"
	"gitlab.com/tokene/doorman/internal/service/handlers"
	"gitlab.com/tokene/doorman/internal/service/helpers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
		),
	)
	r.Route("/doorman", func(r chi.Router) {
		r.Post("/validate_token", handlers.ValidateJWT)
		r.Post("/refresh_token", handlers.RefreshJwt)
		r.Post("/get_token_pair", handlers.GenerateJwtPair)
	})

	return r
}
