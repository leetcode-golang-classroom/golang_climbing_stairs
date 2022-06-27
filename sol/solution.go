package sol

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	prev_2_step := 2
	prev_1_step := 3
	way := 0
	for stair := 4; stair <= n; stair++ {
		way = prev_1_step + prev_2_step
		prev_2_step = prev_1_step
		prev_1_step = way
	}
	return way
}
