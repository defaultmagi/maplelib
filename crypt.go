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

package maplelib

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

/*
   An encryption key for MapleStory packets.
   It consists of a 4-byte value repeated four times for a total of 16 bytes.

   MapleStory uses two keys: one is used to encrypt sent packets and the other
   is used to decrypt received packets. This is valid for both the client and
   the server.
   The two keys are randomly generated when a client connects and sent
   in the unencrypted handshake packet which is the first packet that is sent
   to a client that connects to a server.
   The MapleStory protocol will shift encryption keys every time a packet is
   send or received, so you will have to call .Shuffle() every time this
   happens.
*/
type Crypt struct {
	mapleVersion uint16
	Key          [16]byte
}

const encryptedHeaderSize = 4
const blocksize = 1460

var aeskey = [32]byte{0x13, 0x00, 0x00, 0x00,
	0x08, 0x00, 0x00, 0x00,
	0x06, 0x00, 0x00, 0x00,
	0xB4, 0x00, 0x00, 0x00,
	0x1B, 0x00, 0x00, 0x00,
	0x0F, 0x00, 0x00, 0x00,
	0x33, 0x00, 0x00, 0x00,
	0x52, 0x00, 0x00, 0x00}

// Initializes and returns encryption key
func NewCrypt(key [4]byte, mapleVersion uint16) Crypt {
	var res Crypt

	// Repeats the key 4 times
	for i := 0; i < 4; i++ {
		copy(res.Key[4*i:], key[:])
	}

	res.mapleVersion = encodeMapleVersion(mapleVersion)
	return res
}

func (c Crypt) String() string {
	return fmt.Sprintf("MapleCrypt[v%d]{% X}", c.GetMapleVersion(), c.Key)
}

// Returns the target MapleStory version for this encryption key
func (c *Crypt) GetMapleVersion() uint16 {
	return decodeMapleVersion(c.mapleVersion)
}

// Encrypts the given array of bytes.
// NOTE: the array must have 4 bytes of space at the beginning for the encrypted header
func (c *Crypt) Encrypt(buffer []byte) {
	c.makeHeader(buffer)
	mapleCrypt(buffer[encryptedHeaderSize:])
	c.aesCrypt(buffer[encryptedHeaderSize:])
}

// Decrypts the given array of bytes.
// NOTE: you must omit the first 4 bytes (encrypted header)
func (c *Crypt) Decrypt(buffer []byte) {
	c.aesDecrypt(buffer[:])
	mapleDecrypt(buffer[:])
}

func (c *Crypt) makeHeader(buffer []byte) {
	cb := uint16(len(buffer) - encryptedHeaderSize)

	// I have no idea what I'm doing
	// this encryption shit was reversed from the game itself
	iiv := uint16(c.Key[3] & 0xFF) // is the & 0xFF even needed?
	iiv |= uint16(c.Key[2]) << 8 & 0xFF00

	iiv ^= c.mapleVersion
	mlength := uint16((cb << 8 & 0xFF00) | cb>>8)
	xoredIv := iiv ^ mlength

	buffer[0] = byte(iiv >> 8 & 0xFF)
	buffer[1] = byte(iiv & 0xFF)
	buffer[2] = byte(xoredIv >> 8 & 0xFF)
	buffer[3] = byte(xoredIv & 0xFF)
}

// Decodes the packet length from the encrypted header.
// NOTE: this size does not include the 4-byte encrypted header.
func GetPacketLength(encryptedHeader []byte) int {
	return int((uint16(encryptedHeader[0]) + uint16(encryptedHeader[1])*0x100) ^
		(uint16(encryptedHeader[2]) + uint16(encryptedHeader[3])*0x100))
}

func encodeMapleVersion(mapleVersion uint16) (res uint16) {
	// I dunno why they do this, looks like some kind of way to encrypt the game version
	res = 0xFFFF - mapleVersion
	res = (res >> 8 & 0xFF) | (res << 8 & 0xFF00)
	return
}

func decodeMapleVersion(mapleVersion uint16) (res uint16) {
	res = (mapleVersion >> 8 & 0xFF) | (mapleVersion << 8 & 0xFF00)
	res = -(res - 0xFFFF)
	return
}

// Shuffles the current key after an encryption or decryption
func (c *Crypt) Shuffle() {
	// I have no idea what I'm doing
	// this encryption shit was reversed from the game itself
	// credits to vana for the key shuffle code
	var im12andwhatisthis = [256]byte{
		0xEC, 0x3F, 0x77, 0xA4, 0x45, 0xD0, 0x71, 0xBF, 0xB7, 0x98, 0x20, 0xFC, 0x4B, 0xE9, 0xB3, 0xE1,
		0x5C, 0x22, 0xF7, 0x0C, 0x44, 0x1B, 0x81, 0xBD, 0x63, 0x8D, 0xD4, 0xC3, 0xF2, 0x10, 0x19, 0xE0,
		0xFB, 0xA1, 0x6E, 0x66, 0xEA, 0xAE, 0xD6, 0xCE, 0x06, 0x18, 0x4E, 0xEB, 0x78, 0x95, 0xDB, 0xBA,
		0xB6, 0x42, 0x7A, 0x2A, 0x83, 0x0B, 0x54, 0x67, 0x6D, 0xE8, 0x65, 0xE7, 0x2F, 0x07, 0xF3, 0xAA,
		0x27, 0x7B, 0x85, 0xB0, 0x26, 0xFD, 0x8B, 0xA9, 0xFA, 0xBE, 0xA8, 0xD7, 0xCB, 0xCC, 0x92, 0xDA,
		0xF9, 0x93, 0x60, 0x2D, 0xDD, 0xD2, 0xA2, 0x9B, 0x39, 0x5F, 0x82, 0x21, 0x4C, 0x69, 0xF8, 0x31,
		0x87, 0xEE, 0x8E, 0xAD, 0x8C, 0x6A, 0xBC, 0xB5, 0x6B, 0x59, 0x13, 0xF1, 0x04, 0x00, 0xF6, 0x5A,
		0x35, 0x79, 0x48, 0x8F, 0x15, 0xCD, 0x97, 0x57, 0x12, 0x3E, 0x37, 0xFF, 0x9D, 0x4F, 0x51, 0xF5,
		0xA3, 0x70, 0xBB, 0x14, 0x75, 0xC2, 0xB8, 0x72, 0xC0, 0xED, 0x7D, 0x68, 0xC9, 0x2E, 0x0D, 0x62,
		0x46, 0x17, 0x11, 0x4D, 0x6C, 0xC4, 0x7E, 0x53, 0xC1, 0x25, 0xC7, 0x9A, 0x1C, 0x88, 0x58, 0x2C,
		0x89, 0xDC, 0x02, 0x64, 0x40, 0x01, 0x5D, 0x38, 0xA5, 0xE2, 0xAF, 0x55, 0xD5, 0xEF, 0x1A, 0x7C,
		0xA7, 0x5B, 0xA6, 0x6F, 0x86, 0x9F, 0x73, 0xE6, 0x0A, 0xDE, 0x2B, 0x99, 0x4A, 0x47, 0x9C, 0xDF,
		0x09, 0x76, 0x9E, 0x30, 0x0E, 0xE4, 0xB2, 0x94, 0xA0, 0x3B, 0x34, 0x1D, 0x28, 0x0F, 0x36, 0xE3,
		0x23, 0xB4, 0x03, 0xD8, 0x90, 0xC8, 0x3C, 0xFE, 0x5E, 0x32, 0x24, 0x50, 0x1F, 0x3A, 0x43, 0x8A,
		0x96, 0x41, 0x74, 0xAC, 0x52, 0x33, 0xF0, 0xD9, 0x29, 0x80, 0xB1, 0x16, 0xD3, 0xAB, 0x91, 0xB9,
		0x84, 0x7F, 0x61, 0x1E, 0xCF, 0xC5, 0xD1, 0x56, 0x3D, 0xCA, 0xF4, 0x05, 0xC6, 0xE5, 0x08, 0x49}

	newiv := [4]byte{0xF2, 0x53, 0x50, 0xC6}
	var input, valueinput byte
	var fulliv, shift uint32

	for i := byte(0); i < 4; i++ {
		input = c.Key[i]
		valueinput = im12andwhatisthis[input]

		newiv[0] += im12andwhatisthis[newiv[1]] - input
		newiv[1] -= newiv[2] ^ valueinput
		newiv[2] ^= im12andwhatisthis[newiv[3]] + input
		newiv[3] -= newiv[0] - valueinput

		fulliv = uint32(newiv[3]<<24) | uint32(newiv[2]<<16) | uint32(newiv[1]<<8) | uint32(newiv[0])
		shift = fulliv>>0x1D | fulliv<<0x03

		newiv[0] = byte(shift & uint32(0xFF))
		newiv[1] = byte(shift >> 8 & uint32(0xFF))
		newiv[2] = byte(shift >> 16 & uint32(0xFF))
		newiv[3] = byte(shift >> 24 & uint32(0xFF))
	}

	for i := byte(0); i < 4; i++ {
		copy(c.Key[4*i:], newiv[:])
	}
}

func ror(val byte, num int) byte {
	for i := 0; i < num; i++ {
		var lowbit int

		if val&1 > 0 {
			lowbit = 1
		} else {
			lowbit = 0
		}

		val >>= 1
		val |= byte(lowbit << 7)
	}

	return val
}

func rol(val byte, num int) byte {
	var highbit int

	for i := 0; i < num; i++ {
		if val&0x80 > 0 {
			highbit = 1
		} else {
			highbit = 0
		}

		val <<= 1
		val |= byte(highbit)
	}

	return val
}

func mapleDecrypt(buf []byte) {
	// I have no idea what I'm doing
	// this encryption shit was reversed from the game itself
	var j int32
	var a, b, c byte

	for i := byte(0); i < 3; i++ {
		a = 0
		b = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[j-1]
			c = rol(c, 3)
			c ^= 0x13
			a = c
			c ^= b
			c = byte(int32(c) - j)
			c = ror(c, 4)
			b = a
			buf[j-1] = c
		}

		a = 0
		b = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[int32(len(buf))-j]
			c -= 0x48
			c ^= 0xFF
			c = rol(c, int(j))
			a = c
			c ^= b
			c = byte(int32(c) - j)
			c = ror(c, 3)
			b = a
			buf[int32(len(buf))-j] = c
		}
	}
}

func mapleCrypt(buf []byte) {
	// I have no idea what I'm doing
	// this encryption shit was reversed from the game itself
	var j int32
	var a, c byte

	for i := byte(0); i < 3; i++ {
		a = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[int32(len(buf))-j]
			c = rol(c, 3)
			c = byte(int32(c) + j)
			c ^= a
			a = c
			c = ror(a, int(j))
			c ^= 0xFF
			c += 0x48
			buf[int32(len(buf))-j] = c
		}

		a = 0

		for j = int32(len(buf)); j > 0; j-- {
			c = buf[j-1]
			c = rol(c, 4)
			c = byte(int32(c) + j)
			c ^= a
			a = c
			c ^= 0x13
			c = ror(c, 3)
			buf[j-1] = c
		}
	}
}

func (c *Crypt) aesCrypt(buf []byte) {
	var pos, tpos, cbwrite, cb int32 = 0, 0, 0, int32(len(buf))
	var first byte = 1

	cb = int32(len(buf))

	// I'm not 100% sure what this exactly does but apparently maple
	// decrypts packets in blocks of 1460 bytes to work around
	// packet limitations or something

	for cb > pos {
		tpos = blocksize - int32(first*4)

		if cb > pos+tpos {
			cbwrite = tpos
		} else {
			cbwrite = cb - pos
		}

		block, err := aes.NewCipher(aeskey[:])
		if err != nil {
			panic(err) // cbf to handle this unlikely error
		}

		stream := cipher.NewOFB(block, c.Key[:])
		stream.XORKeyStream(buf[pos:pos+cbwrite], buf[pos:pos+cbwrite])

		pos += tpos

		if first == 1 {
			first = 0
		}
	}
}

func (c *Crypt) aesDecrypt(buf []byte) {
	var pos, tpos, cbread, cb int32 = 0, 0, 0, int32(len(buf))
	var first byte = 1

	for cb > pos {
		tpos = blocksize - int32(first*4)

		if cb > pos+tpos {
			cbread = tpos
		} else {
			cbread = cb - pos
		}

		block, err := aes.NewCipher(aeskey[:])
		if err != nil {
			panic(err) // cbf to handle this unlikely error
		}

		stream := cipher.NewOFB(block, c.Key[:])
		stream.XORKeyStream(buf[pos:pos+cbread], buf[pos:pos+cbread])

		pos += tpos

		if first == 1 {
			first = 0
		}
	}
}
