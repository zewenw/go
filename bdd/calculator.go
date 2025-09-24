package bdd

type Calculator struct {
	A, B, Result int
}

func (c *Calculator) Add() {
	c.Result = c.A + c.B
}
