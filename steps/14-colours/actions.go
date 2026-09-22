package main

type Action interface{}

type EscapeAction struct{}

type MovementAction struct {
	DX, DY int
}
