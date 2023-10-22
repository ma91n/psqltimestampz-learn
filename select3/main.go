package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	connURL := "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	config, err := pgx.ParseConfig(connURL)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("pgx connect: %v", err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, "select * from event")
	if err != nil {
		log.Fatalf("select query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.At); err != nil {
			log.Fatalf("scan: %v", err)
		}
		fmt.Printf("%s %s\n", e.ID, e.At.Format(time.RFC3339))
	}

}
