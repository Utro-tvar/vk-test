package rest

import (
	"log/slog"
	"net/http"

	"github.com/Utro-tvar/vk-test/backend/internal/models"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Service interface {
	GetStatistics() ([]models.Container, error)
	UpdateStatistics([]models.Container)
}

type App struct {
	router chi.Router
}

func New(logger *slog.Logger, service Service) *App {
	app := App{}

	app.router = chi.NewRouter()

	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)
	app.router.Use(middleware.URLFormat)

	app.router.Get("/read", func(w http.ResponseWriter, r *http.Request) {
		stats, err := service.GetStatistics()
		if err != nil {
			logger.Error("rest.GET /read", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, stats)
	})

	app.router.Post("/update", func(w http.ResponseWriter, r *http.Request) {
		var containers []models.Container
		err := render.DecodeJSON(r.Body, &containers)
		if err != nil {
			logger.Error("rest.POST /update", slog.Any("error", err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		service.UpdateStatistics(containers)
		w.WriteHeader(http.StatusOK)
	})

	return &app
}

func (a *App) MustRun(addr string) {
	err := http.ListenAndServe(addr, a.router)
	if err != nil {
		panic(err)
	}
}
