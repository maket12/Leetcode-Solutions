class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_table = {}
        for i in range(len(nums)):
            if hash_table.get(nums[i]) is None:
                hash_table[nums[i]] = []
            hash_table[nums[i]].append(i)
        
        for i in range(len(nums)):
            diff = target - nums[i]
            if diff in hash_table:
                for idx in hash_table[diff]:
                    if idx != i:
                        return [i, idx]
        return []
