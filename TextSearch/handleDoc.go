package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Documents struct {
	Docs []Document `json:"docs"`
}

func handleAddDocs(c *fiber.Ctx, pool *pgxpool.Pool) error {
	var docs Documents
	if err := c.BodyParser(&docs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request format!!"})
	}
	if len(docs.Docs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Documents required!!"})
	}
	ctx := context.Background()
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "System failure, Please wait 1 second and try again."})
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to begin transaction"})
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO documents (title, content, tsv) VALUES ($1, $2, to_tsvector($1 || ' ' || $2) )`

	for _, doc := range docs.Docs {
		_, err := tx.Exec(ctx, query, doc.Title, doc.Content)
		if err != nil {
			log.Println("Error inserting the document:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "error", "message": "Failed to insert the documents!!",
			})

		}
	}
	if err := tx.Commit(ctx); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Failed to commit transaction"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "Documents added successfully!"})
}
