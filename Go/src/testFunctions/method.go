package main

import "fmt"

type Rect struct {
	w, h float64
}

func (r Rect) Area() float64 {
	return r.w * r.h
}

func main() {
	fmt.Println("test method")
	r := Rect{3, 5}
	area := r.Area()
	fmt.Println(area)
}
