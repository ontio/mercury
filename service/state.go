package service

type ConnectionState int

const (
	ConnectionInit ConnectionState = iota
	ConnectionRequestReceived
	ConnectionResponseSent
	ConnectionACKRec
)
