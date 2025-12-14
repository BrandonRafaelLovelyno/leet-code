class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for i in range(0,3):
            for j in range(0,3):
                x1, y1 = i * 3, j * 3
                my_set = set()
                for k in range(x1, x1 + 3):
                    for l in range(y1, y1 + 3):
                        if board[k][l] != "." and board[k][l] in my_set:
                            return False
                        my_set.add(board[k][l])

        for i in range(0, 9):
            for j in range(0,9):
                if board[i][j] != ".":
                    for k in range(i+1, 9):
                        if board[k][j] == board[i][j]:
                            return False
                    for k in range(j+1, 9):
                        if board[i][k] == board[i][j]:
                            return False
        return True