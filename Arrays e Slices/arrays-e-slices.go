package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Aline"
	array1[1] = "Cristine"
	array1[2] = "Barbosa"
	array1[3] = "Fabio"
	array1[4] = "Carito"
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{20, 30, 40, 50, 60, 70}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 80)
	fmt.Println(slice)

	slice2 := array2[1:2]
	fmt.Println(slice2)

	//Arrays Internos
	fmt.Println("><><><><><")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //lenght
	fmt.Println(cap(slice3)) //capacidade
}
