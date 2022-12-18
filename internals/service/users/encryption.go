package users

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/BigNutJaa/users/internals/config"
)

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func StartEncrypt(password string) string {
	StringToEncrypt := password

	// To encrypt the StringToEncrypt
	encText, err := Encrypt(StringToEncrypt, config.MySecret)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	//fmt.Println("Encrypt =", encText)
	return encText
}

func StartDecrypt(password string) string {
	// To decrypt the StringToEncrypt
	decText, err := Decrypt(password, config.MySecret)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	//fmt.Println("Decrypt =", decText)
	return decText
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
