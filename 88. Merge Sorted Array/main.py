class Solution(object):
    def merge(self, nums1, m, nums2, n):
        """
        :type nums1: List[int]
        :type m: int
        :type nums2: List[int]
        :type n: int
        :rtype: None Do not return anything, modify nums1 in-place instead.
        """
        new = []
        k = 0
        i, j = 0, 0
        while i < m and j < n:
            if j < n and nums1[i] > nums2[j]:
                new[k] = nums2[j]
                j += 1
            else:
                new[k]
