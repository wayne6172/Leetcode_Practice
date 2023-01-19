package Q974

func subarraysDivByK(nums []int, k int) int {

	m := make([]int, k)
	m[0] = 1

	ans := 0
	sum := 0
	for _, num := range nums {
		sum = (sum + num%k + k) % k
		ans += m[sum]
		m[sum]++
	}

	return ans
}
