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
	"bytes"
	"testing"
)

func TestEncode1(t *testing.T) {
	var val, out, bytecount = byte(0xBA), []byte{0xBA}, 1
	p := NewPacket()
	p.Encode1(val)

	if len(p) < bytecount {
		t.Errorf("len(p) after p.Encode1(0x%X) = %d, expected %d", val, len(p), bytecount)
	}

	if !bytes.Equal(p, out) {
		t.Errorf("p after p.Encode1(0x%X) = %v, expected % X", val, p, out)
	}
}

func TestEncode2(t *testing.T) {
	var val, out, bytecount = uint16(0xBAAD), []byte{0xAD, 0xBA}, 2
	p := NewPacket()
	p.Encode2(val)

	if len(p) < bytecount {
		t.Errorf("len(p) after p.Encode2(0x%X) = %d, expected %d", val, len(p), bytecount)
	}

	if !bytes.Equal(p, out) {
		t.Errorf("p after p.Encode2(0x%X) = %v, expected % X", val, p, out)
	}
}

func TestEncode4(t *testing.T) {
	var val, out, bytecount = uint32(0xBAADF00D), []byte{0x0D, 0xF0, 0xAD, 0xBA}, 4
	p := NewPacket()
	p.Encode4(val)

	if len(p) < bytecount {
		t.Errorf("len(p) after p.Encode4(0x%X) = %d, expected %d", val, len(p), bytecount)
	}

	if !bytes.Equal(p, out) {
		t.Errorf("p after p.Encode4(0x%X) = %v, expected % X", val, p, out)
	}
}

func TestEncode8(t *testing.T) {
	var val, out, bytecount = uint64(0xBAADF00DBADDCAFE), []byte{0xFE, 0xCA, 0xDD, 0xBA, 0x0D, 0xF0, 0xAD, 0xBA}, 8
	p := NewPacket()
	p.Encode8(val)

	if len(p) < bytecount {
		t.Errorf("len(p) after p.Encode8(0x%X) = %d, expected %d", val, len(p), bytecount)
	}

	if !bytes.Equal(p, out) {
		t.Errorf("p after p.Encode8(0x%X) = %v, expected % X", val, p, out)
	}
}

func TestEncodeBuffer(t *testing.T) {
        var val, out, bytecount = 
                []byte{0xAA, 0xBB, 0xCC, 0xDD}, 
                []byte{0x04, 0x00, 0xAA, 0xBB, 0xCC, 0xDD}, 
                4 + 2
                
        p := NewPacket()
        p.EncodeBuffer(val)
        
        if len(p) < bytecount {
            t.Errorf("len(p) after p.EncodeBuffer(% X) = %d, expected %d", val, len(p), bytecount)
        }
        
        if !bytes.Equal(p, out) {
            t.Errorf("p after p.EncodeBuffer(%v) = %v, expected % X", val, p, out)
        }   
}

func TestEncodeString(t *testing.T) {
        var val, out, bytecount = 
                "loli", 
                []byte{0x04, 0x00, 'l', 'o', 'l', 'i'}, 
                4 + 2
                
        p := NewPacket()
        p.EncodeString(val)
        
        if len(p) < bytecount {
            t.Errorf("len(p) after p.EncodeString(%s) = %d, expected %d", val, len(p), bytecount)
        }
        
        if !bytes.Equal(p, out) {
            t.Errorf("p after p.EncodeString(%s) = %v, expected % X", val, p, out)
        }   
}

func TestMultipleEncode(t *testing.T) {
	var val1, val2, val4, val8, valbuf, valstr, out, bytecount = 
	        byte(0xAA),
		uint16(0xCCBB),
		uint32(0x00FFEEDD),
		uint64(0x8877665544332211),
		[]byte{0xAA, 0xBB, 0xCC, 0xDD}, 
		"loli", 
		[]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00, 0x11, 0x22, 0x33, 
		       0x44, 0x55, 0x66, 0x77, 0x88, 0x04, 0x00, 0xAA, 0xBB, 0xCC, 
		       0xDD, 0x04, 0x00, 'l', 'o', 'l', 'i'},
		27

	p := NewPacket()
	p.Encode1(val1)
	p.Encode2(val2)
	p.Encode4(val4)
	p.Encode8(val8)
	p.EncodeBuffer(valbuf)
	p.EncodeString(valstr)

	if len(p) < bytecount {
		t.Errorf("len(p) after encoding 0x%X, 0x%X, 0x%X, 0x%X, % X, %s = %d, expected %d",
			val1, val2, val4, val8, valbuf, valstr, len(p), bytecount)
	}

	if !bytes.Equal(p, out) {
		t.Errorf("p after encoding 0x%X, 0x%X, 0x%X, 0x%X, % X, %s = %v, expected % X",
			val1, val2, val4, val8, valbuf, valstr, p, out)
	}
}

func TestDecode1(t *testing.T) {
	const fun = "packet.Decode1(&it)"
	var packet, out = Packet{0xAA}, byte(0xAA)

	it := packet.Begin()
	res, err := packet.Decode1(&it)

	if err != nil {
		t.Errorf("%s: %v", fun, err)
	}

	if res != out {
		t.Errorf("%s = %X, expected %X", fun, res, out)
	}

	if len(it) != 0 {
		t.Errorf("len(it) is non-zero after %s, expected zero", fun)
	}
}

func TestDecode2(t *testing.T) {
	const fun = "packet.Decode2(&it)"
	var packet, out = Packet{0xBB, 0xAA}, uint16(0xAABB)

	it := packet.Begin()
	res, err := packet.Decode2(&it)

	if err != nil {
		t.Errorf("%s: %v", fun, err)
	}

	if res != out {
		t.Errorf("%s = %X, expected %X", fun, res, out)
	}

	if len(it) != 0 {
		t.Errorf("len(it) is non-zero after %s, expected zero", fun)
	}
}

func TestDecode4(t *testing.T) {
	const fun = "packet.Decode4(&it)"
	var packet, out = Packet{0xDD, 0xCC, 0xBB, 0xAA}, uint32(0xAABBCCDD)

	it := packet.Begin()
	res, err := packet.Decode4(&it)

	if err != nil {
		t.Errorf("%s: %v", fun, err)
	}

	if res != out {
		t.Errorf("%s = %X, expected %X", fun, res, out)
	}

	if len(it) != 0 {
		t.Errorf("len(it) is non-zero after %s, expected zero", fun)
	}
}

func TestDecode8(t *testing.T) {
	const fun = "packet.Decode8(&it)"
	var packet, out = Packet{0x11, 0x00, 0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA},
		uint64(0xAABBCCDDEEFF0011)

	it := packet.Begin()
	res, err := packet.Decode8(&it)

	if err != nil {
		t.Errorf("%s: %v", fun, err)
	}

	if res != out {
		t.Errorf("%s = %X, expected %X", fun, res, out)
	}

	if len(it) != 0 {
		t.Errorf("len(it) is non-zero after %s, expected zero", fun)
	}
}

func TestDecodeBuffer(t *testing.T) {
        const fun = "packet.DecodeBuffer(&it)"    
        var packet, out = 
                Packet{0x08, 0x00, 0x11, 0x00, 0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA},
                []byte{0x11, 0x00, 0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA}
        
        it := packet.Begin()
        res, err := packet.DecodeBuffer(&it)
        
        if err != nil {
            t.Errorf("%s: %v", fun, err)
        }
        
        if !bytes.Equal(res, out) {
            t.Errorf("%s = % X, expected % X", fun, res, out)
        }
        
        if len(it) != 0 {
            t.Errorf("len(it) is non-zero after %s, expected zero", fun)
        }
}

func TestDecodeString(t *testing.T) {
        const fun = "packet.DecodeString(&it)"    
        var packet, out = 
                Packet{0x07, 0x00, 'l', 'o', 'l', 'i', 'c', 'o', 'n'},
                "lolicon"
        
        it := packet.Begin()
        res, err := packet.DecodeString(&it)
        
        if err != nil {
            t.Errorf("%s: %v", fun, err)
        }
        
        if res != out {
            t.Errorf("%s = %s, expected %s", fun, res, out)
        }
        
        if len(it) != 0 {
            t.Errorf("len(it) is non-zero after %s, expected zero", fun)
        }
}

func TestMultipleDecode(t *testing.T) {
	var packet, out1, out2, out4, out8, outbuf, outstr = 
	        Packet{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11, 0x00, 0xFF, 
	               0xEE, 0xDD, 0xCC, 0xBB, 0xAA, 0x04, 0x00, 0xAA, 0xBB, 0xCC, 
		       0xDD, 0x04, 0x00, 'l', 'o', 'l', 'i'},
		byte(0x88),
		uint16(0x6677),
		uint32(0x22334455),
		uint64(0xAABBCCDDEEFF0011), 
		[]byte{0xAA, 0xBB, 0xCC, 0xDD}, 
		"loli"

	it := packet.Begin()
	res1, err := packet.Decode1(&it)
	res2, err := packet.Decode2(&it)
	res4, err := packet.Decode4(&it)
	res8, err := packet.Decode8(&it)
	resbuf, err := packet.DecodeBuffer(&it)
	resstr, err := packet.DecodeString(&it)

	if err != nil {
		t.Errorf("multiple decodes: %v", err)
	}

	if res1 != out1 {
		t.Errorf("packet.Decode1(&it) = %X, expected %X", res1, out1)
	}

	if res2 != out2 {
		t.Errorf("packet.Decode2(&it) = %X, expected %X", res2, out2)
	}

	if res4 != out4 {
		t.Errorf("packet.Decode4(&it) = %X, expected %X", res4, out4)
	}

	if res8 != out8 {
		t.Errorf("packet.Decode8(&it) = %X, expected %X", res8, out8)
	}
	
	if !bytes.Equal(resbuf, outbuf) {
	        t.Errorf("packet.DecodeBuffer(&it) = % X, expected % X", resbuf, outbuf)
	}
	
        if resstr != outstr {
                t.Errorf("packet.DecodeString(&it) = %s, expected %s", resstr, outstr)
        }

	if len(it) != 0 {
		t.Errorf("len(it) is non-zero after multiple decodes, expected zero")
	}
}

func TestDecodeFail(t *testing.T) {
	const fun = "packet.Decode8(&it)"
	var packet, out = Packet{0xDD, 0xCC, 0xBB, 0xAA},
		uint64(0xAABBCCDDEEFF0011)

	it := packet.Begin()
	res, err := packet.Decode8(&it)

	if err == nil {
		t.Errorf("%s was supposed to fail but it didn't", fun, err)
	}

	if res == out {
		t.Errorf("%s = %X, but the function was supposed to fail", fun, res, out)
	}

	if len(it) != 4 {
		t.Errorf("len(it) = %d, expected 4", len(it))
	}
}
