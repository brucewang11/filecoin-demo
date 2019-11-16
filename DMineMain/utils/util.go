package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"strconv"
)

//验证port
func VerifyPort(port string) bool{
	intPort, err := strconv.Atoi(port)
	if err!=nil {
		return false
	}
	if intPort<=0 || (intPort>65535) {
		return false
	}
	return true

}
//随机key
func GenerateKey(bits int) string{
	private, _ := rsa.GenerateKey(rand.Reader, bits)
	derStream := x509.MarshalPKCS1PrivateKey(private)
	str := hex.EncodeToString(derStream)
	return str
}
