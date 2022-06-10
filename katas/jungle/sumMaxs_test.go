package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSumMaxsQS(b *testing.B) {

	var numbers []int
	for i := 0; i < 10000000; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	for n := 0; n < b.N; n++ {
		sumMaxsQS(numbers)
	}
}

func BenchmarkSumMaxs(b *testing.B) {

	var numbers []int
	for i := 0; i < 10000000; i++ {
		numbers = append(numbers, rand.Intn(100))
	}

	for n := 0; n < b.N; n++ {
		sumMaxs(numbers)
	}
}
