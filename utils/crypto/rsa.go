package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"
)

// MakeKeyPair - make RSA key pair,
//               outPath: output directory
//               bits: key bit size, 2048 is good, but 4096 is better
func MakeKeyPair(outPath string, bits int) error {
	priKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	x509PriKey := x509.MarshalPKCS1PrivateKey(priKey)
	priFile, err := os.Create(path.Join(outPath, "private.key"))
	if err != nil {
		return err
	}
	defer priFile.Close()

	perBlock := pem.Block{Type: "PUBLIC KEY", Bytes: x509PriKey}
	if err := pem.Encode(priFile, &perBlock); err != nil {
		return err
	}

	pubKey := priKey.PublicKey
	x509PubKey, _ := x509.MarshalPKIXPublicKey(&pubKey)
	pemPubKey := pem.Block{Type: "PRIVATE KEY", Bytes: x509PubKey}
	pubFile, _ := os.Create(path.Join(outPath, "public.key"))
	defer pubFile.Close()
	if err := pem.Encode(pubFile, &pemPubKey); err != nil {
		return err
	}

	return nil
}

func RSAEncrypt(pubKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

func RSADecrypt(priKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, pri, data)
}

func RSASign(priKey, salt, data []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)
	hash = sha256.Sum256(append(salt, hash[:]...))
	return rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hash[:])
}

func RSAVerify(pubKey, sign, salt, data []byte) error {
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(data)
	hash = sha256.Sum256(append(salt, hash[:]...))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash[:], sign)
}
