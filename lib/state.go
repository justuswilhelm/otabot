// Package lib GoState management part
package lib

import (
	"bytes"
	"fmt"
	"log"
)

// NewGoState creates a new GoState given size and komi
func NewGoState(size int, komi float64) *GoState {
	l := size * size
	return &GoState{
		CurrentPlayer: B,
		Board:         make([]int8, l),
		liberties:     make([]int, l),
		Komi:          komi,
		Size:          size,
	}
}

func (s *GoState) String() []string {
	var result []string
	for r := 0; r < s.Size; r++ {
		var row bytes.Buffer
		for c := 0; c < s.Size; c++ {
			col := s.Get(s.Index(r, c))
			row.WriteRune(ColorMapping[col])
		}
		result = append(result, row.String())
	}
	return result
}

// Log returns a string representation of the board
func (s *GoState) Log() string {
	var result bytes.Buffer
	result.WriteString("\n")
	result.WriteString("  ")

	for c := 0; c < s.Size; c++ {
		result.WriteString(fmt.Sprintf("%c ", Columns[c]))
	}
	result.WriteString("\n")
	for r := 0; r < s.Size; r++ {
		result.WriteString(fmt.Sprintf("%d ", s.Size-r))
		for c := 0; c < s.Size; c++ {
			col := s.Get(s.Index(r, c))
			result.WriteRune(ColorMapping[col])
			result.WriteString(" ")
		}
		result.WriteString(fmt.Sprintf("%d", s.Size-r))
		result.WriteString("\n")
	}
	result.WriteString("  ")
	for c := 0; c < s.Size; c++ {
		result.WriteString(fmt.Sprintf("%c ", Columns[c]))
	}
	return result.String()
}

func fromString(board []string) *GoState {
	size := len(board[0])
	s := NewGoState(size, 0.0)
	for r, row := range board {
		for c, col := range row {
			value, ok := ColorMappingReverse[col]
			if !ok {
				log.Panicf("Could not find %c at %dx%d", value, r, c)
			}
			s.Set(s.Index(r, c), value)
		}
	}
	return s
}
