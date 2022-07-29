package routes

import (
	"audioTest/src/application"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

type Router struct {
	router chi.Router
}

func Routes(audioService application.IAudioService) *Router {
	router := &Router{}
	corsOptions := accessControl()
	routes := chi.NewRouter()
	routes.Use(corsOptions.Handler)

	routes.Route("/", func(r chi.Router) {
		audioHandler := AudioHandler{
			audioService: audioService,
		}
		r.Mount("/", audioHandler.router())
	})
	router.router = routes
	return router
}

func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
