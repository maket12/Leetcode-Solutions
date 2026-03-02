class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        k = len(nums) - 1
        result = [0] * k

        i, j = 0, len(nums) - 1
        while i <= j:
            if abs(nums[i]) > abs(nums[j]):
                result[k] = nums[i] ** 2
                i += 1
            else:
                result[k] = nums[j] ** 2
                j -= 1
            k -= 1
        
        return result
