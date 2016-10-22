package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Sqrt(x float64) float64 {
	z, zplus := 1.0, 1.0
	for delta := 1.0; delta > 0.000001; z = zplus {
		zplus = z - (math.Pow(z, 2)-x)/(2*z)
		delta = math.Abs(zplus - z)
		fmt.Printf("%.10f -- %.10f\n", zplus, delta)
	}
	return z
}

func main() {
	for _, arg := range os.Args[1:] {
		arg_float, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatalf("strconv.ParseFloat: %v", err)
		}
		fmt.Printf("%.30f\n", Sqrt(arg_float))
	}
	if len(os.Args) == 1 {
		log.Fatalf("Nothing to square!")
	}
}
