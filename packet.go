/*
   Copyright 2014 Franc[e]sco (lolisamurai@tfwno.gf)
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

// Package maplelib contains various go utilities related to MapleStory (encryption, packets, and so on)
package maplelib

import "fmt"

// A Packet is an array of bytes that contains a decrypted MapleStory packet.
// All of the numeric values are encoded in little endian.
type Packet []byte

// A PacketIterator is a slice of the packet array which is used as an iterator when reading values
type PacketIterator []byte

// A EndOfPacketError is returned when trying to read past the end of the packet
type EndOfPacketError struct {
	Bytes     int // bytes we attempted to read
	BytesLeft int // bytes left
}
func (e EndOfPacketError) Error() string {
	return fmt.Sprintf(
		"Tried to read %d bytes with %d bytes left to read.",
		e.Bytes, e.BytesLeft)
}

// NewPacket initializes an empty packet
// NOTE: do not create packets with make or new, as that will cause unexpected behaviour
func NewPacket() Packet {
	return make(Packet, 0)
}

func (p Packet) String() string {
	return fmt.Sprintf("Packet(%d): % X", len(p), []byte(p))
}

// Begin returns a packet iterator that points to the beginning of the packet
func (p Packet) Begin() PacketIterator {
	return PacketIterator(p[:])
}

// At returns a packet iterator that points at the desired position
func (p Packet) At(i int) PacketIterator {
	return PacketIterator(p[i:])
}

// Append appends raw data at the end of the packet
func (p *Packet) Append(data []byte) {
	*p = append(*p, data...)
}

// Encode1 encodes and appends a byte to the packet
func (p *Packet) Encode1(b byte) {
	*p = append(*p, b)
}

// Encode2 encodes and appends a word to the packet
func (p *Packet) Encode2(w uint16) {
	*p = append(*p,
		byte(w),
		byte(w>>8))
}

// Encode4 encodes and appends a dword to the packet
func (p *Packet) Encode4(dw uint32) {
	*p = append(*p,
		byte(dw),
		byte(dw>>8),
		byte(dw>>16),
		byte(dw>>24))
}

// Encode8 encodes and appends a qword to the packet
func (p *Packet) Encode8(qw uint64) {
	*p = append(*p,
		byte(qw),
		byte(qw>>8),
		byte(qw>>16),
		byte(qw>>24),
		byte(qw>>32),
		byte(qw>>40),
		byte(qw>>48),
		byte(qw>>56))
}

// EncodeBuffer encodes and appends a buffer to the packet using 2 bytes for the length
// followed by the data
func (p *Packet) EncodeBuffer(b []byte) {
	p.Encode2(uint16(len(b)))
	p.Append(b)
}

// EncodeString encodes and appends a string to the packet using 2 bytes for the length
// followed by the text bytes
func (p *Packet) EncodeString(str string) {
	p.EncodeBuffer([]byte(str))
}

// hasRoom checks whether the given iterator has enough room
// ahead to read the given number of bytes
func hasRoom(it PacketIterator, byteCount int) bool {
	return len(it) >= byteCount
}

// Decode1 decodes a byte at the position specified
// by the given iterator which is then incremented
func (p Packet) Decode1(it *PacketIterator) (res byte, err error) {
	slice := *it
	if !hasRoom(slice, 1) {
		err = EndOfPacketError{1, len(slice)}
		return
	}

	res = slice[0]
	*it = slice[1:]
	return
}

// Decode2 decodes a word (2 bytes) at the position specified
// by the given iterator which is then incremented
func (p Packet) Decode2(it *PacketIterator) (res uint16, err error) {
	slice := *it
	if !hasRoom(slice, 2) {
		err = EndOfPacketError{2, len(slice)}
		return
	}

	res = uint16(slice[0]) |
		uint16(slice[1])<<8
	*it = slice[2:]
	return
}

// Decode4 decodes a dword (4 bytes) at the position specified
// by the given iterator which is then incremented
func (p Packet) Decode4(it *PacketIterator) (res uint32, err error) {
	slice := *it
	if !hasRoom(slice, 4) {
		err = EndOfPacketError{4, len(slice)}
		return
	}

	res = uint32(slice[0]) |
		uint32(slice[1])<<8 |
		uint32(slice[2])<<16 |
		uint32(slice[3])<<24
	*it = slice[4:]
	return
}

// Decode8 decodes a qword (8 bytes) at the position specified
// by the given iterator which is then incremented
func (p Packet) Decode8(it *PacketIterator) (res uint64, err error) {
	slice := *it
	if !hasRoom(slice, 8) {
		err = EndOfPacketError{8, len(slice)}
		return
	}

	res = uint64(slice[0]) |
		uint64(slice[1])<<8 |
		uint64(slice[2])<<16 |
		uint64(slice[3])<<24 |
		uint64(slice[4])<<32 |
		uint64(slice[5])<<40 |
		uint64(slice[6])<<48 |
		uint64(slice[7])<<56
	*it = slice[8:]
	return
}

// DecodeBuffer decodes a buffer and returns a slice of the packet that points to the buffer
// NOTE: the returned slice is NOT a copy and any operation on it will affect the packet
func (p Packet) DecodeBuffer(it *PacketIterator) (res []byte, err error) {
	buflen, err := p.Decode2(it)
	if err != nil {
		return
	}

	slice := *it
	if !hasRoom(slice, int(buflen)) {
		err = EndOfPacketError{int(buflen), len(slice)}
		return
	}

	res = slice[:buflen]
	*it = slice[buflen:]
	return
}

// DecodeString decodes a string and returns it as a copy of the data
func (p Packet) DecodeString(it *PacketIterator) (res string, err error) {
	bytes, err := p.DecodeBuffer(it)
	res = string(bytes[:])
	return
}
