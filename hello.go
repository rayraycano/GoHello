package main

import (
	"fmt"

	"github.com/rayraycano/stringutil"
)

func main() {
	s := "Ni Hao bruh bruh"
    fmt.Println(s)
    fmt.Println("Missy Elliot Says: 'REVERSE!'")
    fmt.Println(stringutil.Reverse(s))

    m := stringutil.RuneCount(s)
    fmt.Println("\nCounts:", m)
    // fmt.Printf(stringutil.RuneCount(s))
}
