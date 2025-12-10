// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-12-09
// Fileoverview: This program prints a 10 by 10 grid of X's.

package main

import "fmt"

func main() {
	// variables
	var xs string

	// Loop through the X's 1 to 10 for the rows
	for rowVariable := 1; rowVariable <= 10; rowVariable++ {
		// Loop through the X's 1 to 10 for the columns
		for columnVariable := 1; columnVariable <= 10; columnVariable++ {
			// Calculate the product and format it to align properly
			xs += "X   "
		}
		// Print the completed row
		fmt.Println(xs)
		// Reset the row variable for the next iteration
		xs = ""
	}

	fmt.Println("\nDone.")
}
