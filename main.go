package main

import (
	todotrpt "Project/module/item/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", todotrpt.HandleCreateProfile(db))        // create item
		v1.GET("/items", todotrpt.HandleListProfile(db))           // list items
		v1.GET("/items/:id", todotrpt.HandleFindAProfile(db))      // get an item by ID
		v1.PUT("/items/:id", todotrpt.HandleUpdateAnItem(db))      // edit an item by ID
		v1.DELETE("/items/:id", todotrpt.HandleDeleteAProfile(db)) // delete an item by ID
	}

	router.Run()
}
