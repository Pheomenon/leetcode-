package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，
判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：
答案中不可以包含重复的四元组。

示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
通过次数138,941提交次数352,562

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	nums := []int{1, 0, -1, 0, -2, 2, 0}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		tuple := threeSum(nums, i+1, target-nums[i])
		for _, v := range tuple {
			v = append(v, nums[i])
			ans = append(ans, v)
		}
		for i < n && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func threeSum(nums []int, start, target int) [][]int {
	ans := make([][]int, 0)
	n := len(nums) - 1
	for i := start; i < n; i++ {
		two := twoSum(nums, i+1, target-nums[i])
		for _, v := range two {
			v = append(v, nums[i])
			ans = append(ans, v)
		}
		for i < n && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func twoSum(nums []int, start, target int) [][]int {
	ans := make([][]int, 0)
	end := len(nums) - 1
	for start < end {
		left := nums[start]
		right := nums[end]
		sum := left + right
		if sum == target {
			ans = append(ans, []int{nums[start], nums[end]})
			for start < end && nums[start] == left {
				start++
			}
			for start < end && nums[end] == right {
				end--
			}
		} else if sum > target {
			for start < end && nums[end] == right {
				end--
			}
		} else if sum < target {
			for start < end && nums[start] == left {
				start++
			}
		}
	}
	return ans
}