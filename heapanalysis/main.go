package main

type mystruct struct {
	a int
	b string
}

func main() {
	//fmt.Println("Called heapAnalysis", heapAnalysis())
	m := make([]int, 0, 2)
	//heapAnalysis(&m)
	//fmt.Println(m)
	m[0] = 1
}

//go:noinline
func heapAnalysis(m *[]int) *int {
	data := 55
	return &data
}
