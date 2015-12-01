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

package wz

// 90% of this package is ported directly from OdinMS, so credits to them

import (
	"errors"
	"os"
)

// A MapleDataProvider is a generic interface for an object that
// parses or provides wz data
type MapleDataProvider interface {
	Get(path string) (MapleData, error)
	Root() MapleDataDirectoryEntry
}

// NewMapleDataProvider analyzes the given path and provides the appropriate
// MapleDataProvider for the format if supported. If the format is not supported
// it will return nil.
// At the moment, only wz xml files are supported.
func NewMapleDataProvider(path string) (res MapleDataProvider, err error) {
	res = nil
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if !fi.IsDir() { // TODO: support Wz files
		err = errors.New("Wz files are not supported yet, please use wz xml's")
		return
	}

	return NewXml(path)
}
