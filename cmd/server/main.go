package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}