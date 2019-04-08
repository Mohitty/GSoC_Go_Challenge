# GSoC_Go_Challenge
This is a qualification task for GSoc 2019

## Problem Statement:
Input: Given a hexadecimal encoded ciphertext, which has been ciphered using single byte repeating key.
Output: The original plaintext

### Solution Approach:
1. Calculated the frequency of characters from the following text: http://www.gutenberg.org/files/11/11-0.txt
2. Performed the XOR with all the ASCII characters and calculater the score of resulting string based on above calculated frequencies.
3. Chose the plain text with the biggest score.

#### Usage:
To run the program:
1. Clone the repository and go to cloned directory.
2. Run `go run decipher.go`.
