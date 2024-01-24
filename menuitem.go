package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type MenuItem struct {
	Type        byte
	DisplayName string
	Selector    string
	Host        string
	Port        int
}

func (dirent MenuItem) String() string {
	return fmt.Sprintf("%s%s\t%s\t%s\t%d\n", string(dirent.Type), dirent.DisplayName, dirent.Selector, dirent.Host, dirent.Port)
}

func NewFromDirEntry(dirent os.DirEntry, path string, srv Server) MenuItem {
	var result MenuItem

	result.DisplayName = dirent.Name()

	relpath, _ := filepath.Rel(srv.Rootdir, path)

	if relpath != "." {
		result.Selector = "/" + relpath
	}
	result.Selector += "/" + dirent.Name()

	var err error
	result.Type, err = MenuItemType(dirent, path)
	if err != nil {
		log.Fatal(err)
	}
	result.Host = srv.Host
	result.Port = srv.Port
	return result
}
