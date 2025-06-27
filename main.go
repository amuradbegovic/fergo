package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	host := flag.String("n", "", "Host name to show in menu entries ponting to local resources")
	netInterface := flag.String("i", "", "Network interface to bind to")
	port := flag.Int("p", 70, "TCP port to listen")
	rootdir := flag.String("d", "", "Directory to serve content from (default is the current working directory)")
	logfile := flag.String("l", "", "Log file (optional, logs to stderr by default)")
	ipv4 := flag.Bool("4", false, "IPv4 only")
	ipv6 := flag.Bool("6", false, "IPv6 only")
	flag.Parse()

	network := "tcp"
	if *ipv4 {
		network = "tcp4"
	} else if *ipv6 {
		network = "tcp6"
	}

	srv, _ := NewServer(*host, *netInterface, *port, network, *rootdir, *logfile)

	fmt.Println(srv.Address())
	err := srv.Serve()

	if err != nil {
		log.Fatal(err)
	}

}
