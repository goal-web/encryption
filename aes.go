package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/goal-web/contracts"
	"github.com/pkg/errors"
)

type aesEncryptor struct {
	key   []byte
	block cipher.Block
}

func AES(key string) contracts.Encryptor {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)

	if err != nil {
		panic(EncryptException{Err: err})
	}

	return &aesEncryptor{key: keyBytes, block: block}
}

func (aes *aesEncryptor) Encrypt(origData []byte) []byte {
	// 获取秘钥块的长度
	blockSize := aes.block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(aes.block, aes.key[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(encrypted, origData)

	return []byte(base64.StdEncoding.EncodeToString(encrypted))
}

func (aes *aesEncryptor) EncryptString(value string) string {
	// 转成字节数组
	origData := []byte(value)
	// 获取秘钥块的长度
	blockSize := aes.block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(aes.block, aes.key[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(encrypted, origData)

	return base64.StdEncoding.EncodeToString(encrypted)
}

func (aes *aesEncryptor) DecryptString(encrypted string) (result string, err error) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			switch value := panicValue.(type) {
			case error:
				err = value
			default:
				err = errors.Errorf("%v", value)
			}
		}

	}()
	// 转成字节数组
	encryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// 获取秘钥块的长度
	blockSize := aes.block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(aes.block, aes.key[:blockSize])
	// 创建数组
	orig := make([]byte, len(encryptedByte))
	// 解密
	blockMode.CryptBlocks(orig, encryptedByte)
	// 去补全码
	return string(PKCS7UnPadding(orig)), nil
}

func (aes *aesEncryptor) Decrypt(encrypted []byte) (result []byte, err error) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			switch value := panicValue.(type) {
			case error:
				err = value
			default:
				err = errors.Errorf("%v", value)
			}
		}

	}()
	// 转成字节数组
	encryptedByte, err := base64.StdEncoding.DecodeString(string(encrypted))
	if err != nil {
		return nil, err
	}

	// 获取秘钥块的长度
	blockSize := aes.block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(aes.block, aes.key[:blockSize])
	// 创建数组
	orig := make([]byte, len(encryptedByte))
	// 解密
	blockMode.CryptBlocks(orig, encryptedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return orig, nil
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
