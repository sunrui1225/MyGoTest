package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func modifySlice(s []int) {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := *(*[3]int)(unsafe.Pointer(hdr.Data))
	fmt.Println("modifySlice_hdr", hdr, "modifySlice_data", data)

	s = append(s, 2048)

	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data2 := *(*[4]int)(unsafe.Pointer(hdr2.Data))
	fmt.Println("modifySlice_hdr_append4", hdr2, "modifySlice_data_append4", data2)

	//s = append(s, 2049)
	//hdr5 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	//data5 := *(*[5]int)(unsafe.Pointer(hdr5.Data))
	//fmt.Println("modifySlice_hdr_append5",hdr5, "modifySlice_data_append5",data5)

	s[0] = 1024

	hdr3 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data3 := *(*[4]int)(unsafe.Pointer(hdr3.Data))
	fmt.Println("modifySlice_hdr_1024", hdr3, "modifySlice_data_1024", data3)
}

func main() {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := *(*[3]int)(unsafe.Pointer(hdr.Data))
	fmt.Println("main_hdr", hdr, "main_data", data)

	modifySlice(s)

	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data2 := *(*[3]int)(unsafe.Pointer(hdr2.Data))
	fmt.Println("main_hdr_end", hdr2, "main_data_end", data2)
}

// 0,1,2,2048
