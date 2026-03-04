class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        nums.sort()
        result = []
        for n in range(len(nums) - 2):
            if n != 0 and nums[n] == nums[n-1]:
                continue

            i, j = n+1, len(nums) - 1
            target = -1 * nums[n]
            while i < j:
                if target == nums[i] + nums[j]:
                    result.append([nums[n], nums[i], nums[j]])
                    i += 1
                    j -= 1

                    while i < j and nums[i] == nums[i-1]:
                        i += 1
                    while j > i and nums[j] == nums[j+1]:
                        j -= 1
                        
                elif nums[i] + nums[j] > target:
                    j -= 1
                else:
                    i += 1
        
        return result
