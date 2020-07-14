package service

import "git.ont.io/ontid/otf/message"

func ReverseConnection(conn message.Connection) message.Connection {
	return message.Connection{
		MyDid:       conn.TheirDid,
		MyRouter:    conn.TheirRouter,
		TheirDid:    conn.MyDid,
		TheirRouter: conn.MyRouter,
		//MyServiceId:    conn.TheirServiceId,
		//TheirServiceId: conn.MyServiceId,
	}
}
