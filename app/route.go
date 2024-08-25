package app

import (
	"blog/app/controller"
	"blog/app/repo"
	"blog/app/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func apiRouter(db *sql.DB) chi.Router {

	// Author
	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	// Blog
	blogRepo := repo.NewBlogRepo(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// User
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetBlog)
		r.Delete("/{id}", blogController.DeleteBlog)
		r.Post("/create", blogController.CreateBlog)
		r.Put("/{id}", blogController.UpdateBlog)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAllUsers)
		r.Get("/{id}", userController.GetUser)
		r.Delete("/{id}", userController.DeleteUser)
		r.Post("/create", userController.CreateUser)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetAllAuthors)
		r.Get("/{id}", authorController.GetAuthor)
		r.Delete("/{id}", authorController.DeleteAuthor)
		r.Post("/create", authorController.CreateAuthor)
		r.Put("/{id}", authorController.UpdateAuthor)
	})

	return r
}
