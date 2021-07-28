package main

import (
	"fmt"
	"testing"
)


func callSlice(s []int){

}
func callArray(a [10000]int) {
	
}
func BenchmarkSlice(t *testing.B){
	s:=make([]int,10000)
	for i:=0;i<t.N;i++{
		callSlice(s)
	}
}

func BenchmarkArray(t *testing.B){
	var a [10000]int
	for i:=0;i<t.N;i++{
		callArray(a)
	}
}
func TestFmt(t *testing.T){
	fmt.Println("fmt")
	t.Log("log")
}