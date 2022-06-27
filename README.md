# golang_climbing_stairs

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

## Examples

**Example 1:**

```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

```

**Example 2:**

```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

```

**Constraints:**

- `1 <= n <= 45`

## 解析

題目給定一個正整數 n 代表要爬到 第n 階樓梯

假設每次 可以選擇爬 1 階或是 2 階

寫一個演算法算是 爬到第 n 階總共有幾種方式

要看爬到 n 階 

可以先思考有幾總可能性

最後一步是走 2 階 也就事前一步事先爬了 n-2 階

最後一步是走 1 階 也就事前一步事先爬了 n-1 階

climb(n) = climb(n-2) + climb(n-1), 對所有 n ≥3

![](https://i.imgur.com/uvoZBAr.png)

其中可以從分析知道

climb(2) = 2, 

爬兩階有 {1,1},{2} 這兩種選擇

climb(3) = 3

爬3階 有 {1,1,1}, {1,2}, {2,1} 三種選擇

![](https://i.imgur.com/acLHwwu.png)

## 程式碼
```go
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

```
## 困難點

1. 要看出問題具有遞迴結構，大問題可以拆分成多個子問題

## Solve Point

- [x]  找出初始條件 climbstair_2 = 2, climbstair_3 = 3
- [x]  逐步計算 climbstair_n =climbstair_n-1 + climbstair_n-2