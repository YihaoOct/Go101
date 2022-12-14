// Package demo explores Go's basic syntax
package demo

import (
	"fmt"
	"strings"
)

// SayHello says hello :)
func SayHello() {
	fmt.Println("Hello, world!")
}

// SayHelloTo returns a greeting string to `name`.
func SayHelloTo(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}

// IsPrime tests whether a number is a prime number
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Primes gives a slice of prime numbers in [from, to)
func Primes(from, to int) []int {
	var results []int
	for i := from; i < to; i++ {
		if IsPrime(i) {
			results = append(results, i)
		}
	}
	return results
}

// CountWords returns the occurrences of `target` in text
func CountWords(text string, target string) int {
	occurrences := make(map[string]int)
	for _, word := range strings.Fields(text) {
		occurrences[word]++
	}
	return occurrences[target]
}

// The following code is similar to
// public class WordCounter : Counter
// {
//     private Dictionary<string, int> occurrences;
//     public WordCounter(string text) { ... }
//     public int GetCount(string target) { ... }
// }

// WordCounter counts occurrences of words in a string
type WordCounter struct {
	occurrences map[string]int
}

// NewWordCounter creates an instance of WordCounter based on a string
func NewWordCounter(text string) WordCounter {
	occurrences := make(map[string]int)
	for _, word := range strings.Fields(text) {
		occurrences[word]++
	}
	return WordCounter{
		occurrences: occurrences,
	}
}

// GetCount returns the number of occurrences of `target` in a WordCounter
func (w WordCounter) GetCount(target string) int {
	return w.occurrences[target]
}

// Counter retrieves the number of count of strings
type Counter interface {
	GetCount(string) int
}
