package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()		//add user data here dummy
	if err := run(ctx, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}