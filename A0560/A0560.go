package main

import "fmt"

/*
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :
输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。

说明 :
数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
通过次数84,197提交次数187,423

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{-1, -1, 1}
	k := 1
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	m := map[int]int{}
	m[0]++
	sum := 0
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if _, ok := m[sum-k]; ok {
			ans += m[sum-k]
		}
		m[sum]++
	}
	return ans
}
