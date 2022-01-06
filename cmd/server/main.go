package main

import (
	customers "github.com/igvargas/HackthonGo/internal/customers"
	"github.com/igvargas/HackthonGo/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.StorageDB
	repo := customers.NewRepositorySQL(db)
	service := customers.NewServiceSQL(repo)
	controller := handler.NewProduct(service)
	router := gin.Default()
	// router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	// router.PUT("/products/:id", controller.Update())
	// router.PATCH("/products/:id", controller.UpdateProd())
	// router.DELETE("/products/:id", controller.Delete())
	err2 := router.Run()
	if err2 != nil {
		panic(err2)
	}
}
