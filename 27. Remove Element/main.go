func removeElement(nums []int, val int) int {
    var k int
    for i, el := range nums {
        if el != val {
            nums[k] = nums[i]
            k++
        }
    }
    return k
}