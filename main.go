package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/pavel-one/anderdog/internal/controller"
	"github.com/pavel-one/anderdog/internal/database"
	"github.com/pavel-one/anderdog/internal/geo"
	"github.com/pavel-one/anderdog/internal/repository"
	"log"
)

func main() {
	app := fiber.New()

	db, err := database.GetInstance()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	g, err := geo.New()
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	rep := repository.NewVisitRepository(db)
	ctrl := controller.New(rep, g)

	app.Get("/", ctrl.Index)
	app.Get("/*", static.New("frontend"))
	log.Fatal(app.Listen(":8080"))
}
