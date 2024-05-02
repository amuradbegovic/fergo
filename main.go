package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	host := flag.String("n", "", "Host name prikazan u linkovima na lokalne resurse")
	netInterface := flag.String("i", "", "Network interface that the server binds to")
	port := flag.Int("p", 70, "Port za kojeg se veže server")
	rootdir := flag.String("d", "", "Root direktorij iz kojeg se poslužuje sadržaj (zadana vrijednost je trenutni radni directory)")
	logfile := flag.String("l", "", "Datoteka u koju se bilježe zahtjevi klijenata (opcionalno)")
	ipv4 := flag.Bool("4", false, "Koristi samo IPv4")
	ipv6 := flag.Bool("6", false, "Koristi samo IPv6")
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
