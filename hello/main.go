package main

import (
	"fmt"
	"runtime"
)

var x int
var y string
var z bool

type x1 int

const (
	a = iota << 1
	b
	c
	d
	e
)

func main() {
	x = 42
	y = "James bond"
	z = true
	var y1 x1
	y1 = 42
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Print(s)
	fmt.Printf("\n%T %v\n", y1, y1)
	x = int(y1)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("\n%v %v\n", a, d)
	somehing()
	learnslice()
	learnmap()
	return
}

func somehing() {
	var x [5]int

	x[2] = 1
	fmt.Println(x)
}

func learnslice() {
	x := []mystruct{{1, "a"}, {2, "b"}, {3, "c"}}
	fmt.Println(x[:], len(x))
	y := []mystruct{{4, "a"}, {5, "b"}, {6, "c"}}
	//x = append(x, y...)
	for key, val := range x {
		fmt.Println(key)
		fmt.Println(val)
	}
	//x = append(x[:2], x[4:]...)

	z := [][]mystruct{x, y}
	fmt.Println(z)
}

type mystruct struct {
	data int
	name string
}

func learnmap() {
	m := map[string]mystruct{
		"a": {1, "a1"},
		"b": {2, "b1"},
	}
	m["c"] = mystruct{3, "c1"}
	fmt.Println(m)

}
