package main

//S(n)=S(n-1)+S(n-2)
func climbStairs(n int) int {
	s := make([]int, n)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	s[0] = 1
	s[1] = 2
	for i := 2; i <= n-1; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n-1]
}
