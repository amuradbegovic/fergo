package main

import (
	"fmt"
	"os"
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

	relpath := srv.RelPath(path)

	if relpath != "." {
		result.Selector = "/" + relpath
	}
	result.Selector += "/" + dirent.Name()

	result.Type = MenuItemType(dirent, path)

	result.Host = srv.Host
	result.Port = srv.Port
	return result
}
