package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := [24]float64{20.1, 24, 27.3, 30.1, 26.4, 22.2, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3, 30.1, 26.4, 20.1, 24, 27.3}
	//Insert Code here
	var total float64

	for _, temp := range temperature {
		total += temp
	}
	// avgTemp := total/24
	fmt.Println(math.Round(total/24*10)/10)
}
