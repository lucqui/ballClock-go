package model

import "strconv"

type Track struct {
	size  int
	index int
	balls []*Ball
}

func NewTrack(size int) *Track {
	balls := make([]*Ball, size)
	return &Track{
		size:  size,
		balls: balls,
		index: -1,
	}
}

func (t *Track) AddBall(b *Ball) ([]*Ball, *Ball) {
	ind := t.index + 1
	t.index = ind % t.size
	if ind >= t.size {
		t.index = -1
		return reverseArray(t.balls), b
	} else {
		t.balls[t.index] = b
		return nil, nil
	}
}

func (t *Track) IsFull() bool {
	return t.index == t.size-1
}

func (t *Track) TakeBall() *Ball {
	b := t.balls[0]
	t.balls = t.balls[1:]
	t.balls = append(t.balls, nil)
	t.index--
	return b
}

func (t *Track) NumBalls() int {
	return t.index + 1
}

func reverseArray(balls []*Ball) []*Ball {
	b := make([]*Ball, len(balls))
	j := len(balls)
	for i := 0; i < len(balls); i++ {
		j--
		b[i] = balls[j]
	}
	return b
}

func (t *Track) String() string {
	s := "["
	for i := 0; i < t.index; i++ {
		s += strconv.Itoa(t.balls[i].Id)
		if i != t.index-1 {
			s += ", "
		}
	}
	s += "]"
	return s
}
