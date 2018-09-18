package model

import (
	"errors"
	"fmt"
)

type Clock struct {
	MinuteTrack     *Track
	FiveMinuteTrack *Track
	HourTrack       *Track
	MainTrack       *Track
	Balls           []*Ball
}

func NewClock(size int) (*Clock, error) {
	if size < 27 || size > 127 {
		return nil, errors.New("Error: wrong input size")
	}
	minuteTrack := NewTrack(4)
	fiveMinuteTrack := NewTrack(11)
	hourTrack := NewTrack(11)
	mainTrack := NewTrack(size)
	balls := make([]*Ball, size)
	for i := 0; i < size; i++ {
		balls[i] = NewBall(i)
		mainTrack.AddBall(balls[i])
	}
	clock := &Clock{
		MinuteTrack:     minuteTrack,
		FiveMinuteTrack: fiveMinuteTrack,
		HourTrack:       hourTrack,
		MainTrack:       mainTrack,
		Balls:           balls,
	}
	return clock, nil
}

func (c *Clock) MoveArm() (bool, error) {
	b := c.moveArm()
	nr := c.MinuteTrack.NumBalls() + c.FiveMinuteTrack.NumBalls() + c.HourTrack.NumBalls() + c.MainTrack.NumBalls()
	if nr != len(c.Balls) {
		return b, errors.New(fmt.Sprintf("Wrong number of balls expect: %d, got: %d", len(c.Balls), nr))
	}
	return b, nil
}

func (c *Clock) moveArm() bool {
	b := c.MainTrack.TakeBall()
	balls, ball := c.MinuteTrack.AddBall(b)
	if balls == nil {
		return false
	}
	for _, b := range balls {
		c.MainTrack.AddBall(b)
	}
	balls, ball = c.FiveMinuteTrack.AddBall(ball)
	if balls == nil {
		return false
	}
	for _, b := range balls {
		c.MainTrack.AddBall(b)
	}
	balls, ball = c.HourTrack.AddBall(ball)
	if balls == nil {
		return false
	}
	for _, b := range balls {
		c.MainTrack.AddBall(b)
	}
	c.MainTrack.AddBall(ball)
	if !c.MainTrack.IsFull() {
		return false
	}
	balls = c.MainTrack.balls
	bo := true
	for i, b := range balls {
		if i != b.Id {
			bo = false
		}
	}
	return bo
}

func (c *Clock) String() string {
	return fmt.Sprintf("{\"Min\": %s, \"FiveMin\": %s, \"Hour\": %s, \"Main\": %s}", c.MinuteTrack.String(), c.FiveMinuteTrack.String(), c.HourTrack.String(), c.MainTrack.String())
}
