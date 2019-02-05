package filter

// Для вывода всех найденных простых чисел:
// for i := 2; i <= num; i++ {
//		if prime[i] == false {
//			fmt.Println(i)
//		}
//	}
func EratosthenesSieve(num int) []bool {

	prime := make([]bool, num+1)

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true
			}
		}

	}

	return prime

}
