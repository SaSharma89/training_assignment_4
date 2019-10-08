/*
 * To explore defer
 * Use multiple defer with print statement.
 */

package main

import (
	"fmt"
	"os"
	"io"
)

func createFile( filename string){
	fmt.Println("createFile start")

	file, err := os.Create(filename)
	if err != nil{
		fmt.Println("Error while creating file with name ", filename)
		return
	}

	defer func() {
		fmt.Println("createFile defer : going to close file")
		if file != nil {
			file.Close()
		}
	}()

	fmt.Println("createFile end")
}


func writeFile( filename string){
	fmt.Println("writeFile start")

	src, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error while opening a file")
		return
	}

	defer func() {
		fmt.Println("writeFile Defer : Going to close file")
		if src != nil {
			src.Close()
		}
	}()

	message := "This is test string to write in file"
	fmt.Fprintf(src, "%s\n", message)

	fmt.Println("writeFile end")
}

func copyFile( srcFile, dstFile string){
	fmt.Println("copyFile start")

	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("Error while opening a file")
		return
	}

	defer func() {
		fmt.Println("copyFile Defer : Going to close src file")
		if src != nil {
			src.Close()
		}
	}()

	//! create dst file
	createFile(dstFile)

	dst, err := os.OpenFile(dstFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println("Error while opening a file")
		return
	}

	defer func() {
		fmt.Println("copyFile Defer : Going to close dst file")
		if dst != nil {
			dst.Close()
		}
	}()

	byteCopyied, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error while coping a file")
		return
	}

	fmt.Println("copyFile end, with byteCopyied :", byteCopyied)
}

func main() {
	fmt.Println("First program to explore defer")

	createFile("t.txt")

	writeFile("t.txt")

	copyFile("t.txt", "t2.txt")

	defer fmt.Println("main Defer : calling defer for main")

	fmt.Println("main is over")
}