package Problem

/*
	简单的搜索，首先判断点击目标是否是地雷，是的话直接标记结束
	如果不是，则判断目标周围地雷数量，为 0 则标 B，对目标周围进行判断；非 0 则标数字，结束当前判断
*/

// leetcode submit region begin(Prohibit modification and deletion)
//goland:noinspection GoUnusedFunction
func updateBoard(board [][]byte, click []int) [][]byte {
	r, c := click[0], click[1]
	if board[r][c] == 'M' {
		board[r][c] = 'X'
		return board
	} else {
		return dfs(board, r, c)
	}
}

func dfs(board [][]byte, r, c int) [][]byte {
	if board[r][c] != 'E' {
		return board
	}
	count := getCountMines(board, r, c)
	if count == 0 {
		board[r][c] = 'B'
		for i := r - 1; i <= r+1; i++ {
			for j := c - 1; j <= c+1; j++ {
				if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
					board = dfs(board, i, j)
				}
			}
		}
	} else {
		board[r][c] = byte(count + '0')
	}

	return board
}

func getCountMines(board [][]byte, r, c int) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
				if board[i][j] == 'M' {
					count++
				}
			}
		}
	}
	return count
}

// leetcode submit region end(Prohibit modification and deletion)

// 2020-08-20 09:05:47

// Let's play the minesweeper game (Wikipedia, online game)!
//
// You are given a 2D char matrix representing the game board. 'M' represents an
// unrevealed mine, 'E' represents an unrevealed empty square, 'B' represents a re
// vealed blank square that has no adjacent (above, below, left, right, and all 4 d
// iagonals) mines, digit ('1' to '8') represents how many mines are adjacent to th
// is revealed square, and finally 'X' represents a revealed mine.
//
// Now given the next click position (row and column indices) among all the unre
// vealed squares ('M' or 'E'), return the board after revealing this position acco
// rding to the following rules:
//
//
// If a mine ('M') is revealed, then the game is over - change it to 'X'.
// If an empty square ('E') with no adjacent mines is revealed, then change it t
// o revealed blank ('B') and all of its adjacent unrevealed squares should be reve
// aled recursively.
// If an empty square ('E') with at least one adjacent mine is revealed, then ch
// ange it to a digit ('1' to '8') representing the number of adjacent mines.
// Return the board when no more squares will be revealed.
//
//
//
//
// Example 1:
//
//
// Input:
//
// [['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'M', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E'],
// ['E', 'E', 'E', 'E', 'E']]
//
// Click : [3,0]
//
// Output:
//
// [['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
// Explanation:
//
//
//
// Example 2:
//
//
// Input:
//
// [['B', '1', 'E', '1', 'B'],
// ['B', '1', 'M', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
// Click : [1,2]
//
// Output:
//
// [['B', '1', 'E', '1', 'B'],
// ['B', '1', 'X', '1', 'B'],
// ['B', '1', '1', '1', 'B'],
// ['B', 'B', 'B', 'B', 'B']]
//
// Explanation:
//
//
//
//
//
// Note:
//
//
// The range of the input matrix's height and width is [1,50].
// The click position will only be an unrevealed square ('M' or 'E'), which also
// means the input board contains at least one clickable square.
// The input board won't be a stage when game is over (some mines have been reve
// aled).
// For simplicity, not mentioned rules should be ignored in this problem. For ex
// ample, you don't need to reveal all the unrevealed mines when the game is over,
// consider any cases that you will win the game or flag any squares.
//
//
// 👍 102 👎 0
