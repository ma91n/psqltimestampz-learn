package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	ctx := context.Background()

	connURL := "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	conn, err := sqlx.Connect("pgx", connURL)
	if err != nil {
		log.Fatalf("pgx connect: %v", err)
	}
	defer conn.Close()

	var events []Event
	//if err := conn.SelectContext(ctx, &events, "select *, to_char(event_at, 'yyyy-mm-dd HH24:MI:SS') as event_time from event"); err != nil {
	if err := conn.SelectContext(ctx, &events, "select * from event"); err != nil {
		log.Fatalf("select query: %v", err)
	}

	for _, e := range events {
		fmt.Printf("event: %+v\n", e)
	}

}
