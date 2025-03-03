package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
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

func Test_readUserInput(t *testing.T) {
	// to test this fucntion we need a channel
	// we also need instance of io.Reader
	doneChan := make(chan bool)
	//create reference to bytes.Buffer and assign it to stdin
	//bytes.Buffer implements the io.Reader interface
	var stdin bytes.Buffer
	// simulate user typing 1 then return then q and return
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)

}

func Test_checkNumbers(t *testing.T) {

	checknumberTests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a whole number"},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
		{"prime", "7", "7 is a prime number"},
		{"not prime", "8", "8 is not a prime number as it is divisible by 2"},
		{"negative", "-1", "-1 is a negative number and not a prime"},
		{"zero", "0", "0 is not a prime number"},
		{"invalid", "three", "Please enter a whole number"},
		{"decimal", "3.14", "Please enter a whole number"},
		{"greek", "εβδομάδα", "Please enter a whole number"},
	}

	for _, entry := range checknumberTests {
		//bufio.NewScanner takws an io.Reader as an argument
		input := strings.NewReader(entry.input)
		// so we can use strings.NewReader which satisfies the io.Reader interface
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold((res), entry.expected) {
			t.Errorf("%s expected %s but got %s", entry.name, entry.expected, res)
		}
	}
}

func Test_intro(t *testing.T) {
	//save a copy of os.stdout
	oldOut := os.Stdout
	//This line creates a pipe using `os.Pipe()`,
	// which returns a pair of file descriptors: `r` (read end) and `w` (write end).
	// This pipe will be used to capture the output of the `prompt` function.
	r, w, _ := os.Pipe()
	//set os.stdout to the write pipe
	os.Stdout = w
	intro()
	//close the write pipe
	_ = w.Close()
	//reset os.stdout to what it was before
	os.Stdout = oldOut
	//This line reads all the data from the read end of the pipe (`r`) using `io.ReadAll`.
	// The captured output is stored in the variable `out`.
	out, _ := io.ReadAll(r)
	//check if the output is as expected
	if !strings.Contains(string(out), "Is it prime?") {
		t.Errorf("Incorrect intro text not correct %s", string(out))
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
