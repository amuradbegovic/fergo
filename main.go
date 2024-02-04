package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	host := flag.String("host", "localhost", "Host name shown in directory listings")
	port := flag.Int("p", 70, "Port that the server listens to")
	rootdir := flag.String("d", "", "Root directory to serve content from (default is current directory)")
	logfile := flag.String("l", "", "File where log output is written to (optional)")
	ipv4 := flag.Bool("4", false, "Only use IPv4")
	ipv6 := flag.Bool("6", false, "Only use IPv6")
	flag.Parse()

	network := "tcp"
	if *ipv4 {
		network = "tcp4"
	} else if *ipv6 {
		network = "tcp6"
	}

	srv, _ := NewServer(*host, *port, network, *rootdir, *logfile)

	fmt.Println(srv.Address())
	err := srv.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
