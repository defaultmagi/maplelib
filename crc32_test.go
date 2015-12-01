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

package maplelib

import "testing"

func crc32test(t *testing.T, data []byte, expect uint32) {
	checksum := Crc32(0, data)
	if checksum != expect {
		t.Errorf("Crc32(0, % 02X) returned %08X, expected %08X",
			data, checksum, expect)
	}
}

func TestCrc32(t *testing.T) {
	crc32test(t, []byte{0x42, 0xE8}, 0x02153982)
	crc32test(t, []byte{0xAC, 0x76}, 0x9771DC10)
	crc32test(t, []byte{0x9E, 0x13}, 0xBC1142D6)
	crc32test(t, []byte{0xFF, 0x03}, 0x43FC0210)
	crc32test(t, []byte{0x30, 0xF2}, 0xB0C86B52)
	crc32test(t, []byte{0x7D, 0x59}, 0x369F3956)
	crc32test(t, []byte{0x35, 0xFD}, 0x1FBA369A)
	crc32test(t, []byte{0x9B, 0x5D}, 0x232472DE)
	crc32test(t, []byte{0x6C, 0x02}, 0xA70B03FF)
	crc32test(t, []byte{0x65, 0x5C}, 0xBB438A7C)
	crc32test(t, []byte{0xE6, 0x29}, 0xBC557BBB)
	crc32test(t, []byte{0x61, 0xE8}, 0x505D5077)
	crc32test(t, []byte{0xD3, 0x37}, 0x6B054D81)
	crc32test(t, []byte{0x04, 0x6C}, 0xDC8D7C5A)
	crc32test(t, []byte{0x87, 0x31}, 0x65B0D6C5)
	crc32test(t, []byte{0x69, 0xE7}, 0xE25AAE98)
	crc32test(t, []byte{0x9D, 0x93}, 0xA7F6FDEB)
	crc32test(t, []byte{0xAD, 0x34}, 0x786C56D5)
	crc32test(t, []byte{0x14, 0xE6}, 0x8A5AD171)
	crc32test(t, []byte{0x81, 0x87}, 0x27DEA9AF)
	crc32test(t, []byte{0x27, 0x5C}, 0x5AF7783F)
	crc32test(t, []byte{0xBD, 0x32}, 0x72BB8074)
	crc32test(t, []byte{0x36, 0x20}, 0x4D07A473)
	crc32test(t, []byte{0xFC, 0x0F}, 0x041BC6A7)
	crc32test(t, []byte{0xDB, 0x55}, 0x446AF32A)
	crc32test(t, []byte{0x73, 0xDE}, 0x2E4A7549)
	crc32test(t, []byte{0x72, 0xE2}, 0x1D6D4261)
	crc32test(t, []byte{0x52, 0x4B}, 0xEE28D246)
	crc32test(t, []byte{0x33, 0xC7}, 0x01D4327A)
	crc32test(t, []byte{0xB4, 0xB3}, 0x4727FFA3)
}
