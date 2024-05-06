package main

import (
	"os"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func MIMETypeToMenuItemType(detectedMIME *mimetype.MIME) byte {
	for mtype := detectedMIME; mtype != nil; mtype = mtype.Parent() {
		if mtype.Is("text/plain") {
			return '0'
		}
		if strings.Contains(mtype.String(), "binhex") {
			return '4'
		}
		if strings.Contains(mtype.String(), "image") {
			if strings.Contains(mtype.String(), "image/gif") {
				return 'g'
			} else {
				return 'I'
			}
		}
	}
	return '9'
}

func MenuItemType(dirent os.DirEntry, path string) byte {
	if dirent.IsDir() || strings.HasSuffix(dirent.Name(), ".gph") {
		return '1'
	}
	detectedMIME, err := mimetype.DetectFile(path + "/" + dirent.Name())
	if err != nil {
		return '0'
	}
	return MIMETypeToMenuItemType(detectedMIME)
}
