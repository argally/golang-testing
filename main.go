package main

import "fmt"

func main() {
	n := 7
	_, msg := isPrime(n)
	fmt.Println(msg)
}

// isPrime determines if a given integer is a prime number.
// It returns a boolean indicating whether the number is prime,
// and a string message explaining the result.
//
// Parameters:
//   n - the integer to be checked.
//
// Returns:
//   bool - true if the number is prime, false otherwise.
//   string - a message explaining the result.

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not a prime number", n)
	}
	if n < 0 {
		return false, fmt.Sprintf("%d is a negative number and not a prime", n)
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number as it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)

}
