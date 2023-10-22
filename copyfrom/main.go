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
	conn, err := pgx.Connect(ctx, connURL)
	if err != nil {
		log.Fatalf("pgx connect: %v", err)
	}
	defer conn.Close(ctx)

	jstZone := time.FixedZone("Asia/Tokyo", 9*60*60)

	args := [][]any{
		{"0001", time.Now().UTC()},
		{"0002", time.Now().In(jstZone)},
	}

	copyCount, err := conn.CopyFrom(ctx, pgx.Identifier{"event"},
		[]string{"event_id", "event_at"}, pgx.CopyFromRows(args))

	if err != nil {
		log.Fatalf("copy exec: %v", err)
	}

	fmt.Printf("copy: %d\n", copyCount)

}
