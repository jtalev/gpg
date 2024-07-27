package main

import (
	"context"
	"fmt"
	"gpg/portal/internal/database"
	"gpg/portal/internal/localdb"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	if getenv("ENV") == "dev" {
		localDb := localdb.NewLocalDb()
		db := database.NewDb(localDb.UserRepo)
		log.Println(db)
		if err := run(ctx, getenv, db); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}
