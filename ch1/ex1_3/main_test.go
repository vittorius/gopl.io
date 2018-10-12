// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments. Must be run as `go test -bench=. -args 1 2 3`
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func printIterate() {
	os.Stdout,_ = os.Open(os.DevNull)
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func printRange() {
	os.Stdout,_ = os.Open(os.DevNull)
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func printJoin() {
	os.Stdout,_ = os.Open(os.DevNull)
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// BenchmarkPrintIterate Benchmarking for printIterate function
func BenchmarkPrintIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printIterate()
	}
}

// BenchmarkRange Benchmarking for printRange function
func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printRange()
	}
}

// BenchmarkJoin Benchmarking for printJoin function
func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printJoin()
	}
}

//!+
func main() {

}

//!-
