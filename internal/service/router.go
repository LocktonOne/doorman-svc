package service

import (
	"gitlab.com/tokene/doorman/internal/config"
	"gitlab.com/tokene/doorman/internal/service/handlers"
	"gitlab.com/tokene/doorman/internal/service/helpers"
	gosdk "gitlab.com/tokene/go-sdk"

	"github.com/ethereum/go-ethereum/common"
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
			//TODO change when admin's contracts added
			helpers.CtxNodeAdmins(gosdk.NewNodeAdminsMock(common.HexToAddress("0x750Bd531CEA1f68418DDF2373193CfbD86A69058"))),
		),
	)
	r.Route("/doorman", func(r chi.Router) {
		r.Get("/validate_token", handlers.ValidateJWT)
		r.Get("/refresh_token", handlers.RefreshJwt)
		r.Post("/get_token_pair", handlers.GenerateJwtPair)
		r.Get("/check_permission/{owner}", handlers.CheckResourcePermission)
		r.Get("/check_purpose", handlers.CheckPurpose)
	})

	return r
}
