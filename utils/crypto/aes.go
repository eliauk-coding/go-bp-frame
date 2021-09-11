package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESEncrypt - The key argument should be the AES key,
//              either 16, 24, or 32 bytes to select
//              AES-128, AES-192, or AES-256.
func AESEncrypt(payload, key []byte) ([]byte, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(payload))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	enc := cipher.NewCFBEncrypter(block, iv)
	enc.XORKeyStream(encrypted, payload)
	return encrypted, nil
}

// AESDecrypt - The key argument should be the AES key,
//              either 16, 24, or 32 bytes to select
//              AES-128, AES-192, or AES-256.
func AESDecrypt(encrypted, key []byte) ([]byte, error) {
	var err error
	var block cipher.Block
	var iv = key[:aes.BlockSize]
	decrypted := make([]byte, len(encrypted))
	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(decrypted, encrypted)
	return decrypted, nil
}
