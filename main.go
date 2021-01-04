package main

import (
	"fmt"
)

func main() {
	dev := DeviceNew("B2027E6DA891CBA1E658291398D082127DBBF993")
	if dev == nil {
		fmt.Println("dev nil")
		return
	}
	lckd := LockdownClientNewWithHandshake(dev)
	if lckd == nil {
		fmt.Println("lckd nil")
	}
	str := LockdownGetValue(lckd, "")
	fmt.Println(str)
}
