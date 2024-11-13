package main

import (
	"context"
	"github.com/gofiber/websocket/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

// Message defines the structure of a WebSocket message.
type Message struct {
	Action  string `json:"action"`  // The action the client wants to perform, e.g., "search"
	Query   string `json:"query"`   // The query string
	Message string `json:"message"` // Additional message or status information
}

// Result defines the structure of the result message sent back to the client.
type Result struct {
	Status  string      `json:"status"`  // Success or error status
	Payload interface{} `json:"payload"` // Result data
	Message string      `json:"message"` // Any additional status message
}

func handleWebSocket(c *websocket.Conn, pool *pgxpool.Pool) {
	defer c.Close()
	for {
		var msg Message
		if err := c.ReadJSON(&msg); err != nil {
			log.Println("Error reading JSON:", err)
			c.WriteJSON(Result{
				Status:  "error",
				Message: "Invalid JSON format",
			})
			break
		}

		switch msg.Action {
		case "search":
			handleSearchAction(c, pool, msg.Query)
		default:
			c.WriteJSON(Result{
				Status:  "error",
				Message: "Unknown action",
			})
		}

	}
}

func handleSearchAction(c *websocket.Conn, pool *pgxpool.Pool, query string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Println("Error acquiring connection: ", err)
		c.WriteJSON(Result{Status: "error", Message: "Error connecting to the database"})
		return
	}
	defer conn.Release()
	results, err := searchDatabase(query)

	if err != nil {
		log.Println("Error searching database:", err)
		c.WriteJSON(Result{
			Status:  "error",
			Message: "Error processing search query",
		})
		return

	}
	if len(results) == 0 {
		c.WriteJSON(Result{
			Status:  "success",
			Payload: "No match for the searched item!!",
			Message: "Search completed"})
		log.Println("No results found")
		return
	}

	if err := c.WriteJSON(Result{
		Status:  "success",
		Payload: results,
		Message: "Search completed",
	}); err != nil {
		log.Println("Error sending results:", err)
	}
}
