class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for row in range(len(board)):
            frequency = {}
            for i in board[row]:
                if i == ".":
                    continue

                if i not in frequency:
                    frequency[i] = 0

                frequency[i] += 1
                if frequency[i] > 1:
                    return False
        
        for column in range(len(board)):
            frequency = {}
            for row in range(len(board)):
                if board[row][column] == ".":
                    continue

                if board[row][column] not in frequency:
                    frequency[board[row][column]] = 0

                frequency[board[row][column]] += 1
                if frequency[board[row][column]] > 1:
                    return False
        
        for row in range(2, len(board), 3):
            for column in range(2, len(board), 3):
                frequency = {}
                for i in range(row - 2, row + 1):
                    for j in range(column - 2, column + 1):
                        if board[i][j] == ".":
                            continue

                        if board[i][j] not in frequency:
                            frequency[board[i][j]] = 0

                        frequency[board[i][j]] += 1
                        if frequency[board[i][j]] > 1:
                            return False

        return True
        