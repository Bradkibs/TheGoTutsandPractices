package main

import (
	"github.com/gofiber/websocket/v2"
	"log"
)

func handleWebSocket(c *websocket.Conn) {
	defer c.Close()
	for {
		var query string
		if err := c.ReadJSON(&query); err != nil {
			log.Println("Error reading query:", err)
			break
		}

		results, err := searchDatabase(query)
		if err != nil {
			log.Println("Error searchinng database:", err)
			c.WriteMessage(websocket.TextMessage, []byte("Error processing search query!!"))
			continue
		}
		if err := c.WriteJSON(results); err != nil {
			log.Println("Error sending results: ", err)
			break
		}
	}
}
