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

// A FileEntry holds the information for a wz file entry
type FileEntry struct {
	*Entry
	feoffset int // not used in xml wz, todo: check if we can use Entry's offset
}

// NewFileEntry initializes a new wz file entry object
// NOTE: size and checksum are not used in wz xml parsing and can be left zeroed
func NewFileEntry(name string, size, checksum int, parent MapleDataEntity,
) *FileEntry {

	return &FileEntry{
		Entry: NewEntry(name, size, checksum, parent),
	}
}

func (e *FileEntry) Offset() int {
	return e.feoffset
}

func (e *FileEntry) SetOffset(offset int) {
	e.feoffset = offset
}
