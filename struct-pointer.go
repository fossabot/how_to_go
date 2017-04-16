package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

/*
结构体字段可以通过结构体指针来访问。
通过指针间接的访问是透明的。
*/
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e2
	fmt.Println(v)
}
