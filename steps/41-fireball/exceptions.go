package main

type Impossible struct {
	Msg string
}

func (i Impossible) Error() string { return i.Msg }
