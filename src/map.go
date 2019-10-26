package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

func main() {
	var m = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
	var m2 = make(map[string]int)
	m2["Lucknow"] = 1
	m2["Bengaluru"] = 2
	fmt.Println(m2["Lucknow"])
	fmt.Println(m2["Bengaluru"])

	m3 := make(map[string]int)
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}