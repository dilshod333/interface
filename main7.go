package main

import "fmt"

type Bike struct {
	brand string
	model string
}
type bike interface {
	start()
	stop()
}

func (bike Bike) start(a string) string {
	return a
}

func (bike Bike) stop(b int) int {
	return b
}

func main() {
	bike := Bike{
		brand: "star",
		model: "x7z",
	}
	c, d := "Dilshod", 333
	a := bike.start(c)
	b := bike.stop(d)

	fmt.Println(a, b)
}
