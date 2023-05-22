package main

import (
	"fmt"
	"os"

	"github.com/gabivega362/gabthebox/server"
	"github.com/gabivega362/gabthebox/server/config"
)

func main() {
	ctx, err := config.NewContext()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	server.ListenAndServe(ctx)
}
