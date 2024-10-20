package main

import (
	"encoding/binary"
	"fmt"
)

//*   ==========================
//*       Utility Functions
//*   ==========================

// ? Convert bytes into bits
func bits(b []byte) []int {
	var result []int
	for i := 0; i < len(b); i++ {
		for j := 7; j >= 0; j-- {
			result = append(result, int(b[i]>>uint(j)&1))
		}
	}
	return result
}

func roundUpToNearestMultipleOf512(n int) int {
	return ((n + 512 - 1) / 512) * 512
}

// ? Convert an integer size to a 64-bit representation
func to64Bit(size int) []int {
	sizeBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(sizeBytes, uint64(size))

	var bits []int
	for _, b := range sizeBytes {
		for j := 7; j >= 0; j-- {
			bits = append(bits, int(b>>uint(j)&1))
		}
	}
	return bits
}

// ? Convert bits to bytes
func bitsToBytes(bits []int) []byte {
	bytes := make([]byte, len(bits)/8)
	for i := 0; i < len(bits); i += 8 {
		for j := 0; j < 8; j++ {
			bytes[i/8] = bytes[i/8]<<1 + byte(bits[i+j])
		}
	}
	return bytes
}

//*		===============================
//* 	  Padding and Block Handling
//* 	==============================

// ? Pad the bit array according to the SHA-256 specification
func padding(bits []int, totalLength int) []int {

	bits = append(bits, 1)

	noOfZerosToPad := totalLength - len(bits) - 64
	for i := 0; i < noOfZerosToPad; i++ {
		bits = append(bits, 0)
	}

	return bits
}

// ? Divide the padded message into 512-bit blocks and split into 32-bit words
func divideIntoBlocks(paddedMessage []int) [][]uint32 {
	var blocks [][]uint32
	bytes := bitsToBytes(paddedMessage)
	blockCount := len(bytes) / 64 // 512 bits = 64 bytes per block

	for i := 0; i < blockCount; i++ {
		var block []uint32
		for j := 0; j < 16; j++ {
			word := binary.BigEndian.Uint32(bytes[i*64+j*4 : i*64+j*4+4])
			block = append(block, word)
		}
		blocks = append(blocks, block)
	}

	return blocks
}

//* 		==========================
//* 			 SHA-256 Functions
//* 		==========================

// ? SHA-256 auxiliary functions
func rotr(x uint32, n uint) uint32 {
	return (x >> n) | (x << (32 - n))
}

func sigma0(x uint32) uint32 {
	return rotr(x, 7) ^ rotr(x, 18) ^ (x >> 3)
}

func sigma1(x uint32) uint32 {
	return rotr(x, 17) ^ rotr(x, 19) ^ (x >> 10)
}

func summation0(x uint32) uint32 {
	return rotr(x, 2) ^ rotr(x, 13) ^ rotr(x, 22)
}

func summation1(x uint32) uint32 {
	return rotr(x, 6) ^ rotr(x, 11) ^ rotr(x, 25)
}

func ch(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func maj(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

// ? Generate the message schedule for the current block
func generateMessageSchedule(block []uint32) []uint32 {

	W := make([]uint32, 64)
	copy(W[:16], block)

	for i := 16; i < 64; i++ {
		s0 := sigma0(W[i-15])
		s1 := sigma1(W[i-2])
		W[i] = s1 + W[i-7] + s0 + W[i-16]
	}
	return W
}

//* ==========================
//* Main Function
//* ==========================

func main() {
	x := []byte("")
	bitsX := bits(x)
	numberOfBits := len(bitsX)
	nextMultipleOf512 := roundUpToNearestMultipleOf512(numberOfBits + 65)

	paddedBits := padding(bitsX, nextMultipleOf512)
	last64Bits := to64Bit(numberOfBits)
	mergedArray := append(paddedBits, last64Bits...)

	blocks := divideIntoBlocks(mergedArray)

	for _, block := range blocks {
		generateMessageSchedule(block)
	}

	//! SHA-256 constants
	var HConstants = []uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}

	var KConstants = []uint32{
		0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5,
		0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3,
		0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc,
		0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7,
		0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13,
		0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
		0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5,
		0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208,
		0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2,
	}

	// Process each block
	for _, block := range blocks {
		W := generateMessageSchedule(block)

		a, b, c, d, e, f, g, h := HConstants[0], HConstants[1], HConstants[2], HConstants[3], HConstants[4], HConstants[5], HConstants[6], HConstants[7]

		for i := 0; i < 64; i++ {
			T1 := h + summation1(e) + ch(e, f, g) + KConstants[i] + W[i]
			T2 := summation0(a) + maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
		}

		// Update hash values
		HConstants[0] += a
		HConstants[1] += b
		HConstants[2] += c
		HConstants[3] += d
		HConstants[4] += e
		HConstants[5] += f
		HConstants[6] += g
		HConstants[7] += h
	}

	fmt.Printf("Final hash: %08x%08x%08x%08x%08x%08x%08x%08x\n", HConstants[0], HConstants[1], HConstants[2], HConstants[3], HConstants[4], HConstants[5], HConstants[6], HConstants[7])

}
