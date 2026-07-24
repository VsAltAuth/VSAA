package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"golang.org/x/crypto/bcrypt"
)

func privkeyFile() string {
	file := os.Getenv("PRIVKEY_FILE")
	if file == "" {
		return "private.key"
	}
	return file
}

func PubkeyFile() string {
	file := os.Getenv("PUBKEY_FILE")
	if file == "" {
		return "public_key.pem"
	}
	return file
}

// I'm not sure if throwing panic so much is a good idea... will consult smart people later
func PrivateKey() (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(privkeyFile())
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("privateKey() - decoded PEM data is nil")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privateKey, nil
	}
	return nil, err
}

func Sign(unsignedval string) (string, error) {
	val := []byte(unsignedval)
	hashed := sha256.Sum256(val)
	privateKey, err := PrivateKey()
	if err != nil {
		return "no privatekey", err
	}
	signed, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "invalid", err
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
