package main

/*
#cgo windows,amd64 CFLAGS: -D__USE_MINGW_ANSI_STDIO=1
#include <stdlib.h>

long my_random() {
    return rand();
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("corss-build")
	fmt.Println("c random:", C.my_random())
}
