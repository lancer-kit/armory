package infoworker

import (
	"context"
	"net/http"
	"time"

	"gitlab.inn4science.com/gophers/service-kit/routines"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"gitlab.inn4science.com/gophers/service-kit/api"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
)

type Info struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Tag     string `json:"tag"`
	Build   string `json:"build"`
}

type InfoWorker struct {
	ParentCtx context.Context
	Info      Info
	api.Server
}

func GetInfoWorker(cfg api.Config, ctx context.Context, info Info) *InfoWorker {
	res := &InfoWorker{
		ParentCtx: ctx,
		Info:      info,
		Server: api.Server{
			Name: "info worker",
			GetConfig: func() api.Config {
				return cfg
			},
		},
	}
	res.GetRouter = res.GetInfoRouter
	return res
}

func (iw *InfoWorker) GetInfoRouter(logger *logrus.Entry, cfg api.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))

	r.Route("/debug", func(r chi.Router) {
		r.Route("/stats", func(r chi.Router) {
			r.Get("/version", iw.Version)
			r.Get("/workers", iw.Workers)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}

func (iw *InfoWorker) Version(w http.ResponseWriter, r *http.Request) {
	parentChief := iw.ParentCtx.Value("chief").(*routines.Chief)
	if parentChief == nil {
		render.ResultServerError.Render(w)
		return
	}

}

func (iw *InfoWorker) Workers(w http.ResponseWriter, r *http.Request) {
	parentChief := iw.ParentCtx.Value("chief").(*routines.Chief)
	if parentChief == nil {
		render.ResultServerError.Render(w)
		return
	}
}
