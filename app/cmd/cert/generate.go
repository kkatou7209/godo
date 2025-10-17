package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

func main() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2024)

	if err != nil {
		log.Fatalf("fail to generate key: %v", err)
		return
	}

	info, err := os.Stat("cert")

	if os.IsNotExist(err) || !info.IsDir() {

		err = os.Mkdir("cert", 0755)

		if err != nil {
			log.Fatalf("fail create ket file directory: %v", err)
			return
		}
	}

	privateKeyFile, err := os.Create("cert/private.pem")

	if err != nil {
		log.Fatalf("fail create key file: %v", err)
		return
	}

	defer privateKeyFile.Close()

	publicKeyFile, err := os.Create("cert/public.pem")

	if err != nil {
		log.Fatalf("fail create key file: %v", err)
		return
	}

	defer publicKeyFile.Close()

	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	privateBlock := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: privateBytes,
	}

	publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	
	if err != nil {
		log.Fatalf("fail to create bytes from public key: %v", err)
		return
	}

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	}

	if err = pem.Encode(privateKeyFile, privateBlock); err != nil {
		log.Fatalf("fail to write private key on file: %v", err)
		return
	}

	if err = pem.Encode(publicKeyFile, publicBlock); err != nil {
		log.Fatalf("fail to write public key on file: %v", err)
		return
	}
}
