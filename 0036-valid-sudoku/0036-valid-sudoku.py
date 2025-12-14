class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        tuples = []
        for i in range(0, 9):
            for j in range(0, 9):
                element = board[i][j]
                if element == ".":
                    continue
                tuples += [(i, element), (element, j), (i//3, j//3, element)]
        return len(tuples)== len(set(tuples))