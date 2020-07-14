package packager

type MessageData struct {
	Data []byte `json:"data,omitempty"`
	MsgType int          `json:"msgtype,omitempty"`
	Sign []byte `json:"sign,omitempty"`
}

type Envelope struct {
	Message *MessageData `json:"message,omitempty"`
	FromDID string       `json:"fromdid,omitempty"`
	ToDID   string       `json:"todid,omitempty"`
}
