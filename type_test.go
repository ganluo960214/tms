package main

//go:generate tms -type=TType
type TType byte

const (
	TTypeCA TType = iota + 1
	TTypeCB
	TTypeCC
	TTypeCD
	TTypeCE
)
