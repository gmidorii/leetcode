package main

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for m := k + 1; m < len(nums); m++ {
					if (nums[i] + nums[j] + nums[k] + nums[m]) == target {
						res = append(res, []int{nums[i], nums[j], nums[k], nums[m]})
					}
				}
			}
		}
	}
	//TODO: fix duplicate
	return res
}
