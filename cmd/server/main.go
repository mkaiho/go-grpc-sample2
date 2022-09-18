package main

import (
	"flag"
	"log"
)

var (
	port = flag.Int("port", 3000, "listening port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Printf("listening port: %d", *port)
}
