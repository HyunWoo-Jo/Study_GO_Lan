package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

//https://sourceforge.net/projects/mingw-w64/files/ 설치
//path C:\Program Files\mingw-w64\x86_64-8.1.0-posix-seh-rt_v6-rev0\mingw64\bin 추가
func main() {
	r := C.rand()

	fmt.Println(r)
}
