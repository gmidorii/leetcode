package main

func isValidSudoku(board [][]byte) bool {
	// rows
	for _, rows := range board {
		rowSet := map[byte]bool{}
		for _, col := range rows {
			if col == []byte(".")[0] {
				continue
			}
			_, ok := rowSet[col]
			if ok {
				return false
			}
			rowSet[col] = true
		}
	}
	// cols
	for i := 0; i < len(board); i++ {
		colSet := map[byte]bool{}
		for j := 0; j < len(board); j++ {
			if board[j][i] == []byte(".")[0] {
				continue
			}
			_, ok := colSet[board[j][i]]
			if ok {
				return false
			}
			colSet[board[j][i]] = true
		}
	}
	// 3x3 boards
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			boardSet := map[byte]bool{}
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[i+x][j+y] == []byte(".")[0] {
						continue
					}
					_, ok := boardSet[board[i+x][j+y]]
					if ok {
						return false
					}
					boardSet[board[i+x][j+y]] = true
				}
			}
		}
	}
	return true
}
