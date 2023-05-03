package main

import (
	"github.com/gabivega362/gabthebox/cli"
	"github.com/gabivega362/gabthebox/server"
)

func main() {
	args := cli.ParseArgs()
	server.ListenAndServe(args.Address)
}
