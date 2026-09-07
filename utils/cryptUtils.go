package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"golang.org/x/crypto/bcrypt"
)

var publicKey string
var privateKey *rsa.PrivateKey
var pubKeyOnce sync.Once
var privKeyOnce sync.Once

func privkeyFile() string {
	file := os.Getenv("PRIVKEY_FILE")
	if file == "" {
		return "private.key"
	}
	return file
}

func GetPublicKey() string {
	pubKeyOnce.Do(func() {
		if privateKey == nil {
			GetPrivateKey()
		}
		pubKey := &privateKey.PublicKey
		pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
		if err != nil {
			fmt.Printf("error marshaling public key: %v", err)
			return
		}
		pemBlock := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubKeyBytes,
		}
		publicKey = string(pem.EncodeToMemory(pemBlock))
	})
	return publicKey
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	var critError error
	privKeyOnce.Do(func() {
		keyData, err := os.ReadFile(privkeyFile())
		if err != nil {
			critError = err
			return
		}
		block, _ := pem.Decode(keyData)
		if block == nil {
			critError = fmt.Errorf("GetPrivateKey() - decoded PEM data is nil")
			return
		}
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err == nil {
			critError = err
			return
		}
		critError = err
	})
	if critError != nil {
		return nil, critError
	}
	return privateKey, nil

}

func Sign(unsignedval string) (string, error) {
	if privateKey == nil {
		GetPrivateKey()
	}
	val := []byte(unsignedval)
	hashed := sha256.Sum256(val)
	signed, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	signedString := string(signed)
	return signedString, nil
}

func BHashPass(pass string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "ERROR", err
	}
	return string(hashed), err
}

func BCompare(hashed string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	return err
}

func generateSalt() (string, error) {
	a := int(bits.Reverse64(uint64(time.Now().Unix())))
	b := []byte(strconv.Itoa(a))
	if b == nil {
		return "", fmt.Errorf("generateSalt(): b is nil but cannot be nil!")
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func hashPassword(pass string, salt string) string {
	saltedPass := pass + salt
	hash := sha256.Sum256([]byte(saltedPass))
	return hex.EncodeToString(hash[:])
}
