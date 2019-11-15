package common

import "fmt"

var (
	RegisterUrl = ""
	SealResulUrl = ""
	PostResultUrl = ""
)
//初始化url
func InitUrl(address string){
	address = fmt.Sprintf("http://%s/", address)
	RegisterUrl = address + "register"
	SealResulUrl = address + "seal_result"
	PostResultUrl = address + "post_result"
}

