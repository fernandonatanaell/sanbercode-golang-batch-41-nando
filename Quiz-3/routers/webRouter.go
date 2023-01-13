package routers

import (
	"quiz-3/controllers"
	"quiz-3/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CekLogin())

	// ROUTE UNTUK BANGUN DATAR
	bangunDatar := router.Group("/bangun-datar")
	bangunDatar.GET("/segitiga-sama-sisi", controllers.SegitigaSamaSisi)
	bangunDatar.GET("/persegi", controllers.Persegi)
	bangunDatar.GET("/persegi-panjang", controllers.PersegiPanjang)
	bangunDatar.GET("/lingkaran", controllers.Lingkaran)

	// ROUTE UNTUK CATEGORIES
	categories := router.Group("/categories")
	categories.GET("/", controllers.GetAllCategories)
	categories.POST("/", controllers.InsertCategory)
	categories.PUT("/:id", controllers.UpdateCategory)
	categories.DELETE("/:id", controllers.DeleteCategory)
	categories.GET("/:id/books", controllers.GetAllBooksFromCategory)

	// ROUTE UNTUK BOOKS
	books := router.Group("/books")
	books.GET("/", controllers.GetAllBooks)
	books.POST("/", controllers.InsertBook)
	books.PUT("/:id", controllers.UpdateBook)
	books.DELETE("/:id", controllers.DeleteBook)

	return router
}
