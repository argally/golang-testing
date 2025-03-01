package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter got true expected false", 0)

	}
	if msg != "0 is not a prime number" {
		t.Errorf("wrong message returned: %s", msg)
	}

}
