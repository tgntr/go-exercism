package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank = []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) (count int) {
	for _, file := range cb[rank] {
		if file {
			count++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (count int) {
	if file < 1 || file > 8 {
		return
	}

	for _, rank := range cb {
		if rank[file-1] {
			count++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (count int) {
	for _, rank := range cb {
		count += len(rank)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (count int) {
	for _, rank := range cb {
		for _, file := range rank {
			if file {
				count++
			}
		}
	}
	return
}
