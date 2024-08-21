package app

import (
	"blog/app/controller"
	"blog/app/repo"
	"blog/app/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func apiRouter() chi.Router {
	var db *sql.DB
	blogController := controller.NewBlogController()
	userController := controller.NewUserController()

	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

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
