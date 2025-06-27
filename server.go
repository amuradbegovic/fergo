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
	Host      string
	Interface string
	Port      int
	Network   string
	RootDir   string
	LogPath   string
}

func NewServer(host, netInterface string, port int, network, rootdir, logpath string) (Server, error) {
	if host == "" {
		host = "localhost"
		if netInterface != "" {
			host = netInterface
		}
	}
	if rootdir == "" {
		var err error
		rootdir, err = os.Getwd()
		if err != nil {
			return Server{}, err
		}
	}
	return Server{host, netInterface, port, network, rootdir, logpath}, nil
}

func (srv Server) Address() string {
	return srv.Interface + ":" + fmt.Sprintf("%d", srv.Port)
}

func (srv Server) RelPath(path string) string {
	relpath, _ := filepath.Rel(srv.RootDir, path)
	return relpath
}

func (srv Server) Serve() error {
	if srv.LogPath != "" {
		f, err := os.OpenFile(srv.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("Error opening file: %v. Logging to stdout.", err)
		} else {
			defer f.Close()
			log.SetOutput(f)
		}
	}

	lsn, err := net.Listen(srv.Network, srv.Address())
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

func (srv Server) ServeFile(path string) (string, error) {
	if strings.HasSuffix(path, ".gph") {
		return ParseGPHFile(path, srv)
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (srv Server) ServeDir(path string) (string, error) {
	indexFile, err := ParseGPHFile(path+"/index.gph", srv)
	if err == nil {
		return indexFile, nil
	}

	menu := ""
	files, err := os.ReadDir(path)
	if err != nil {
		return menu, err
	}

	for _, file := range files {
		menu += NewFromDirEntry(file, path, srv).String()
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
	if tabIndex := strings.Index(selector, "\t"); tabIndex != -1 {
		selector = selector[:tabIndex]
	}

	if !strings.HasPrefix(selector, "/") {
		selector = "/" + selector
	}

	log.Printf("%s\t%s", conn.RemoteAddr().String(), selector)

	response := ""

	if strings.Contains(selector, "..") {
		response = "Error: selector cannot contain \"..\"\n"
	} else {
		path := srv.RootDir + selector

		fileinfo, err := os.Stat(path)
		if err != nil {
			response = "Error: resource not found\n"
		} else {
			if fileinfo.IsDir() {
				response, _ = srv.ServeDir(path)
			} else {
				response, _ = srv.ServeFile(path)
			}
		}
	}
	conn.Write([]byte(response))
}
