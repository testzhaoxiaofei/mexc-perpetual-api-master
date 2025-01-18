package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// DecryptAESGCM 解密AES-GCM加密的数据
// cipherText: 完整的密文（包含nonce和authTag）
// key: 32字节的AES密钥.(20250114为fp.umd.js中的de = "1b8c71b668084dda9dc0285171ccf753"这个值)
func Aes256GCMDecrypt(cipherText, key []byte) (string, error) {
	const (
		NonceSize = 12 // GCM standard nonce size
		TagSize   = 16 // GCM authentication tag size
	)

	// Validate key length for AES-256
	if len(key) != 32 {
		return "", errors.New("key length must be 32 bytes for AES-256")
	}

	// Validate cipherText length
	if len(cipherText) < NonceSize+TagSize {
		return "", errors.New("cipherText is too short")
	}

	// Extract nonce, ciphertext (encrypted data), and authTag
	nonce := cipherText[:NonceSize]
	authTag := cipherText[len(cipherText)-TagSize:]
	encryptedData := cipherText[NonceSize : len(cipherText)-TagSize]

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	// Wrap block in GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Append authTag to the encrypted data (GCM expects them together)
	fullCipherText := append(encryptedData, authTag...)

	// Decrypt data
	plaintext, err := gcm.Open(nil, nonce, fullCipherText, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plaintext), nil
}

// Aes256GCMEncrypt 加密数据为AES-GCM格式
// plainText: 待加密的明文数据
// key: 32字节的AES密钥
func Aes256GCMEncrypt(plainText, key []byte) ([]byte, error) {
	// Ensure the key length is 32 bytes (256 bits)
	if len(key) != 32 {
		return nil, fmt.Errorf("invalid key length: must be 32 bytes")
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %v", err)
	}

	// Create a GCM (Galois/Counter Mode) instance
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %v", err)
	}

	// Generate a random nonce (IV) of the appropriate length
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	// Encrypt the plaintext using the nonce and GCM instance
	ciphertext := gcm.Seal(nil, nonce, plainText, nil)

	// Prepend the nonce to the ciphertext for use during decryption
	result := append(nonce, ciphertext...)

	return result, nil
}
