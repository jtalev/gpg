package main

import (
	"context"
	"fmt"
	"gpg/portal/internal/localdb"
	"os"
)

func main() {
	ctx := context.Background()
	if getenv("ENV") == "dev" {
		db := localdb.Db{}
		if err := run(ctx, getenv, db); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}