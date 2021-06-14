package main

//go:generate tms -type=Type
type Type byte

const (
	TypeCA Type = iota + 1
	TypeCB
	TypeCC
	TypeCD
	TypeCE
)
