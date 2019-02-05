package main

import (
	"fmt"
	"my-projects/golang-algorithms/filter"
)

func main() {

	num := 1000000
	prime := filter.EratosthenesSieve(num)

	for i := 2; i <= num; i++ {
		if prime[i] == false {
			fmt.Println(i)
		}
	}
}
