package main

import (
	"fmt"
	"os"
	"strings"
)

type DirEntity struct {
	Type        int
	DisplayName string
	Selector    string
	Host        string
	Port        int
}

func (dirent DirEntity) String() string {
	return fmt.Sprintf("%d%s\t%s\t%s\t%d\n", dirent.Type, dirent.DisplayName, dirent.Selector, dirent.Host, dirent.Port)
}

func NewFromDirEntry(dirent os.DirEntry, relPath, host string, port int) DirEntity {
	var result DirEntity
	if dirent.IsDir() || strings.HasSuffix(dirent.Name(), ".gph") || strings.HasSuffix(dirent.Name(), ".gophermap") {
		result.Type = 1
	} else {
		result.Type = 0
	}
	result.DisplayName = dirent.Name()
	result.Selector = "/" + relPath + "/" + dirent.Name()
	result.Host = host
	result.Port = port
	return result
}
