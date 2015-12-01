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

// A DirectoryEntry holds the information for a wz directory
type DirectoryEntry struct {
	*Entry
	subdirs []MapleDataDirectoryEntry
	files   []MapleDataFileEntry
	entries map[string]MapleDataEntry
}

// NewDirectoryEntry initializes a wz directory entry with the given data
func NewDirectoryEntry(name string, size, checksum int,
	parent MapleDataEntry) *DirectoryEntry {

	return &DirectoryEntry{
		Entry:   NewEntry(name, size, checksum, parent),
		subdirs: make([]MapleDataDirectoryEntry, 0),
		files:   make([]MapleDataFileEntry, 0),
		entries: make(map[string]MapleDataEntry),
	}
}

// EmptyDirectoryEntry initializes a zeroed wz directory entry object
func EmptyDirectoryEntry() *DirectoryEntry {
	return &DirectoryEntry{
		Entry:   NewEntry("", 0, 0, nil),
		subdirs: make([]MapleDataDirectoryEntry, 0),
		files:   make([]MapleDataFileEntry, 0),
		entries: make(map[string]MapleDataEntry),
	}
}

// AddDirectory adds a wz subdirectory to the wz directory
func (e *DirectoryEntry) AddDirectory(dir MapleDataDirectoryEntry) {
	e.subdirs = append(e.subdirs, dir)
	e.entries[dir.Name()] = dir
}

// AddFile adds a file to the wz directory
func (e *DirectoryEntry) AddFile(file MapleDataFileEntry) {
	e.files = append(e.files, file)
	e.entries[file.Name()] = file
}

// Subdirectories returns a slice of the wz subdirectories inside
// the wz directory
func (e *DirectoryEntry) Subdirectories() []MapleDataDirectoryEntry {
	return e.subdirs
}

// Files returns a slice of the wz files inside the wz directory
func (e *DirectoryEntry) Files() []MapleDataFileEntry {
	return e.files
}

// GetEntry returns the wz entry inside this folder that matches the given name
func (e *DirectoryEntry) GetEntry(name string) MapleDataEntry {
	return e.entries[name]
}
