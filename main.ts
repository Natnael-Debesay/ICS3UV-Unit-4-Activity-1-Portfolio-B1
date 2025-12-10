/**
 * @author Natnael Debesay
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program prints a 10 by 10 grid of X's.
 */

// variables
let xs = "";

// Loop through the X's 1 to 10 for the rows
for (let rowVariable = 1; rowVariable <= 10; rowVariable++) {
  // Loop through the X's 1 to 10 for the columns
  for (let columnVariable = 1; columnVariable <= 10; columnVariable++) {
    // Calculate the product and format it to align properly
    xs += 'X ';
  }
// Print the completed row
console.log(xs);
// Reset the row variable for the next iteration
xs = "";
}

console.log("\nDone.");
