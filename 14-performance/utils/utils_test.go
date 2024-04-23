package utils

import "testing"

func Benchmark_IsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(771)
	}
}

func Benchmark_GenPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenPrimes(3, 9999)
	}
}
