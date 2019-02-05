package filter

import "testing"

/*
	go test -bench . eratosthenesSieve_test.go eratosthenesSieve.go
	go test -bench . -benchmem eratosthenesSieve_test.go eratosthenesSieve.go
	go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1 eratosthenesSieve_test.go eratosthenesSieve.go

*/

func BenchmarkEratosthenesSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EratosthenesSieve(10000)
	}
}
