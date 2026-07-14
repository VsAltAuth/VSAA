package utils

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	_ "github.com/joho/godotenv/autoload"
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
func PrivateKey() *rsa.PrivateKey {
	keyData, err := os.ReadFile(privkeyFile())
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(keyData)
	if block == nil {
		panic("privateKey() - decoded PEM data is nil")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return privateKey
	}
	panic(err)
}

func Sign(unsignedval string) (string, error) {
	val := []byte(unsignedval)
	signed, err := rsa.SignPKCS1v15(nil, PrivateKey(), crypto.SHA256, val[:])
	if err != nil {
		return "invalid", err
	}
	signedString := string(signed)
	return signedString, nil
}
