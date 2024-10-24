package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var secretKey = make([]byte, 32)
var nonce = make([]byte, 12)

func init() {
	// Generate a random key (replace with a secure key in production)
	// keyBytes, _ := base64.StdEncoding.DecodeString("SkI3UnOLNzrAh1/RFcVjmd+4DaGqsW8=")
	keyBytes := []byte("SkI3UnOLNzrAh1/RFcVjmd+4DaGqsW8=")
	copy(secretKey, keyBytes)
	// copy(secretKey, tempSecretKey)
	// if _, err := rand.Read(secretKey); err != nil {
	// 	panic(err)
	// }

	// Generate a random nonce (must be 16 bytes for AES-256-GCM)
	// nonceBytes, _ := base64.StdEncoding.DecodeString("+SFP1asdW9VS5lg=")
	nonceBytes := []byte("+SFP1asdW9VS5lg=")
	copy(nonce, nonceBytes)
	// if _, err := rand.Read(nonce); err != nil {
	// 	panic(err)
	// }
}

func main() {
	fmt.Println("Secret Key: ", secretKey, len(secretKey), string(secretKey))
	fmt.Println("Nonce: ", nonce, len(nonce), string(nonce))
	http.HandleFunc("/", index)
	http.HandleFunc("/authenticate", authenticate)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("input")
		encryptedValue, err := encrypt(e, secretKey)
		if err != nil {
			http.Error(w, "Encryption error ", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		c.Value = e + `|` + getCodeHmac(e) + `|` + getCodeBase64(e) + `|` + string(encryptedValue)
	}

	http.SetCookie(w, c)

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
		  <input type="text" name="input">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This `+c.Value+`</a>
	  </body>
	</html>`)
}

func getCodeHmac(data string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func decodeBase64(data string) string {
	bs, _ := base64.URLEncoding.DecodeString(data)
	return string(bs)
}

func getCodeBase64(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func encrypt(plaintext string, key []byte) (string, error) {
	// Create a new AES-GCM cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext
	decodedText := []byte(plaintext)
	ciphertext := cipher.Seal(nil, nonce, decodedText, nil)

	// Combine ciphertext and nonce for storage or transmission
	combined := append(nonce, ciphertext...)

	// Encode the combined data as a base64 string
	encoded := base64.StdEncoding.EncodeToString(combined)

	return encoded, nil
}

func decrypt(encoded string, key []byte) (string, error) {
	// Decode the base64 string
	combined, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Separate the nonce and ciphertext
	nonceSize := cipher.NonceSize()
	nonce := combined[:nonceSize]
	ciphertext := combined[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func authenticate(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeHmacRcvd := xs[1]
	codeBase64Rcvd := xs[2]
	codeAesRcvd := xs[3]
	codeCheckHmac := getCodeHmac(email)
	codeCheckBaseCode64 := getCodeBase64(email)
	codeCheckAes, err := decrypt(codeAesRcvd, secretKey)

	if err != nil {
		fmt.Println("AES codes didn't match")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if codeHmacRcvd != codeCheckHmac {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeHmacRcvd)
		fmt.Println(codeCheckHmac)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if codeBase64Rcvd != codeCheckBaseCode64 {
		fmt.Println("BASE64 codes didn't match")
		fmt.Println(codeBase64Rcvd)
		fmt.Println(codeCheckBaseCode64)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>HMAC</h1>
	  	<h1>`+codeHmacRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheckHmac+` - RECALCULATED </h1>
		<br></br>
		<h1>BASE64</h1>
		<h1>`+codeBase64Rcvd+` - RECEIVED </h1>
	  	<h1>`+decodeBase64(codeBase64Rcvd)+` - RECALCULATED </h1>
		<br></br>
		<h1>AES</h1>
		<h1>`+codeAesRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheckAes+` - RECALCULATED </h1>
	  </body>
	</html>`)
}
