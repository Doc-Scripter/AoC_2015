package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Hex(input string) string {
	// input:=str
	// 999999
	for i :=0;i<10000000 ; i++ {
		// num:= 
		str:=input+strconv.Itoa(i)
		hash := md5.Sum([]byte(str))
		// h := md5.New()
		// h.Write(hash)
		out := fmt.Sprintf("%x", hash)
		fmt.Println(out)
		if strings.HasPrefix(out,"00000"){
			return str
		}
	}
return ""
	
}

func main() {
	input := "bgvyzdsv"
	result:=Hex(input)
	fmt.Println(result)
}
