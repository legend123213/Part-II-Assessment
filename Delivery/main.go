package main

import (
	"context"
	"log"

	config "github.com/legend123213/ass2/Config"
	controller "github.com/legend123213/ass2/Delivery/Controller"
	routes "github.com/legend123213/ass2/Delivery/Routes"
	repositery "github.com/legend123213/ass2/Repositery"
	usecase "github.com/legend123213/ass2/Usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//mongo connection
func mongoconnection() *mongo.Client {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// mongo connection
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Database.Uri))
	if err != nil {
		log.Fatalf("error creating mongo client: %v", err)
	}
	if err := client.Connect(context.Background()); err != nil {
		log.Fatalf("error connecting to mongo: %v", err)
	}
	return client
	
}
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	mongoClint := mongoconnection()
	userRepo := repositery.NewUserServiceRepo(mongoClint.Database(config.Database.Name))
	userUsecase := usecase.NewUserServiceUsecase(userRepo)
	userController := controller.NewUserController(*userUsecase)
	mainRoute :=routes.NewRoute(userController)
	mainRoute.SetupRouter().Run(config.Port)

}