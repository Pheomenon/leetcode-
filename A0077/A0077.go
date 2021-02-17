package main

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:
输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	combine(4, 2)

}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var helper func(start int, path []int)
	helper = func(start int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			helper(start+1, path)
			path = path[:len(path)-1]
		}
	}
	helper(1, []int{})
	return res
}