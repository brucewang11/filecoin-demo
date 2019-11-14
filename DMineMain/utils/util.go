package utils

import (

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
