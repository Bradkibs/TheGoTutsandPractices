package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"www.github.com/Bradkibs/PostgresTextSearch/DB"
)

var conn *pgxpool.Conn

func main() {
	var err error
	conn, err = DB.Connect()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	app := fiber.New()

	app.Get("/ws/search", websocket.New(handleWebSocket))

	log.Fatal(app.Listen(":3000"))
}
