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
        iv := [4]byte { 0xFE, 0xCA, 0xDD, 0xBA }
        
        pcopy := make([]byte, len(packet))
        copy(pcopy, []byte(packet))
        crypt := NewCrypt(iv, 62)
        crypt.Encrypt(pcopy)
        packetlen := GetPacketLength(pcopy)
        
        if packetlen != len(packet[4:]) {
                t.Errorf("Encrypted packet length = %d, expected %d", packetlen, len(packet[4:]))                
        }
        
        crypt.Decrypt(pcopy[4:])
        
        if !bytes.Equal(pcopy[4:], packet[4:]) {
                t.Errorf("Encrypted and decrypted packet = %v, expected %v", Packet(pcopy), packet)  
        }
}