package main

import (
	"flag"
	"log"
)

var (
	addr = flag.String("addr", "localhost:3000", "server address")
)

func init() {
	flag.Parse()
}

func main() {
	log.Printf("server address: %s", *addr)
}
