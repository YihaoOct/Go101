package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSayHelloTo(t *testing.T) {
	want := "Hello, Golang\n"
	got := SayHelloTo("Golang")

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func ExampleSayHelloTo() {
	greetings := SayHelloTo("Jimmy")
	fmt.Println(greetings)
	// output:
	// Hello, Jimmy
}

func TestIsPrime(t *testing.T) {
	cases := map[int]bool{
		-1: false,
		0:  false,
		1:  false,
		2:  true,
		3:  true,
		4:  false,
		5:  true,
	}

	for input, want := range cases {
		got := IsPrime(input)

		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestPrimes(t *testing.T) {
	from, to := 0, 10
	want := []int{2, 3, 5, 7}
	got := Primes(from, to)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCountWords(t *testing.T) {
	text := "Hello world"
	cases := map[string]int{
		"Hello": 1,
		"world": 1,
		"magic": 0,
	}

	for target, want := range cases {
		if got := CountWords(text, target); got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestWordCounter(t *testing.T) {
	counter := NewWordCounter("Hello world")
	cases := map[string]int{
		"Hello": 1,
		"world": 1,
		"magic": 0,
	}

	for target, want := range cases {
		if got := counter.GetCount(target); got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}

func TestCounter(t *testing.T) {
	var counter Counter
	counter = NewWordCounter("Hello world")
	cases := map[string]int{
		"Hello": 1,
		"world": 1,
		"magic": 0,
	}

	if _, ok := counter.(Counter); !ok {
		t.Errorf("expected counter to be of type Counter, but it is not")
	}

	for target, want := range cases {
		if got := counter.GetCount(target); got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	}
}
