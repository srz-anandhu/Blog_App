package app

import (
	"blog/app/controller"

	"github.com/go-chi/chi/v5"
)

func apiRouter() chi.Router {
	blogController := controller.NewBlogController()
	userController := controller.NewUserController()
	authorController := controller.NewAuthorController()

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetBlog)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAllUsers)
		r.Get("/{id}", userController.GetUser)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetAllAuthors)
		r.Get("/{id}", authorController.GetAuthor)
	})
	return r
}
