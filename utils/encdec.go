package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func Encrypt(plainText []byte, key []byte, iv []byte) (string, error) {
	var plainTextBlock []byte
	l := len(plainText)

	if l%16 != 0 {
		extendBlock := 16 - (l % 16)
		plainTextBlock = make([]byte, l+extendBlock)
		copy(plainTextBlock[l:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, l)
	}

	copy(plainTextBlock, plainText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	chiperText := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(chiperText, plainTextBlock)

	if len(chiperText) == 0 {
		return "", errors.New("unable to encrypt data")
	}

	str := base64.StdEncoding.EncodeToString(chiperText)

	return str, nil
}

func Decrypt(encrypted string, key []byte, iv []byte) (string, error) {
	chiperText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(chiperText)%aes.BlockSize != 0 {
		return "", errors.New("chiper text too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(chiperText, []byte(chiperText))

	unpadding := int(chiperText[len(chiperText)-1])
	chiperText = chiperText[:len(chiperText)-unpadding]

	if len(chiperText) == 0 {
		return "", errors.New("unable to encrypt data")
	}

	return string(chiperText), nil
}
