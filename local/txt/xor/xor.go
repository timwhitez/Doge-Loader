package main

import (
	"os"
	//"fmt"
	"strconv"
	"io/ioutil"
	"log"
	//"encoding/hex"
	//"fmt"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}

var XorKey []byte = []byte{0x11, 0x34, 0x67, 0x6A, 0xAA, 0xFF, 0x04, 0x7B}

type Xor struct {
}

type m interface {
	enc(src string) string
	dec(src string) string
}

func (a *Xor) enc(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + "\\x" + (s)
		j = (j + 1) % 8
	}
	return result
}

func (a *Xor) dec(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	//fmt.Println(bt)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}
	return result
}




func main(){
	xor := Xor{}

	//读取本地shellcode文件，异或成shellcode_xor
	fileObj, err := os.Open(os.Args[1]);
	bodyBytes, err := ioutil.ReadAll(fileObj)
	if err != nil {
		log.Fatal(err)
	}
	charcode := bodyBytes

	shellcodes := strings.ReplaceAll(string(charcode), "\n", "")
	shellcodes = strings.ReplaceAll(string(shellcodes), "\\x", "")

	xorCode := xor.enc(shellcodes)
	//fmt.Println(xorCode)
	ioutil.WriteFile("xor.txt", []byte(xorCode),0644)

}