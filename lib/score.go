package lib

import (
	"set"
)

// MakeEmptyRegion constructs a region consisting of empty points around p
// if p is not empty, it will return a region consisting of emtpy points
// neighboring p and p
func (s *GoState) MakeEmptyRegion(p int) *Region {
	reg := &Region{
		Points:            set.New(set.ThreadSafe),
		NeighboringColors: set.New(set.ThreadSafe),
	}
	goals := set.New(set.ThreadSafe)
	goals.Add(p)
	for goals.Size() > 0 {
		point := goals.Pop().(int)
		reg.Points.Add(point)
		for _, neighbor := range s.Neighbors(point) {
			neighborColor := s.Get(neighbor)
			if neighborColor == E {
				if !reg.Points.Has(neighbor) {
					goals.Add(neighbor)
				}
			} else {
				reg.NeighboringColors.Add(neighborColor)
			}
		}
	}
	return reg
}

func (s *GoState) ScoreComponents() map[int8]int {
	handled := make([]bool, len(s.Board))
	score := map[int8]int{
		W: 0,
		B: 0,
	}
	for p, c := range s.Board {
		if c != E {
			score[c] += 1
			continue
		}
		if handled[p] {
			continue
		}
		region := s.MakeEmptyRegion(p)
		region_size := region.Points.Size()
		for _, c := range []int8{B, W} {
			if region.NeighboringColors.Has(c) {
				score[c] += region_size
			}
		}
		for _, pRaw := range region.Points.List() {
			p := pRaw.(int)
			handled[p] = true
		}
	}
	return score
}
func (s *GoState) Score() float64 {
	score := s.ScoreComponents()
	return float64(score[B]-score[W]) - s.Komi
}
