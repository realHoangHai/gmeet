package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/realHoangHai/gmeet-biz/ent"
	"github.com/realHoangHai/gmeet-biz/ent/migrate"
	"github.com/realHoangHai/gmeet-biz/pkg/config"
	"github.com/realHoangHai/gmeet-biz/pkg/handlers"
	"github.com/realHoangHai/gmeet-biz/pkg/middleware"
	"github.com/realHoangHai/gmeet-biz/pkg/routes"
	"log"
)

func main() {
	conf := config.New()

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name))
	if err != nil {
		log.Fatalf("Failed opening connection to database: %v", err)
	}
	defer client.Close()

	if err = client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	app := fiber.New()
	middleware.SetupMiddleware(app)
	handler := handlers.NewHandlers(client, conf)
	routes.SetupApiV1(app, handler)

	port := "8088"
	addr := flag.String("addr", port, "The address to bind to.")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
