package medium

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowCheck := map[byte]bool{}
		colCheck := map[byte]bool{}

		for j := 0; j < 9; j++ {
			// Row check
			if board[i][j] != '.' {
				if rowCheck[board[i][j]] {
					return false
				}
				rowCheck[board[i][j]] = true
			}

			// Column check
			if board[j][i] != '.' {
				if colCheck[board[j][i]] {
					return false
				}
				colCheck[board[j][i]] = true
			}
		}
	}

	// Box check
	for blockRow := 0; blockRow < 3; blockRow++ {
		for blockCol := 0; blockCol < 3; blockCol++ {
			boxCheck := map[byte]bool{}
			startRow := blockRow * 3
			startCol := blockCol * 3

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := board[startRow+i][startCol+j]
					if val != '.' {
						if boxCheck[val] {
							return false
						}
						boxCheck[val] = true
					}
				}
			}
		}
	}
	return true
}
