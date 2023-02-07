package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	retailtransport "retail-demo/retail-management/transport"
)

func main() {
	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test:test@cluster0.9myhxyq.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	db := client.Database("retail-demo")

	// create a gin router
	router := gin.Default()
	{
		router.GET("/get/info", retailtransport.HandleGetInfo(db))
		router.POST("/sell/product", retailtransport.HandleSellProducts(db))
		router.POST("/import/product", retailtransport.HandleImportProduct(db))
		router.POST("/move/product", retailtransport.HandleMoveProducts(db))
	}

	// start the router
	router.Run(":12345")
}
