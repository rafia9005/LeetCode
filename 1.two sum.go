package twosum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for l := i + 1; l < len(nums); l++ {
			if nums[i]+nums[l] == target {
				return []int{i, l}
			}
		}
	}
	return nil
}
