package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}


func main() {
	s := make([]int, 0)
	oldCap := cap(s)
	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("before cap = %-4d  |  after append %-4d  cap = %-4d\n",  oldCap, i, newCap)
			oldCap = newCap
		}
	}
}
