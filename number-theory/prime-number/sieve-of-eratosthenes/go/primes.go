package primes

func List(n int) []int {
	result := make([]int, 0, 65536)
	state := make([]bool, n)

	for i := 3; i*i < n; i += 2 {
		if !state[i] {
			for k := i * i; k < n; k += 2 * i {
				state[k] = true
			}
		}
	}
	result = append(result, 2)
	for i := 3; i < n; i += 2 {
		if !state[i] {
			result = append(result, i)
		}
	}
	return result
}
