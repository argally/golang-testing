package main

import (
	"io"
	"os"
	"testing"
)

// Test_isPrime tests the isPrime function to determine if a number is prime.
// It iterates over a set of predefined test cases, each containing a name,
// a test number, the expected result, and a message. The function checks
// if the result from isPrime matches the expected result and reports an
// error if there is a mismatch.
func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not prime", 8, false, "8 is not a prime number as it is divisible by 2"},
		{"negative", -1, false, "-1 is a negative number and not a prime"},
		{"zero", 0, false, "0 is not a prime number"},
	}

	for _, entry := range primeTests {

		result, msg := isPrime(entry.testNum)
		if entry.expected && !result {
			t.Errorf("Test case %s failed: expected true, got false", entry.name)
		}

		if !entry.expected && result {
			t.Errorf("Test case %s failed: expected false, got true", entry.name)
		}

		if entry.msg != msg {
			t.Errorf("Test case %s failed: expected %s, got %s", entry.name, entry.msg, msg)
		}
	}
}

func Test_prompt(t *testing.T) {
	//save a copy of os.stdout
	oldOut := os.Stdout
	//This line creates a pipe using `os.Pipe()`,
	// which returns a pair of file descriptors: `r` (read end) and `w` (write end).
	// This pipe will be used to capture the output of the `prompt` function.
	r, w, _ := os.Pipe()
	//set os.stdout to the write pipe
	os.Stdout = w
	prompt()
	//close the write pipe
	_ = w.Close()
	//reset os.stdout to what it was before
	os.Stdout = oldOut
	//This line reads all the data from the read end of the pipe (`r`) using `io.ReadAll`.
	// The captured output is stored in the variable `out`.
	out, _ := io.ReadAll(r)
	//check if the output is as expected
	if string(out) != "--> " {
		t.Errorf("Incorrect prompt: expected --> but got %s", string(out))
	}

}
