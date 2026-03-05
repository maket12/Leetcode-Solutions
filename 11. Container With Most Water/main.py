class Solution:
    def countArea(self, x1: int, y1: int, x2: int, y2: int) -> int:
        return (x2 - x1) * min(y1, y2)

    def maxArea(self, height: List[int]) -> int:
        area = 0
        i, j = 0, len(height) - 1
        while i < j:
            area = max(area, self.countArea(i, height[i], j, height[j]))
            if height[i] < height[j]:
                i += 1
            else:
                j -= 1
        return area
