package main

import ( // Importing packages that permits:
	"fmt"     // printing to screen
	"log"     // returning an error message if stringconversion fails
	"math"    // calculating power and absolute number functions
	"os"      // taking arguments from the command line
	"strconv" // converting a string to a float64
)

func Sqrt(x float64) float64 { // finding the square root of the value x
	// making a first guess z. Correcting the guess and saving it to zplus
	// and copying it back into z for a new guess.
	// Declaring and initializing the two variables. First guess is 1.
	z, zplus := 1.0, 1.0
	// Loop until delta is below a threshold.
	for delta := 1.0; delta > 0.000001; z = zplus {
		zplus = z - (math.Pow(z, 2)-x)/(2*z)
		delta = math.Abs(zplus - z)
		fmt.Printf("%.10f -- %.10f\n", zplus, delta)
	}
	return z
}

func main() {
	// Skipping the first argument (path), then for each argument
	// convert it to a float64 and run it through the Sqrt function.
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
