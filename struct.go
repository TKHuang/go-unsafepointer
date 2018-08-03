package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Car struct {
		brand string
		power int
		color string
	}

	car := Car{"Nelix", 6000, "Yellow"}
	cp := unsafe.Pointer(&car)
	cpb := (*string)(unsafe.Pointer(uintptr(cp) + unsafe.Offsetof(car.brand)))
	cpp := (*int)(unsafe.Pointer(uintptr(cp) + unsafe.Offsetof(car.power)))
	cpc := (*string)(unsafe.Pointer(uintptr(cp) + unsafe.Offsetof(car.color)))
	fmt.Println("Before...")
	fmt.Println("brand =", *cpb)
	fmt.Println("power =", *cpp)
	fmt.Println("color =", *cpc)
	*cpb = "GTCQ"
	*cpp = 60000
	*cpc = "Red"
	fmt.Println("After...")
	fmt.Println("brand =", *cpb)
	fmt.Println("power =", *cpp)
	fmt.Println("color =", *cpc)
}
