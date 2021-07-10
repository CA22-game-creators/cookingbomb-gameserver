package domain

type StatusEnum int

const (
	UNSPECIFIED StatusEnum = iota
	CONNECTED
	DISCONNECTED_BY_CLIENT
)

func (s StatusEnum) IsConnectable() bool {
	switch s {
	case CONNECTED:
		return false
	default:
		return true
	}
}

func (s StatusEnum) IsActive() bool {
	switch s {
	case CONNECTED:
		return true
	default:
		return false
	}
}
