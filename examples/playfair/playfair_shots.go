package main

import (
	"github.com/rahul0tripathi/go-shots/crypto/playfair"
)

const SECURE_KEY = "PASSWORD"
const PLAIN_TEXT = "ENCRYPT IJ"

func main() {
	playfair.EncryptAndDecrypt(SECURE_KEY, PLAIN_TEXT)
}
