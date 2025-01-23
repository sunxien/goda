package functions

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/sm4"
	"io"
)

// MD5
/*
 * Usage: MD5("ABCabc123")
 * Output: 39a9d01af60d5c6dcd94b61f0f2e086d
 */
func MD5(s *string) *String {
	if s == nil {
		return NullString()
	}
	hash := md5.New()
	hash.Write([]byte(*s))
	return NonNullString(hex.EncodeToString(hash.Sum(nil)))
}

// HMAC
/*
 * Usage: HMAC("ABCabc123", "123456")
 * Output:
 */
func HMAC(s *string, key string) *String {
	if s == nil {
		return NullString()
	}
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(*s))
	return NonNullString(hex.EncodeToString(mac.Sum(nil)))
}

// SHA1
/*
 * Usage: SHA1("ABCabc123")
 * Output:
 */
func SHA1(s *string) *String {
	if s == nil {
		return NullString()
	}
	hash := sha1.New()
	hash.Write([]byte(*s))
	return NonNullString(hex.EncodeToString(hash.Sum(nil)))
}

// SHA256
/*
 * Usage: SHA256("ABCabc123")
 * Output:
 */
func SHA256(s *string) *String {
	if s == nil {
		return NullString()
	}
	hash := sha256.New()
	hash.Write([]byte(*s))
	return NonNullString(hex.EncodeToString(hash.Sum(nil)))
}

// HEX
/*
 * Usage: HEX("ABCabc123")
 * Output:
 */
func HEX(s *string) *String {
	if s == nil {
		return NullString()
	}
	return NonNullString(hex.EncodeToString([]byte(*s)))
}

// UNHEX
/*
 * Usage: UNHEX("ABCabc123")
 * Output:
 */
func UNHEX(s *string) *String {
	if s == nil {
		return NullString()
	}
	b, err := hex.DecodeString(*s)
	if err != nil {
		panic("UNHEX.DecodeString failed. Error: " + err.Error())
	}
	return NonNullString(string(b))
}

// BASE64Encode
/*
 * Usage: BASE64Encode("ABCabc123")
 * Output: QUJDYWJjMTIz
 */
func BASE64Encode(s *string) *String {
	if s == nil {
		return NullString()
	}
	return NonNullString(base64.StdEncoding.EncodeToString([]byte(*s)))
}

// BASE64Decode
/*
 * Usage: BASE64Decode("QUJDYWJjMTIz")
 * Output: ABCabc123
 */
func BASE64Decode(s *string) *String {
	if s == nil {
		return NullString()
	}
	b, err := base64.StdEncoding.DecodeString(*s)
	if err != nil {
		panic("BASE64Decode.DecodeString failed. Error: " + err.Error())
	}
	return NonNullString(string(b))
}

// AESEncrypt
/*
 * Key length: 16, 24, 32 => AES-128, AES-192, AES-256
 * Usage: AESEncrypt("ABCabc123", "0123456789abcdef")
 * Output: Cr+4qKl1RuFLrxstfQpwuxyOhN5bCinbY/LaNbPYoC0=
 */
func AESEncrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := aes.NewCipher([]byte(k))
	if err != nil {
		panic("AESEncrypt.NewCipher failed. Error: " + err.Error())
	}
	bs := b.BlockSize()
	data := PKCS7Padding([]byte(*s), bs)
	buf := make([]byte, bs+len(data))
	iv := buf[:bs]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("AESEncrypt.ReadFull failed. Error: " + err.Error())
	}
	cipher.NewCBCEncrypter(b, iv).CryptBlocks(buf[bs:], data)
	return NonNullString(hex.EncodeToString(buf))
}

// AESDecrypt
/*
 * Key length: 16, 24, 32 => AES-128, AES-192, AES-256
 * Usage: AESEncrypt("Cr+4qKl1RuFLrxstfQpwuxyOhN5bCinbY/LaNbPYoC0=", "0123456789abcdef")
 * Output: ABCabc123
 */
func AESDecrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := aes.NewCipher([]byte(k))
	if err != nil {
		panic("AESDecrypt.NewCipher failed. Error: " + err.Error())
	}
	bs := b.BlockSize()
	if len(*s) < bs {
		panic("Cipher text is too short")
	}
	data, err := hex.DecodeString(*s)
	if err != nil {
		panic("AESDecrypt.DecodeString failed. Error: " + err.Error())
	}
	iv := data[:bs]
	data = data[bs:]
	if len(data)%bs != 0 {
		panic("Cipher text is not a multiple of block size")
	}
	buf := make([]byte, len(data))
	cipher.NewCBCDecrypter(b, iv).CryptBlocks(buf, data)
	return NonNullString(string(PKCS7Unpadding(buf)))
}

// DESEncrypt
/*
 * Usage: DESEncrypt('ABCabc123')
 * Output:
 */
func DESEncrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := des.NewCipher([]byte(k))
	if err != nil {
		panic("DESEncrypt.NewCipher failed. Error: " + err.Error())
	}
	bs := b.BlockSize()
	data := PKCS7Padding([]byte(*s), bs)
	buf := make([]byte, bs+len(data))
	iv := buf[:bs]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("AESEncrypt.ReadFull failed. Error: " + err.Error())
	}
	cipher.NewCBCEncrypter(b, iv).CryptBlocks(buf[bs:], data)
	return NonNullString(hex.EncodeToString(buf))
}

// DESDecrypt
/*
 * DESDecrypt('ABCabc123', '01234567')
 * Output:
 */
func DESDecrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := des.NewCipher([]byte(k))
	if err != nil {
		panic("DESDecrypt.NewCipher failed. Error: " + err.Error())
	}
	bs := b.BlockSize()
	if len(*s) < bs {
		panic("Cipher text is too short")
	}
	data, err := hex.DecodeString(*s)
	if err != nil {
		panic("DESDecrypt.DecodeString failed. Error: " + err.Error())
	}
	iv := data[:bs]
	data = data[bs:]
	if len(data)%bs != 0 {
		panic("Cipher text is not a multiple of block size")
	}
	buf := make([]byte, len(data))
	cipher.NewCBCDecrypter(b, iv).CryptBlocks(buf, data)
	return NonNullString(string(PKCS7Unpadding(buf)))
}

// SM3
/*
 * Usage: SM3("ABCabc123")
 * Output:
 */
func SM3(s *string) *String {
	if s == nil {
		return NullString()
	}
	h := sm3.New()
	h.Write([]byte(*s))
	return NonNullString(hex.EncodeToString(h.Sum(nil)))
}

// SM4Encrypt
/*
 * Usage: SM4Encrypt("ABCabc123", "0123456789abcdef")
 * Output:
 */
func SM4Encrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := sm4.NewCipher([]byte(k))
	if err != nil {
		panic("SM4Encrypt.NewCipher failed. Error: " + err.Error())
	}
	origin := PKCS7Padding([]byte(*s), b.BlockSize())
	iv := make([]byte, sm4.BlockSize)
	buf := make([]byte, len(origin))
	cipher.NewCBCEncrypter(b, iv).CryptBlocks(buf, origin)
	return NonNullString(hex.EncodeToString(buf))
}

// SM4Decrypt
/*
 * Usage: SM4Decrypt('ABCabc123')
 * Output:
 */
func SM4Decrypt(s *string, k string) *String {
	if s == nil {
		return NullString()
	}
	b, err := sm4.NewCipher([]byte(k))
	if err != nil {
		panic("SM4Decrypt.NewCipher failed. Error: " + err.Error())
	}
	data, err := hex.DecodeString(*s)
	if err != nil {
		panic("SM4Decrypt.DecodeString failed. Error: " + err.Error())
	}
	iv := make([]byte, b.BlockSize())
	buf := make([]byte, len(data))
	cipher.NewCBCDecrypter(b, iv).CryptBlocks(buf, data)
	origin := PKCS7Unpadding(buf)
	return NonNullString(string(origin))
}

// PKCS5Padding
/*
 *
 */
func PKCS5Padding(src []byte, bs int) []byte {
	padding := bs - len(src)%bs
	paddingtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingtext...)
}

// PKCS5Unpadding
/*
 *
 */
func PKCS5Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// PKCS7Padding
/*
 *
 */
func PKCS7Padding(src []byte, bs int) []byte {
	padding := bs - len(src)%bs
	paddingtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, paddingtext...)
}

// PKCS7Unpadding
/*
 *
 */
func PKCS7Unpadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
