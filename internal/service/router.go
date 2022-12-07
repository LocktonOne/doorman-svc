package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokene/doorman/internal/config"
	"gitlab.com/tokene/doorman/internal/service/handlers"
	"gitlab.com/tokene/doorman/internal/service/helpers"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxServiceConfig(cfg.ServiceConfig()),
			helpers.CtxRegistryConfig(cfg.RegistryConfig()),
			helpers.CtxEthRPCConfig(cfg.EthRPCConfig()),

			//TODO change when admin's contracts added
			//helpers.CtxNodeAdmins(gosdk.NewNodeAdminsMock(common.HexToAddress("0x750Bd531CEA1f68418DDF2373193CfbD86A69058"))),
		),
	)
	r.Route("/integrations/doorman", func(r chi.Router) {
		r.Get("/validate_token", handlers.ValidateJWT)
		r.Get("/refresh_token", handlers.RefreshJwt)
		r.Post("/get_token_pair", handlers.GenerateJwtPair)
		r.Get("/check_permission/{owner}", handlers.CheckResourcePermission)
		r.Get("/check_permission", handlers.CheckPermission)
		r.Get("/check_purpose", handlers.CheckPurpose)

	})

	return r
}
