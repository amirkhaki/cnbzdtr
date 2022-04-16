package main

import (
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func main() {
	help := `usage:
	./decrypt KEY ENCRYPTED
		KEY is encryption key
		ENCRYPTED is string to decrypt
	`
	if len(os.Args) < 3 {
		fmt.Println(help)
		return
	}
	msg, err := decrypt([]byte(os.Args[1]), os.Args[2])
	if err != nil {
		fmt.Println("error decrypting ",err)
		return
	}
	fmt.Println(msg)

}

func decrypt(key []byte, cryptoText string) (msg string, err error) {
	ciphertext, err := base64.URLEncoding.DecodeString(cryptoText)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		err = fmt.Errorf("ciphertext too short")
		return
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	msg = fmt.Sprintf("%s", ciphertext)
	return
}
