class Solution:
    def diagonalSum(self, mat: list[list[int]]) -> int:
        result = 0
        for shift in range(len(mat)):
            result += mat[shift][shift]
            result += mat[shift][len(mat)-shift-1]
        if len(mat) % 2 != 0:
            diff = len(mat) // 2
            result -= mat[diff][diff]
        return result
