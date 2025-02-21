package main

import ("fmt")

func main() {
	var a []int = []int{1, 2, 3}
	var b []int = a

	fmt.Println(a, b)
	fmt.Println(&a, &b)


	v:=make([]int,3)
	v[0]=1
	v[1]=2
	v[2]=3
	fmt.Println(v)


	//slice
	//append slice

	
}

