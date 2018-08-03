package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [4]int{1, 2, 3, 4}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	fmt.Println("p1 = ", *(*int)(p1))
	fmt.Println("p3 = ", *(*int)(p3))
	*(*int)(p3) = 6
	fmt.Println("p3 = ", *(*int)(p3))
}
