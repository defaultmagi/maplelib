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

// A MapleDataType is an integer that describes the datatype of a wz entry
type MapleDataType int

// Possible values for MapleDataType
const (
	NONE MapleDataType = iota
	IMG_0x00
	SHORT
	INT
	FLOAT
	DOUBLE
	STRING
	EXTENDED
	PROPERTY
	CANVAS
	VECTOR
	CONVEX
	SOUND
	UOL
	UNKNOWN_TYPE
	UNKNOWN_EXTENDED_TYPE
	INVALID
)
