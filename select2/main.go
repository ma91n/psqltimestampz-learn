package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func main() {
	ctx := context.Background()

	connURL := "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable&timezone=japan"
	conn, err := sqlx.Connect("pgx", connURL)
	if err != nil {
		log.Fatalf("pgx connect: %v", err)
	}
	defer conn.Close()

	var events []Event
	if err := conn.SelectContext(ctx, &events, "select * from event"); err != nil {
		log.Fatalf("select query: %v", err)
	}

	for _, e := range events {
		//fmt.Printf("%s %s\n", e.ID, e.At.In(jst))
		fmt.Printf("%s %s\n", e.ID, e.At)
	}

}
