package model

type Ball struct {
	Id int
}

func NewBall(id int) *Ball {
	return &Ball{
		Id: id,
	}
}
