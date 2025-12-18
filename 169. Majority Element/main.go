func majorityElement(nums []int) int {
    need := len(nums) / 2
    m := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        m[nums[i]]++
        if m[nums[i]] > need {
            return nums[i]
        }
    }

    return 0
}