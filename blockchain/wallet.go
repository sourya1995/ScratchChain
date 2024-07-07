package blockchain

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func GenerateRSAKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		// panic(err)
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func NewWallet() (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()
	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
}

func(wallet *Wallet) SignTransaction(transaction *Transaction) (string, error) {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender, transaction.Recipient, transaction.Amount, transaction.Coinbase)
	hash := sha256.Sum256([]byte(dataString))
	signature, err := rsa.SignPKCS1v15(rand.Reader, wallet.PrivateKey, crypto.SHA256, hash[:])
	if(err != nil) {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(signature), nil
}

func VerifyTransaction(transaction *Transaction, publicKey *rsa.PublicKey, signature string) error {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender, transaction.Receiver, transaction.Amount, transaction.Coinbase)
	hash := sha256.Sum256([]byte(dataString))
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signatureBytes)
	if err != nil {
		return errors.New("invalid signature")
	}

	return nil
}





