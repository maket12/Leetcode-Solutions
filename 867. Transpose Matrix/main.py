class Solution:
    def transpose(self, matrix: list[list[int]]) -> list[list[int]]:
        result = []
        for column in range(len(matrix[0])):
            new_row = []
            for row in range(len(matrix)):
                new_row.append(matrix[row][column])
            result.append(new_row)
        return result
