package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"

	go_fiber_helpers "nx-multi-cloud/libs/go/fiber-heplers"
	users "nx-multi-cloud/libs/go/models/users"
	go_mongodb "nx-multi-cloud/libs/go/mongodb"
	go_myutils "nx-multi-cloud/libs/go/utils"
	//"github.com/yurikrupnik/nx-go-playground/my-lib"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name,omitempty" validate:"required,min=3,max=36"`
}

func main() {
	// Connect to the database
	if err := go_mongodb.Connect(); err != nil {
		log.Println("failed to connect")
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: go_fiber_helpers.DefaultErrorHandler,
	})

	app.Use(recover.New())
	// Default middleware config
	app.Use(logger.New())
	//app.Use(csrf.New()) // todo check it - forbidden post events
	app.Use(cors.New())
	apiGroup := app.Group("api")
	apiGroup1 := app.Group("v1")
	apiGroup1.Get("/aris", func(ctx *fiber.Ctx) error {
		return ctx.SendString("tosi")
	})
	apiGroup1.Get("/friends", func(ctx *fiber.Ctx) error {
		return ctx.SendString("friends")
	})
	apiGroup1.Get("/sap", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Yes!")
	})
	//go_models_user.CreateFakeGroup[users.User](apiGroup, "users")
	go_fiber_helpers.CreateFakeGroup[Project](apiGroup, "projects")
	go_fiber_helpers.CreateFakeGroup[users.User](apiGroup, "users")

	app.Get("/dashboard", monitor.New())
	port := go_myutils.Getenv("PORT", "8080")
	log.Println("port", port)
	host := go_myutils.Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	log.Panic(app.Listen(result))
}
