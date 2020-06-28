package service

import "git.ont.io/ontid/otf/message"

type ConnectionState int

const (
	InvitationInit ConnectionState = iota
	InvitationUsed
	ConnectionRequestSent
	ConnectionRequestReceived
	ConnectionResponseReceived
	ConnectionACKRec
)

type InvitationRec struct {
	Invitation message.Invitation `json:"invitation"`
	State      ConnectionState    `json:"state"`
}

type ConnectionRequestRec struct {
	ConnReq message.ConnectionRequest `json:"conn_req"`
	State   ConnectionState           `json:"state"`
}
