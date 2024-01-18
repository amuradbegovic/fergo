package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	host := flag.String("h", "localhost", "Host shown in directory listings")
	port := flag.Int("p", 70, "Port that the server listens to")
	rootdir := flag.String("d", "", "Root directory to serve content from (default is current directory)")
	flag.Parse()
	srv, _ := NewServer(*host, *port, *rootdir)

	fmt.Println(srv.Address())
	err := srv.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
