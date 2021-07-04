package domain

type StatusEnum int

const (
	UNSPECIFIED StatusEnum = iota
	CONNECTED
	DISCONNECTED_BY_CLIENT
)
