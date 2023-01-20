package main

import shots_crypto "github.com/rahul0tripathi/go-shots/crypto"

const SECURE_KEY = "PASSWORD"
const PLAIN_TEXT = "ENCRYPT IJ"

func main() {
	shots_crypto.EncryptAndDecrypt(SECURE_KEY, PLAIN_TEXT)
}
