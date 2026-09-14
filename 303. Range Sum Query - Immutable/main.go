type NumArray struct {
    Arr []int
    PrefixSums []int
}


func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
        return NumArray{Arr: nums, PrefixSums: []int{}}
    }

    prefixSums := make([]int, len(nums))
	prefixSums[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSums[i] = nums[i] + prefixSums[i-1]
	}

	return NumArray{
		Arr: nums,
		PrefixSums: prefixSums,
	}
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.PrefixSums[right]
    }
    return this.PrefixSums[right] - this.PrefixSums[left - 1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */