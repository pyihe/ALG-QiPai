package cards

import "github.com/pyihe/go-pkg/rands"

type Shuffler interface {
	Shuffle(cards []Card)
}

type randShuffle struct {
}

func NewRandomShuffle() Shuffler {
	return &randShuffle{}
}

func (rs *randShuffle) Shuffle(cs []Card) {
	var n = len(cs)
	if n > 0 {
		for i := 0; i < n; i++ {
			pos := rands.Int(0, n-1)
			cs[i], cs[pos] = cs[pos], cs[i]
		}
	}
}
