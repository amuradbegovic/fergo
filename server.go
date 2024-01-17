package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Server struct {
	Host    string
	Port    int
	Rootdir string
}

func NewServer(host string, port int, rootdir string) (Server, error) {
	var srv Server
	if host == "" {
		host = "localhost"
	}
	srv.Host = host
	if rootdir == "" {
		var err error
		rootdir, err = os.Getwd()
		if err != nil {
			return srv, err
		}
	}
	srv.Port = port
	srv.Rootdir = rootdir
	return srv, nil
}

func (srv Server) Address() string {
	return srv.Host + ":" + fmt.Sprintf("%d", srv.Port)
}

func (srv Server) Serve() error {
	lsn, err := net.Listen("tcp", srv.Address())
	if err != nil {
		return err
	}
	defer lsn.Close()
	for {
		conn, err := lsn.Accept()
		if err != nil {
			return err
		}

		go srv.HandleConnection(conn)
	}
}

func ServeFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (srv Server) ServeDir(path string) (string, error) {
	menu := ""
	files, err := os.ReadDir(path)
	if err != nil {
		return menu, err
	}

	relpath, _ := filepath.Rel(srv.Rootdir, path)
	for _, file := range files {
		menu = fmt.Sprintf("%s\n%s", menu, NewFromDirEntry(file, relpath, srv.Host, srv.Port).String())
	}
	menu += "\n.\n"
	return menu, nil
}

func (srv Server) HandleConnection(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)
	nread, _ := conn.Read(request)
	selector := string(request[:nread])
	selector = strings.TrimSuffix(selector, "\r\n")
	selector = strings.TrimSuffix(selector, "\n")

	response := ""

	log.Printf("%s\t%s", selector, conn.RemoteAddr().String())

	path := srv.Rootdir + selector
	fmt.Printf("Requested path: \"%s\"\n", path)

	fileinfo, err := os.Stat(path)
	if err != nil {
		response = fmt.Sprintln(err, "\n.")
	} else {
		if fileinfo.IsDir() {
			response, err = srv.ServeDir(path)
		} else {
			response, err = ServeFile(path)
		}
		if err != nil {
			response = fmt.Sprintln(err, "\n.")
		}
	}

	conn.Write([]byte(response))

	//c.Write([]byte(fmt.Sprintln("iWelcome to port 70\n.")))

}
