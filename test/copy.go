package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args

	if len(list) != 3 {
		fmt.Println("usage: xxx srcFile dstFile")
		return
	}

	srcFile := list[1]
	dstFile := list[2]

	if srcFile == dstFile {
		fmt.Println("源文件和目的文件名称不能相同")
		return
	}

	sF, err1 := os.Open(srcFile)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	dF, err2 := os.Create(dstFile)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
	}

	defer sF.Close()
	defer dF.Close()

	buf := make([]byte, 4*1024)
	
	for {
        n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF { 
				break
			}
			fmt.Println("err = ", err)
		}
		dF.Write(buf[:n])
	}

}
