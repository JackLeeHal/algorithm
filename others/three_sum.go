package others

import "sort"

// leetcode 15
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		// 排除两数相等的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		// 这是后面两个较大的数加起来应该等于的数
		target := -nums[i]
		for j := i + 1; j < n; j++ {
			// 再次排除两数相等的情况
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 如果两个数加起来太大就往前找
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return res
}
