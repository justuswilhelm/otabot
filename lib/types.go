package lib

import "set"

// Board color
const (
	W int8 = -1
	E int8 = 0
	B int8 = 1
	M int8 = 2 // Marker
)

var ColorMapping = map[int8]rune{
	E: '.',
	W: 'O',
	B: 'X',
}

var ColorMappingReverse = map[rune]int8{
	'.': E,
	'O': W,
	'X': B,
}

// Columns maps an index to a Go board column name
const Columns = "ABCDEFGHJKLMNOPQRST"

// GoState stores a go game state
type GoState struct {
	Board         []int8
	liberties     []int
	Komi          float64
	CurrentPlayer int8
	Size          int
	Ko            int
}

// Flood modi
const (
	FirstSweepM = iota
	ReduceM
	RemoveM
	CountM
)

// GoMove encodes information to make one move in the game
type GoMove struct {
	Player int8
	Pass   bool
	Index  int
}

type Group struct {
	Color      int8
	Points     set.Interface
	Surrounded bool
}

type Region struct {
	Points            set.Interface
	NeighboringColors set.Interface
}
