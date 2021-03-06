package utils

import (
	// Native packages
	"crypto/rand"
	"log"

	// 3rd party packages
	"golang.org/x/crypto/bcrypt"
)

/**
 *	Returns expectedString if not empty, or fallbacks to defaultString value.
 *	@NOTE Helper function for ternery string checks in Go.
 *
 *	@param expectedString string - String to expect not to be empty.
 *	@param defaultString string - String to return if expectedString is empty.
 *
 *	@return string
 */
func Pick(expectedString string, defaultString string) string {
	if expectedString == "" {
		return defaultString
	}

	return expectedString
}

/**
 *	Generates a random string at desired length.
 *
 *	@param outputLength int - Output string length
 *
 *	@return string
 */
func RandomString(outputLength int) string {
	var bitLength byte
	var bitMask byte

	availableCharBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	availableCharLength := len(availableCharBytes)

	for bits := availableCharLength - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}

	bitMask = 1<<bitLength - 1
	bufferSize := outputLength + outputLength/3

	outputString := make([]byte, outputLength)
	for i, j, randomBytes := 0, 0, []byte{}; i < outputLength; j++ {
		if j%bufferSize == 0 {
			randomBytes = SecureRandomBytes(bufferSize)
		}
		if n := int(randomBytes[j%outputLength] & bitMask); n < availableCharLength {
			outputString[i] = availableCharBytes[n]
			i++
		}
	}

	return string(outputString)
}

/**
 *	Returns a slice with random bytes.
 *
 *	@param length int - Desired byte length.
 *
 *	@return []byte
 */
func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)

	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}

	return randomBytes
}

/**
 *	@const PasswordEncryptCost int
 */
const PasswordEncryptCost = 13

/**
 *	Creates and returns a password hash using bcrypt.
 *
 *	@param rawPassword string - Unhashed password to encrypt.
 *
 *	@return string
 */
func PasswordCreate(rawPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), PasswordEncryptCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

/**
 *	Verifies a hashed password with an unhashed password to see if they match.
 *
 *	@param hashedPassword string - Hashed password from storage.
 *	@param rawPassword string - Unhashed password from input.
 *
 *	@return bool
 */
func PasswordMatch(hashedPassword string, rawPassword string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)) != nil {
		return false
	}
	return true
}
