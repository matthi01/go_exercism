package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var count int
	if _, ok := cb[rank]; !ok {
		return 0
	}
	r := cb[rank]
	for _, square := range r {
		if square {
			count++
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var count int
	if file <= 0 || file > 8 {
		return 0
	}
	for _, rank := range cb {
		if rank[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var count int
	for rangeIndex := range cb {
		count += len(cb[rangeIndex])
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var count int
	for rankIndex := range cb {
		count += CountInRank(cb, rankIndex)
	}
	return count
}
