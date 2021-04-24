package main

import (
	"flag"

	"github.com/thangnd85/plaincast/server"
)

func main() {
	flag.Parse()

	server.Serve()
}
