/*
   Copyright 2014-2015 Franc[e]sco (lolisamurai@tfwno.gf)
   This file is part of maplelib-go.
   maplelib-go is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   maplelib-go is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with maplelib-go. If not, see <http://www.gnu.org/licenses/>.
*/

// Package wz contains various utilities to parse Wz data
// 90% of this package is ported directly from OdinMS, so credits to them
package wz

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const xmldebug = false

// wz.Xml provides access to the data in a wz xml directory tree
type Xml struct {
	root              string
	rootForNavigation *DirectoryEntry
}

// NewXml walks the given root directory and creates a new wx.Xml object
// that will provide access to the wz xml data
func NewXml(sourcedirpath string) (*Xml, error) {
	x := &Xml{
		root: sourcedirpath,
		rootForNavigation: NewDirectoryEntry(
			filepath.Base(sourcedirpath), 0, 0, nil),
	}

	// walk sourcedir and fill all data entities
	err := fillMapleDataEntities(x.root, x.rootForNavigation)
	return x, err
}

// fillMapleDataEntities is a recursive function that walks the root directory
// of the xml wz files and caches everything into the wzdir object
func fillMapleDataEntities(lrootpath string, wzdir *DirectoryEntry) (
	err error) {

	lroot, err := os.Open(lrootpath)
	if err != nil {
		return
	}
	defer lroot.Close()

	lrootInfo, err := lroot.Stat()
	if err != nil {
		return
	}

	if !lrootInfo.IsDir() {
		err = errors.New("The root directory must be a directory, not a file!")
		return
	}

	// enumerates all files and dirs in lroot
	lrootFilesAndDirs, err := lroot.Readdir(-1)
	if err != nil {
		return
	}

	// iterate all diles and subdirs in the current directory
	for _, fileInfo := range lrootFilesAndDirs {
		fileName := fileInfo.Name()
		if xmldebug {
			fmt.Println(fileName)
		}
		switch {
		// found a directory, walk into it recursively
		case fileInfo.IsDir() && !strings.HasSuffix(fileName, ".img"):
			newdir := NewDirectoryEntry(fileName, 0, 0, wzdir)
			wzdir.AddDirectory(newdir)

			// recursively walk it
			err = fillMapleDataEntities(
				filepath.Join(lrootpath, fileName), newdir)
			if err != nil {
				return
			}

		// found a file, add it
		case strings.HasSuffix(fileName, ".xml"):
			wzdir.AddFile(
				NewFileEntry(fileName[0:len(fileName)-4], 0, 0, wzdir))
		}
	}

	return
}

// Get returns the wz data at the given path
func (x *Xml) Get(path string) (res MapleData, err error) {
	dataFile, err := os.Open(filepath.Join(x.root, path+".xml"))
	if err != nil {
		return
	}
	defer dataFile.Close()

	res, err = NewXMLDomMapleData(dataFile, path)
	return
}

// Root returns the root directory entry of the xml file
func (x *Xml) Root() MapleDataDirectoryEntry { return x.rootForNavigation }
