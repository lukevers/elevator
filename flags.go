package main

import (
	"flag"
)

var (
	host = flag.String("host", "127.0.0.1", "Host to bind webserver to")
	port = flag.Int("port", 2999, "Port to bind webserver to")
)
