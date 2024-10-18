package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// Convert bytes into bits
func bits(b []byte) []int {
	var result []int
	for i := 0; i < len(b); i++ {
		for j := 7; j >= 0; j-- {
			result = append(result, int(b[i]>>uint(j)&1))
		}
	}
	return result
}

func hexToBits(hexString string) ([]int, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	var bits []int
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			bits = append(bits, int((b>>i)&1))
		}
	}

	return bits, nil
}

func roundUpToNearestMultipleOf512(n int) int {
	return ((n + 512 - 1) / 512) * 512
}

func padding(bits []int, totalLength int) []int {

	bits = append(bits, 1)

	noOfZerosToPad := totalLength - len(bits) - 64
	for i := 0; i < noOfZerosToPad; i++ {
		bits = append(bits, 0)
	}

	return bits
}

func to64Bit(size int) []int {
	// Create a byte slice of 8 bytes
	sizeBytes := make([]byte, 8)
	// Convert the size to uint64 and store it in big-endian format
	binary.BigEndian.PutUint64(sizeBytes, uint64(size))

	var bits []int
	for _, b := range sizeBytes {
		for j := 7; j >= 0; j-- {
			bits = append(bits, int(b>>uint(j)&1))
		}
	}
	return bits
}

func main() {
	x := []byte("Hello saldifou oiasfoiasdhf")
	bitsX := bits(x)
	numberOfBits := len(bitsX)
	nextMultipleOf512 := roundUpToNearestMultipleOf512(numberOfBits + 65)

	paddedBits := padding(bitsX, nextMultipleOf512)
	last64Bits := to64Bit(numberOfBits)
	mergedArray := append(paddedBits, last64Bits...)
	fmt.Println("Merged Array:", mergedArray)
	fmt.Println("Merged Array:", len(mergedArray))

	hexString := "6a09e667"

	bits, err := hexToBits(hexString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var HConstants = []int{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}

	bits2, err := hexToBits(fmt.Sprintf("%x", HConstants[0]))

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var KConstants = []int{
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

	fmt.Println("K Constants:", KConstants)

	fmt.Println("H Constants:", HConstants)

	fmt.Println("Bits:", bits)
	fmt.Println("Bits2:", bits2)

}
