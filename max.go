package findmax

import (
	"reflect"
	"fmt"
	"sort"
)


func rrr(i,j int) {
   sort.Slice([]int{1,2,3}, func(i, j int) bool{
	   return true
   })
}

// FindMax comporator
func FindMax(slice interface{},less func(i, j int) bool) interface{} {

	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	swap(1,2)
	length := rv.Len()
	fmt.Println("rv=>",rv)
	fmt.Println("len=>",length)

	// var x float64 = 3.4
	// fmt.Println("type:", reflect.TypeOf(x))
	// v:=reflect.ValueOf(x)
	// fmt.Println("value:", v.String())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println(x == v.Float())

   
	return true


	
	
	
}