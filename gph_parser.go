// Parser za .gph format menija (kao u geomyidae)

package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func SplitGPHLine(line string) []string {
	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")

	var fields []string
	currentField := ""
	for i := 0; i < len((line)); i++ {
		if line[i] == '|' {
			if i > 0 {
				if line[i-1] == '\\' {
					currentField += string(line[i])
					continue
				}
			}
			fields = append(fields, currentField)
			currentField = ""
		} else {
			currentField += string(line[i])
		}
	}
	fields = append(fields, currentField)

	return fields
}

func GPHLineToMenuItem(line, path string, srv Server) (MenuItem, error) {
	var mitem MenuItem

	if !(strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")) {
		return mitem, errors.New("link mora biti okružen uglatim zagradama")
	}

	fields := SplitGPHLine(line)
	if len(fields) != 5 {
		return mitem, errors.New("link mora imati 5 polja")
	}
	if !(len(fields[0]) == 1) {
		return mitem, errors.New("prvo polje mora sadržati samo jedan znak")
	}
	mitem.Type = fields[0][0]
	mitem.DisplayName = fields[1]
	mitem.Selector = fields[2]
	if !strings.HasPrefix(mitem.Selector, "/") && mitem.Selector != "Err" {
		mitem.Selector = "/" + srv.RelPath(filepath.Dir(path)) + "/" + mitem.Selector
	}
	if fields[3] == "server" {
		mitem.Host = srv.Host
	} else {
		mitem.Host = fields[3]
	}
	if fields[4] == "port" || fields[4] == "" {
		mitem.Port = srv.Port
	} else {
		portnumber, err := strconv.Atoi(fields[4])
		if err != nil {
			return mitem, err
		}
		mitem.Port = portnumber
	}

	return mitem, nil
}

func ParseGPHFile(path string, srv Server) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	menu := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mitem, err := GPHLineToMenuItem(line, path, srv)
		if err != nil {
			mitem = MenuItem{'i', strings.TrimPrefix(line, "[|"), "", srv.Host, srv.Port}
		}

		menu += mitem.String()
	}
	return menu + "\n.\n", nil
}
