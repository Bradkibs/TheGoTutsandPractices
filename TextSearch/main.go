package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"www.github.com/Bradkibs/PostgresTextSearch/DB"
)

var pool *pgxpool.Pool

func main() {
	var err error
	pool, err = DB.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()
	app := fiber.New()

	app.Get("/ws/search", websocket.New(func(c *websocket.Conn) { handleWebSocket(c, pool) }))
	// This endpoint adds searchable docs
	app.Post("/docs", func(ctx *fiber.Ctx) error { return handleAddDocs(ctx, pool) })
	log.Fatal(app.Listen(":3000"))
}
