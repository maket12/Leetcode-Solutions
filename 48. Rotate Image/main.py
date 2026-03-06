class Solution:
    def rotate(self, matrix: list[list[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for row in range(n // 2):
            for column in range(row, n - 1 - row):
                matrix[row][column], matrix[column][n-1-row], matrix[n-1-row][n-1-column], matrix[n-1-column][row] = matrix[n-1-column][row], matrix[row][column], matrix[column][n-1-row], matrix[n-1-row][n-1-column]
            