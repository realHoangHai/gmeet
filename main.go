package main

import (
	"context"
	"entgo.io/ent/examples/fs/ent"
	"entgo.io/ent/examples/fs/ent/migrate"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/realHoangHai/gmeet-biz/config"
	"github.com/realHoangHai/gmeet-biz/middleware"
	"github.com/realHoangHai/gmeet-biz/routes"
	"log"
)

func main() {
	conf := config.New()

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name))
	if err != nil {
		log.Fatalf("Failed opening connection to database: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	err = client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	app := fiber.New()
	middleware.SetupMiddleware(app)
	routes.SetupApiV1(app)

	port := "8080"
	addr := flag.String("addr", port, "The address to bind to.")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
