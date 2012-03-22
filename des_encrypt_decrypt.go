package main

import (
	// "bufio"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"io"
)

func main() {
	c, err := des.NewCipher([]byte("32129232"))
	if err != nil {
		panic(err)
	}

	message := bytes.NewBufferString("this isn't a really secret, do you know that? no secret is best secret.")

	fmt.Println("original bytes:")
	fmt.Println([]byte(message.String()))

	encrypted := bytes.NewBuffer(nil)

	// var encrypted bytes.Buffer

	b8 := make([]byte, 8)

	w := cipher.StreamWriter{S: cipher.NewOFB(c, b8), W: encrypted}

	io.Copy(w, message)

	fmt.Println("encrypted bytes:")
	fmt.Println(encrypted.Bytes())

	encryptedOutput := fmt.Sprintf("%02x", encrypted.Bytes())

	fmt.Println(encryptedOutput)

	digital := fmt.Sprintf("%x", []byte("that's cool!"))

	var m []byte
	fmt.Sscanf(digital, "%x", &m)

	fmt.Println(string(m))

	c2, err := des.NewCipher([]byte("32129232"))

	w2 := cipher.StreamReader{S: cipher.NewOFB(c2, b8), R: encrypted}

	decrypted_message := bytes.NewBuffer(nil)

	io.Copy(decrypted_message, w2)

	fmt.Println(decrypted_message.String())

}
