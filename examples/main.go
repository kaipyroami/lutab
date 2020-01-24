package main

import (
	"flag"
	"fmt"
)

func main() {

	// Note: (name, default val, help message)
	file1Ptr := flag.String("table-one", "table1.csv", "The first table input.")
	file2Ptr := flag.String("table-two", "table2.csv", "The second table input.")
	multiplierPtr := flag.Float64("multiplier", 1.0, "The multiplier that the first table is scaled by.")

	wordPtr := flag.String("", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
