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
// TODO: reimplement this more efficiently

// An Entry is a generic wz entry and holds data common to all
// types of wz entries
type Entry struct {
	name     string
	size     int // not used in wz xmls
	checksum int // not used in wz xmls
	offset   int // not used in wz xmls
	parent   MapleDataEntity
}

// NewEntry initializes a generic wz entry object
// NOTE: esize and echecksum are not used in wx xml reading and can be
// left zeroed
func NewEntry(ename string, esize, echecksum int, eparent MapleDataEntity,
) *Entry {

	return &Entry{
		name:     ename,
		size:     esize,
		checksum: echecksum,
		parent:   eparent,
	}
}

func (e *Entry) Name() string            { return e.name }
func (e *Entry) Size() int               { return e.size }
func (e *Entry) Checksum() int           { return e.checksum }
func (e *Entry) Offset() int             { return e.offset }
func (e *Entry) Parent() MapleDataEntity { return e.parent }
