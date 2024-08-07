package app

import (
	"blog/app/controller"

	"github.com/go-chi/chi/v5"
)

func apiRouter() chi.Router {
	blogController := controller.NewBlogController()

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetBlog)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", )
		r.Get("/{id}", )
	})

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", )
		r.Get("/{id}", )
	})
	return r
}
