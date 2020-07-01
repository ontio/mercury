package service

import "git.ont.io/ontid/otf/message"

func ReverseConnection(conn message.Connection) message.Connection {
	return message.Connection{
		MyDid:          conn.TheirDid,
		TheirDid:       conn.MyDid,
		MyServiceId:    conn.TheirServiceId,
		TheirServiceId: conn.MyServiceId,
	}
}
