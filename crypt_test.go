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

import (
	"bytes"
	"testing"
)

func TestEncryption(t *testing.T) {
	packet := NewPacket()
	packet.Encode4(0x00000000) // placeholder for the encrypted header
	packet.Encode2(0x000D)
	packet.Encode4(0xBAADF00D)
	iv := [4]byte{0xFE, 0xCA, 0xDD, 0xBA}

	pcopy := make([]byte, len(packet))
	copy(pcopy, []byte(packet))
	crypt := NewCrypt(iv, 62)
	crypt.Encrypt(pcopy)
	packetlen := GetPacketLength(pcopy)

	if packetlen != len(packet[4:]) {
		t.Errorf("Encrypted packet length = %d, expected %d",
			packetlen, len(packet[4:]))
	}

	crypt.Decrypt(pcopy[4:])

	if !bytes.Equal(pcopy[4:], packet[4:]) {
		t.Errorf("Encrypted and decrypted packet = %v, expected %v",
			Packet(pcopy), packet)
	}
}

func TestShuffle(t *testing.T) {
	iv := [4]byte{0xFE, 0xCA, 0xDD, 0xBA}
	nextiv := [4]byte{0x81, 0xA5, 0x8F, 0x29}
	crypt := NewCrypt(iv, 62)
	crypt.Shuffle()

	if !bytes.Equal(crypt.IV()[:4], nextiv[:]) {
		t.Errorf("nextiv = % X, expected % X", crypt.IV()[:4], nextiv[:])
	}
}
