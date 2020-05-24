//两数之和
func twoSum(nums []int, target int) []int {
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		c, ok := v[dif]
		if ok != false {
			return []int{c, i}
		}
		v[nums[i]] = i
	}
	return []int{-1, -1}
}
