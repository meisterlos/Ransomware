package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

const (
	encryptedFile1 = "Test1.txt"     // Şifrelenmiş dosya adı
	encryptedFile2 = "MRT haber.docx"
	decryptedFile1 = "decrypted.txt" // Çözülmüş dosya adı
	decryptedFile2 = "decrypted.docx"
)

func main() {
	key := []byte("1234567890123456")

	ciphertext, err := ioutil.ReadFile(encryptedFile1)
	if err != nil {
		fmt.Println(err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(ciphertext) < aes.BlockSize {
		fmt.Println("Şifreli metin çok kısa!")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	decryptedText := make([]byte, len(ciphertext)) // Çözülen metni tutacak byte dizisi
	mode.CryptBlocks(decryptedText, ciphertext)

	// Padding'i kaldıralım
	padding := int(decryptedText[len(decryptedText)-1])
	decryptedText = decryptedText[:len(decryptedText)-padding]

	// Çözülen metni yeni bir dosyaya yazalım (txt)
	if err := ioutil.WriteFile(decryptedFile1, decryptedText, 0644); err != nil {
		fmt.Println(err)
		return
	}

	ciphertext, err = ioutil.ReadFile(encryptedFile2)
	if err != nil {
		fmt.Println(err)
		return
	}

	iv = ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode = cipher.NewCBCDecrypter(block, iv)
	decryptedText = make([]byte, len(ciphertext)) // Çözülen metni tutacak byte dizisi
	mode.CryptBlocks(decryptedText, ciphertext)

	// Padding'i kaldıralım
	padding = int(decryptedText[len(decryptedText)-1])
	decryptedText = decryptedText[:len(decryptedText)-padding]

	// Çözülen metni yeni bir dosyaya yazalım (docx)
	if err := ioutil.WriteFile(decryptedFile2, decryptedText, 0644); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Dosyalar başarıyla şifresi çözüldü ve", decryptedFile1, "ve", decryptedFile2, "adlı dosyalara kaydedildi!")
}