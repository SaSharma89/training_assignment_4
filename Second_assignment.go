/*
 * To explore panic, recover, error/errors
 * write a program using concept of error/errors handling, panic, recover with the help of godoc.
 */

package main

import (
	"fmt"
	"os"
	"errors"
)

func openFile(fileName string)( err error){
	fmt.Println("Inside openFile")

	src, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error while opening file : ", fileName, " with system error : ", err )
		err = fmt.Errorf("error while opeing file system error %w", err)
		return
	}

	defer func() {
		fmt.Println("defer : closing file")
		src.Close()
	}()

	fmt.Println("openFile over")

	return
}

func testPanic(){
	fmt.Println("testPanic start")

	panic("testing panic, execution must be stop")

	fmt.Println("testPanic over")
}

func testPanicRecover(){
	fmt.Println(" testPanicRecover start")

	defer func() {
		fmt.Println("testPanicRecover defer : try to recover")
		message := recover()
		fmt.Println("Panic message : ", message)
	}()

	var i int
	fmt.Println("Enter value 1/2 : ")
	fmt.Scanln(&i)

	switch i {
	case 1 :
		fmt.Println("Your input 1")
	case 2:
		fmt.Println("Your input 2")
	default:
		panic(fmt.Sprintf("panic occured due to input %d", i))
	}

	fmt.Println(" testPanicRecover over")
}

func main(){
	fmt.Println("Inside second assignment to explore error,panic and recover")

	fmt.Println("\n\n\nExplore errors:::::::::::")
	err := openFile("t3.txt")

	//! errors explore
	//! unwrap system err
	fmt.Println("system error : ", errors.Unwrap(err))

	//! check occurred error
	fmt.Println("error check : ", errors.Is(err, os.ErrNotExist))

	//! assign exist error
	var existErr *os.PathError
	errors.As(err, &existErr )
	fmt.Println("error as : ", existErr )

	//! panic or recover
	fmt.Println("\n\n\nExplore panic and recover:::::::::::")
	//testPanic()
	testPanicRecover()

	fmt.Println("main over")

}

