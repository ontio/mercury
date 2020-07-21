package controller

import (
	"git.ont.io/ontid/otf/common/log"
	"git.ont.io/ontid/otf/common/message"
	"git.ont.io/ontid/otf/service/common"
)

func ResolveForward(req interface{},
	msgsvr *common.MsgService,
	conn message.Connection,
	mtype common.MessageType) (bool, error) {
	//add forward logic
	routers := common.MergeRouter(conn.MyRouter, conn.TheirRouter)
	if !common.IsReceiver(msgsvr.Cfg.SelfDID, routers) {
		//forward message
		log.Infof("forward ConnectionRequestType message")
		err := msgsvr.HandleOutBound(common.OutboundMsg{
			Msg: common.Message{
				MessageType: mtype,
				Content:     req,
			},
			Conn: conn,
		})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
