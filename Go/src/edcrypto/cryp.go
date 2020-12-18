package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
	"io"
)

func Encrypt(b cipher.Block, plaintext []byte) []byte {
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	mode := cipher.NewCBCEncrypter(b, iv)

	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext

}

func Decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화 된 데이터의 길이는 블록 크기의 배수가")
		return nil
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(b, iv)

	mode.CryptBlocks(plaintext, ciphertext)

	return plaintext
}

func main() {
	// hash
	s := "Hello, world! 12"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("hello "))
	sha.Write([]byte("world"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
	// aes 16바이트 암호화
	key := "Hello, Key 12345"
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	cipherText := make([]byte, len(s))
	block.Encrypt(cipherText, []byte(s))
	fmt.Printf("%x\n", cipherText)

	plainText := make([]byte, len(s))
	block.Decrypt(plainText, cipherText)

	fmt.Println(string(plainText))
	// cbc

	s3 := "동해 물과 백두산이 마르고 닳도록 하느님이 보우하사 우리 나라 만세. 무궁화 삼천리 화려강산 대한 사람, 대한으로 길이 보전하세."

	block3, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	cipherText2 := Encrypt(block3, []byte(s3))
	fmt.Printf("%x\n", cipherText2)

	plaintext2 := Decrypt(block3, cipherText2)
	fmt.Println(string(plaintext2))
	/// rsa 공개키
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	ciphertext3, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s3),
	)

	fmt.Printf("%x\n", ciphertext3)

	plaintext3, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext3,
	)

	fmt.Println(string(plaintext3))

}
