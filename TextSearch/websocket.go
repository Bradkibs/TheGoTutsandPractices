package main

import (
	"context"
	"github.com/gofiber/websocket/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func handleWebSocket(c *websocket.Conn, pool *pgxpool.Pool) {
	defer c.Close()
	for {
		var query string
		if err := c.ReadJSON(&query); err != nil {
			log.Println("Error reading query:", err)
			break
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		conn, err := pool.Acquire(ctx)
		if err != nil {
			log.Println("Error acquiring connection:", err)
			c.WriteMessage(websocket.TextMessage, []byte("Error connecting to database"))
			continue
		}

		defer conn.Release()
		results, err := searchDatabase(query)
		if err != nil {
			log.Println("Error searchinng database:", err)
			err := c.WriteMessage(websocket.TextMessage, []byte("Error processing search query!!"))
			if err != nil {
				log.Println("Error Writting the message", err)
			}
			continue
		}
		if err := c.WriteJSON(results); err != nil {
			log.Println("Error sending results: ", err)
			break
		}
	}
}
