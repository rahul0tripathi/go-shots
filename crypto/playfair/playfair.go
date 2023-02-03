package playfair

import (
	"fmt"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func constructPlayfairMatrix(key string) ([5][5]byte, map[byte]coordinate) {
	key = strings.ToUpper(key)
	matrix := [5][5]byte{}
	i := 0
	j := 0
	referenceMatrix := map[byte]bool{}
	indexMap := map[byte]coordinate{}
	for _, v := range key {
		if byte(v) == 74 {
			continue
		}
		if _, ok := referenceMatrix[byte(v)]; !ok {
			matrix[i][j] = byte(v)
			indexMap[byte(v)] = coordinate{
				x: i,
				y: j,
			}
			if j == 4 {
				j = -1
				i++
			}
			j++
			referenceMatrix[byte(v)] = true

		}
	}
	for k := uint8(65); k < uint8(91); k++ {
		if k == 74 {
			continue
		}
		if _, ok := referenceMatrix[k]; !ok {
			matrix[i][j] = k
			indexMap[k] = coordinate{
				x: i,
				y: j,
			}
			if j == 4 {
				j = -1
				i++
			}
			if i == 5 {
				break
			}
			j++
		}

	}
	return matrix, indexMap
}

func encrypt(plainText string, matrix [5][5]byte, playfairMatrixMap map[byte]coordinate) (string, error) {
	plainText = padRepeatingCharacters(strings.ReplaceAll(strings.ToUpper(plainText), " ", ""))
	if len(plainText)%2 != 0 {
		plainText += "Z"
	}
	fmt.Printf("MODIFIED PLAIN TEXT := %s\n", plainText)
	fmt.Println("-------------------------------")
	encrypted := ""
	iterator := 0
ENCRYPT:
	fmt.Printf("| encrypting  | %c | %c\n", plainText[iterator], plainText[iterator+1])
	w1 := playfairMatrixMap[plainText[iterator]]
	w2 := playfairMatrixMap[plainText[iterator+1]]
	encryptedW1 := coordinate{}
	encryptedW2 := coordinate{}
	if w1.x == w2.x {
		// same row case
		encryptedW1.x = w1.x
		encryptedW1.y = w1.y + 1
		encryptedW2.x = w2.x
		encryptedW2.y = w2.y + 1
	} else if w1.y == w2.y {
		// same column case
		encryptedW1.x = w1.x + 1
		encryptedW1.y = w1.y
		encryptedW2.x = w2.x + 1
		encryptedW2.y = w2.y
	} else {
		encryptedW1.x = w1.x
		encryptedW1.y = w2.y
		encryptedW2.x = w2.x
		encryptedW2.y = w1.y
	}
	parsedW1 := string(matrix[wrapIndex(encryptedW1.x)][wrapIndex(encryptedW1.y)])
	parsedW2 := string(matrix[wrapIndex(encryptedW2.x)][wrapIndex(encryptedW2.y)])
	fmt.Printf("| encrypted   | %s | %s\n\n", parsedW1, parsedW2)
	encrypted += parsedW1 + parsedW2
	iterator += 2
	if iterator < len(plainText) {
		goto ENCRYPT
	}
	return encrypted, nil
}

func decrypt(cipherText string, matrix [5][5]byte, playfairMatrixMap map[byte]coordinate) (string, error) {
	decrypted := ""
	iterator := 0
DECRYPT:
	fmt.Printf("| decrypting  | %c | %c\n", cipherText[iterator], cipherText[iterator+1])
	w1 := playfairMatrixMap[cipherText[iterator]]
	w2 := playfairMatrixMap[cipherText[iterator+1]]
	decryptedW1 := coordinate{}
	decryptedW2 := coordinate{}
	if w1.x == w2.x {
		// same row case
		decryptedW1.x = w1.x
		decryptedW1.y = w1.y - 1
		decryptedW2.x = w2.x
		decryptedW2.y = w2.y - 1
	} else if w1.y == w2.y {
		// same column case
		decryptedW1.x = w1.x - 1
		decryptedW1.y = w1.y
		decryptedW2.x = w2.x - 1
		decryptedW2.y = w2.y
	} else {
		decryptedW1.x = w1.x
		decryptedW1.y = w2.y
		decryptedW2.x = w2.x
		decryptedW2.y = w1.y
	}
	parsedW1 := string(matrix[wrapIndex(decryptedW1.x)][wrapIndex(decryptedW1.y)])
	parsedW2 := string(matrix[wrapIndex(decryptedW2.x)][wrapIndex(decryptedW2.y)])
	fmt.Printf("| decrypted   | %s | %s\n\n", parsedW1, parsedW2)
	decrypted += parsedW1 + parsedW2
	iterator += 2
	if iterator < len(cipherText) {
		goto DECRYPT
	}
	return decrypted, nil
}

func EncryptAndDecrypt(key string, plainText string) {
	matrix, indexMap := constructPlayfairMatrix(key)
	logMatrix(matrix)
	fmt.Printf("ENCRYPTING := %s \n", plainText)
	fmt.Println("-------------------------------")
	encrypted, err := encrypt(plainText, matrix, indexMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("-------------------------------")
	fmt.Printf("ENCRYPTED := %s\n\n", encrypted)
	decrypted, err := decrypt(encrypted, matrix, indexMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("-------------------------------")
	fmt.Printf("DECRYPTED := %s", decrypted)
}

// util functions that no one cares about
func logMatrix(matrix [5][5]byte) {
	fmt.Println("---- MATRIX ----")
	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			fmt.Printf("%c ", matrix[m][n])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func padRepeatingCharacters(input string) string {
	paddedString := ""
	for _, v := range input {
		if len(paddedString) > 0 && paddedString[len(paddedString)-1] == uint8(v) {
			paddedString += "X"
		}
		// replace J by I
		if byte(v) == 74 {
			v = 73
		}
		paddedString += fmt.Sprintf("%c", v)
	}
	return paddedString
}

func wrapIndex(i int) int {
	if i < 0 {
		return 4
	}
	return i % 5
}
