package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/azizurrahaman/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := db.New(conn)

	// Example: Create a new account
	account, err := queries.CreateAccount(context.Background(), db.CreateAccountParams{
		Owner:    "john_doe",
		Balance:  1000,
		Currency: "USD",
	})

	if err != nil {
		log.Fatal("cannot create account:", err)
	}

	fmt.Printf("Created account: %+v\n", account)

	// Example: Get the account back
	retrievedAccount, err := queries.GetAccount(context.Background(), account.ID)
	if err != nil {
		log.Fatal("cannot get account:", err)
	}

	fmt.Printf("Retrieved account: %+v\n", retrievedAccount)
}
