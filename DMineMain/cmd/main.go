package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"github.com/dMineMain/hashring"
)

var rands = 100000

//测试一致性哈希
func main() {
	cHashRing := hashring.NewConsistent()
	cHashRing.Add(hashring.NewNode(1, "172.18.1.221",0, "", 1))
	cHashRing.Add(hashring.NewNode(2, "172.18.1.222", 0, "", 1))
	cHashRing.Add(hashring.NewNode(3, "172.18.1.223", 0, "", 1))
	cHashRing.Add(hashring.NewNode(4, "172.18.1.224", 0, "", 1))
	cHashRing.Add(hashring.NewNode(5, "172.18.1.225", 0, "", 1))

	maps := make(map[string]int)
	for i:=0;i<10000 ;i++  {

		aa := GenerateKey(128)
		node := cHashRing.Get(aa)
		maps[node.Ip] = maps[node.Ip] + 1


	}
	for k,v:= range maps{
		fmt.Println(k,v)
	}
}
func GenerateKey(bits int) string{
	private, _ := rsa.GenerateKey(rand.Reader, bits)
	derStream := x509.MarshalPKCS1PrivateKey(private)
	str := hex.EncodeToString(derStream)
	return str
}



