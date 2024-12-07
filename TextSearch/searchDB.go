package main

import (
	"context"
	"fmt"
)

type Doc struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func searchDatabase(query string) ([]Doc, error) {
	searchQuery := fmt.Sprintf("%s:*", query)

	sql := `SELECT id, title, content FROM documents WHERE tsv @@ to_tsquery('english', $1) ORDER BY ts_rank_cd(tsv,to_tsquery('english',$1));`

	rows, err := pool.Query(context.Background(), sql, searchQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Doc
	for rows.Next() {
		var doc Doc
		if err := rows.Scan(&doc.ID, &doc.Title, &doc.Content); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}
