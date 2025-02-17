// *EXERCISE 1*
/*package add

// Add is a function that takes in two integer values and returns their sum.
func Add(a, b int) int {
	return a + b
}*/

// Created repository
// Cloned the repo to my local machine using git clone https://github.com/IsaiahZoop/chap10.git
// Once code is on GitHub with the v1.0.0 tag, you can use it in another Go project with:
// go get github.com/IsaiahZoop/chap10@v1.0.0

// Run
//git commit -m "Added Add function"
//git tag v1.0.0
//git push origin main
//git push origin v1.0.0

// ---------------------------------------------------------------------------------------------

// *EXERCISE 2*
// Package chap10 provides basic math operations, including addition.
// More details can be found at https://www.mathsisfun.com/numbers/addition.html.
/*package add

// Add is a function that takes in two integer values and returns their sum.
// For more information, visit: https://www.mathsisfun.com/numbers/addition.html.
func Add(a, b int) int {
    return a + b
}*/

// Run
// git commit -m "Added godoc comments to Add function"
// git tag v1.0.1
// git push origin main
// git push origin v1.0.1

// ---------------------------------------------------------------------------------------------

// *EXERCISE 3*
// Package chap10 provides basic math operations with support for generics.
package add

import "golang.org/x/exp/constraints"

// Number is a generic type constraint that accepts both Integer and Float types.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two values of type Number and returns their sum.
func Add[T Number](a, b T) T {
	return a + b
}

// Run
// git commit -m "Refactored Add to be generic using Number interface"
// git tag v2.0.0
// git push origin main
// git push origin v2.0.0
