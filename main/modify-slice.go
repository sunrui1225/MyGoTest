package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func modifySlice(s []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("modifySlice_hdr", hdr)
	data := *(*[3]int)(unsafe.Pointer(hdr.Data))
	fmt.Println("modifySlice_data", data)

	s = append(s, 2048)

	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("modifySlice_hdr_append4", hdr2)
	data2 := *(*[4]int)(unsafe.Pointer(hdr2.Data))
	fmt.Println("modifySlice_data_append4", data2)

	//s = append(s, 2049)
	//hdr5 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	//fmt.Println("modifySlice_hdr_append5",hdr5)
	//data5 := *(*[5]int)(unsafe.Pointer(hdr5.Data))
	//fmt.Println("modifySlice_data_append5",data5)

	s[0] = 1024

	hdr3 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("modifySlice_hdr_1024", hdr3)
	data3 := *(*[5]int)(unsafe.Pointer(hdr3.Data))
	fmt.Println("modifySlice_data_1024", data3)
}

func main() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("main_hdr", hdr)
	data := *(*[3]int)(unsafe.Pointer(hdr.Data))
	fmt.Println("main_data", data)

	modifySlice(s)

	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("main_hdr_end", hdr2)
	data2 := *(*[3]int)(unsafe.Pointer(hdr2.Data))
	fmt.Println("main_data_end", data2)

	fmt.Println(s)
}

// 0,1,2,2048
