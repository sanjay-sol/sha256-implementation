# SHA-256 Implementation in Go

Reference Paper : [The cryptographic hash function SHA-256](https://helix.stormhub.org/papers/SHA-256.pdf)

## Overview

SHA-256 (Secure Hash Algorithm 256) is a cryptographic hash function that produces a 256-bit (32-byte) hash value for any input. 

## How It Works

The implementation follows the SHA-256 algorithm, which consists of the following main steps:

1. **Input Preprocessing**: Convert the input message into a bit representation and apply necessary padding.
2. **Block Processing**: Divide the padded message into 512-bit blocks.
3. **Hash Computation**: For each block, compute the hash using a series of logical operations and constants.
4. **Output the Final Hash**: Combine the computed hash values into a final output.



### Key SHA-256 Functions

- **`rotr(x uint32, n uint) uint32`**: Right rotates a 32-bit integer.
- **`sigma0(x uint32) uint32`**: Auxiliary function for SHA-256.
- **`sigma1(x uint32) uint32`**: Auxiliary function for SHA-256.
- **`summation0(x uint32) uint32`**: Auxiliary function for SHA-256.
- **`summation1(x uint32) uint32`**: Auxiliary function for SHA-256.
- **`ch(x, y, z uint32) uint32`**: Choice function for SHA-256.
- **`maj(x, y, z uint32) uint32`**: Majority function for SHA-256.
- **`generateMessageSchedule(block []uint32) []uint32`**: Generates the message schedule for the current block.





