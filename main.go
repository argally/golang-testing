package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Print a welcome message
	intro()
	//create a channel to indicate the end of the program
	doneChan := make(chan bool)
	//start a gorouting to read user input and run program
	go readUserInput(doneChan)
	//block until the doneChan gets a value
	<-doneChan
	//close channel and say goodbye
	close(doneChan)
	fmt.Println("Goodbye!")

}

// readUserInput reads user input from the standard input (stdin) and sends a signal
// to the provided doneChan channel when the reading is complete.

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	// **Exit Condition**: The `checkNumbers` function returns two values:
	// a result message (`res`) and a boolean (`done`). If `done` is `true`,
	// it means the user has entered "q" to quit.
	// The goroutine then sends a `true` signal to the `doneChan` channel and exits the loop,
	// effectively ending the goroutine.

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}
		// **Output and Prompt**: If `done` is `false`, it prints the result message (`res`)
		//  and prompts the user for more input.
		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// **Reading User Input**: The function starts by reading a line of input from the user
	// using the `scanner.Scan()` method.
	// This method reads the next token from the input
	// and makes it available via `scanner.Text()`.

	scanner.Scan()
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}
	//*Input Conversion**: The function attempts to convert the user input
	// from a string to an integer using `strconv.Atoi()`.
	numToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number", false
	}

	_, msg := isPrime(numToCheck)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// isPrime determines if a given integer is a prime number.
	// It returns a boolean indicating whether the number is prime,
	// and a string message explaining the result.
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

func intro() {
	fmt.Println("Is it prime?")
	fmt.Println("------------")
	fmt.Println("Enter a whole number, and we'll tell you if it's prime. Enter q to quit.")
	prompt()
}

func prompt() {
	fmt.Print("--> ")
}
