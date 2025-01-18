package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RSA加密函数
//
// @param publicKey string 公钥
//
// @param plainText string 明文
func RSAEncrypt(publicKey string, plainText string) ([]byte, error) {

	// 解码PEM格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	// 使用RSA_PKCS1_PADDING进行加密
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, []byte(plainText))
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func RSADecrypt(privateKey string, cipherText []byte) (string, error) {
	// 解码PEM格式的私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the private key")
	}

	// 解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	privKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not an RSA private key")
	}

	// 使用RSA_PKCS1_PADDING进行解密
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cipherText)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
