package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture/config"

	_bookHttpDelivery "go-clean-architecture/book/delivery/http"
	_bookMiddleware "go-clean-architecture/book/delivery/http/middleware"
	_bookRepo "go-clean-architecture/book/repository/psql"
	_bookUcase "go-clean-architecture/book/usecase"
)

func main() {
	r := gin.Default()
	config.SetupModels() // new
	db := config.GetDBConnection()
	port := config.GetPortConnection()

	r.Use(_bookMiddleware.Cors())

	repo := _bookRepo.NewPsqlBookRepository(db)
	us := _bookUcase.NewBookUsecase(repo)
	api := r.Group("/v1")

	_bookHttpDelivery.NewBooksHandler(api, us)

	r.Run(port)
}
