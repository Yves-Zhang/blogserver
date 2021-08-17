package main

import (
	"yyds/routers"

	middleware "yyds/middleWare"

	"github.com/kataras/iris/v12"
)

func main(){
	app := iris.New()

	app.HandleDir("www", "./www")

	app.Use(middleware.LogMiddleware)

	books := routers.BooksRouter{}

	booksAPI := app.Party("/books")
	{	
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/list", books.List)
		booksAPI.Post("/list", books.List)

		// POST: http://localhost:8080/books
		booksAPI.Get("/create", books.Create)
		booksAPI.Post("/create", books.Create)
	}

	app.Listen(":8000")
}