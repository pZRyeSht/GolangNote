func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    t := len(nums) / 2
    if nums[t] != nums[t-1] && nums[t] != nums[t+1] {
        return nums[t]
    } else if nums[t] == nums[t-1] {
        if t % 2 ==0 {
            return singleNonDuplicate(nums[:t-1])
        } else {
            return singleNonDuplicate(nums[t+1:])
        }
    } else {
        if t % 2 ==0 {
            return singleNonDuplicate(nums[t+2:])
        } else {
            return singleNonDuplicate(nums[:t])
        }
    }
}