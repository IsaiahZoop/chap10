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
// | (pipe operator): Allows you to combine multiple constraints.
// It means that the Number interface accepts types that either match integer or float here.
type Number interface {
	constraints.Integer | constraints.Float
}

// This function uses Go generics to work with different types (integers and floats).
// It ensures that only integer and floating-point types can be passed as arguments, thanks to the Number interface.
// Add is a generic function that takes in two parameters of type T,
// where T must satisfy the Number interface (either an integer or a float).
func Add[T Number](a, b T) T {
	// The function returns the sum of a and b, both of type T.
	return a + b
}

// Run
// git commit -m "Refactored Add to be generic using Number interface"
// git tag v2.0.0
// git push origin main
// git push origin v2.0.0

// To remove tag localled: git tag -d v2.0.0
// To remove tag from GitHub remotely: git push --delete origin v2.0.0
// To re-add tag: git tag v2.0.0
// To push tag: git push origin v2.0.0

// To push everything from this directory run...
/*
git add .
git commit -m "Pushed all updates"
git push origin main
git push --tags
*/
