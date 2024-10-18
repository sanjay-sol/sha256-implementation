package main

import (
	"encoding/binary"
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


}
